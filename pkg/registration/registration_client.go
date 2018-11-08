package registration

import (
	"github.com/golang/glog"
	"github.com/turbonomic/turbo-go-sdk/pkg/builder"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

const (
	TargetIdentifierField string = "targetIdentifier"
	propertyId            string = "id"
)

// Implements the TurboRegistrationClient interface
type FaasIstioTurboRegistrationClient struct {
}

func (p *FaasIstioTurboRegistrationClient) GetSupplyChainDefinition() []*proto.TemplateDTO {
	glog.Infoln("Building a supply chain ..........")

	supplyChainFactory := &SupplyChainFactory{}
	templateDtos, err := supplyChainFactory.CreateSupplyChain()
	if err != nil {
		glog.Error("Error creating Supply chain for KnativeTurbo")
		return nil
	}
	glog.Infoln("Supply chain for KnativeTurbo is created.")
	return templateDtos
}

func (p *FaasIstioTurboRegistrationClient) GetIdentifyingFields() string {
	return TargetIdentifierField
}

func (p *FaasIstioTurboRegistrationClient) GetAccountDefinition() []*proto.AccountDefEntry {

	targetIDAcctDefEntry := builder.NewAccountDefEntryBuilder(TargetIdentifierField, "kubeconfig",
		"Path to the kubeconfig", ".*", true, false).Create()

	return []*proto.AccountDefEntry{
		targetIDAcctDefEntry,
	}
}

func (rclient *FaasIstioTurboRegistrationClient) GetEntityMetadata() []*proto.EntityIdentityMetadata {
	glog.V(3).Infof("Begin to build EntityIdentityMetadata")

	result := []*proto.EntityIdentityMetadata{}

	entities := []proto.EntityDTO_EntityType{
		proto.EntityDTO_VIRTUAL_APPLICATION,
	}

	for _, etype := range entities {
		meta := rclient.newIdMetaData(etype, []string{propertyId})
		result = append(result, meta)
	}

	glog.V(4).Infof("EntityIdentityMetaData: %++v", result)

	return result
}

func (rclient *FaasIstioTurboRegistrationClient) newIdMetaData(etype proto.EntityDTO_EntityType, names []string) *proto.EntityIdentityMetadata {
	data := []*proto.EntityIdentityMetadata_PropertyMetadata{}
	for _, name := range names {
		dat := &proto.EntityIdentityMetadata_PropertyMetadata{
			Name: &name,
		}
		data = append(data, dat)
	}

	result := &proto.EntityIdentityMetadata{
		EntityType:            &etype,
		NonVolatileProperties: data,
	}

	return result
}

func (rClient *FaasIstioTurboRegistrationClient) GetActionPolicy() []*proto.ActionPolicyDTO {
	glog.V(23).Infof("Begin to build Action Policies")
	ab := builder.NewActionPolicyBuilder()
	supported := proto.ActionPolicyDTO_SUPPORTED
	notSupported := proto.ActionPolicyDTO_NOT_SUPPORTED

	//1. containerPod: move, provision; not resize;
	pod := proto.EntityDTO_VIRTUAL_APPLICATION
	podPolicy := make(map[proto.ActionItemDTO_ActionType]proto.ActionPolicyDTO_ActionCapability)
	podPolicy[proto.ActionItemDTO_MOVE] = supported
	podPolicy[proto.ActionItemDTO_PROVISION] = notSupported
	podPolicy[proto.ActionItemDTO_RIGHT_SIZE] = notSupported
	podPolicy[proto.ActionItemDTO_SUSPEND] = notSupported

	rClient.addActionPolicy(ab, pod, podPolicy)

	return ab.Create()
}

func (rClient *FaasIstioTurboRegistrationClient) addActionPolicy(ab *builder.ActionPolicyBuilder,
	entity proto.EntityDTO_EntityType,
	policies map[proto.ActionItemDTO_ActionType]proto.ActionPolicyDTO_ActionCapability) {

	for action, policy := range policies {
		ab.WithEntityActions(entity, action, policy)
	}
}
