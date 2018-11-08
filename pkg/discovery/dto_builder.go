package discovery

import (
	"fmt"
	"github.com/turbonomic/turbo-go-sdk/pkg/builder"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

type FaasIstioEntityProp string

const (
	StitchingAttr string = "VappIds"

	LocalNameAttr    string = "LocalName"
	AltNameAttr      string = "altName"
	ExternalNameAttr string = "externalnames"
	NsAttr  FaasIstioEntityProp = "namespace"
	VsAttr  FaasIstioEntityProp = "virtualservice"
	FuncAttr FaasIstioEntityProp = "function"
	ClientAttr FaasIstioEntityProp = "client"

	DefaultPropertyNamespace string = "DEFAULT"
	PropertyUsed                    = "used"
	PropertyCapacity                = "capacity"
)

type FaasIstioDTOBuilder struct {
}

func (dtoBuilder *FaasIstioDTOBuilder) buildFunctionDto(clientApp *ClientApp) (*builder.EntityDTOBuilder, error) {
	if clientApp == nil {
		return nil, fmt.Errorf("Null service for %++v", clientApp)
	}

	// id.
	altNameId := fmt.Sprintf("%s/%s/%s",
			clientApp.AppLabel, clientApp.GatewayHost, clientApp.GatewayEndpoint)
	localNameId := clientApp.FunctionEndpoint

	// uuid and display name.
	vappId := fmt.Sprintf("%s-%s/%s/%s", "virtualservice",
			clientApp.AppLabel, clientApp.GatewayEndpoint, clientApp.FunctionEndpoint)
	fmt.Printf("**** vapp id : %s\n", vappId)

	transKey := fmt.Sprintf("%s/%s", clientApp.GatewayHost, clientApp.GatewayEndpoint)
	commodities := []*proto.CommodityDTO{}
	commodity, _ := builder.NewCommodityDTOBuilder(proto.CommodityDTO_TRANSACTION).
			Key(transKey).
			Create()
	commodities = append(commodities, commodity)

	entityDTOBuilder := builder.NewEntityDTOBuilder(proto.EntityDTO_VIRTUAL_APPLICATION, vappId).
		DisplayName(vappId).
		//WithProperty(getEntityProperty(StitchingAttr, vappId)).
		WithProperty(getEntityProperty(LocalNameAttr, localNameId)).	//TODO: CHANGE to alNameId for LOCAL DEBUG
		//WithProperty(getEntityProperty(LocalNameAttr, altNameId)).	//TODO: CHANGE BACK to localNameId AFTER LOCAL DEBUD
		WithProperty(getEntityProperty(AltNameAttr, altNameId)).
		WithProperty(getEntityProperty(string(ClientAttr), clientApp.AppLabel)).
		WithProperty(getEntityProperty(string(FuncAttr), clientApp.FunctionEndpoint)).
		WithProperty(getEntityProperty(string(VsAttr), clientApp.VirtualServiceName)).
		WithProperty(getEntityProperty(string(NsAttr), clientApp.KubernetesNamespace)).
		SellsCommodities(commodities).
		ReplacedBy(getReplacementMetaData(proto.EntityDTO_VIRTUAL_APPLICATION)) //for stitching with kubeturbo vapps

	fmt.Printf("Created function dto builder\n")
	return entityDTOBuilder, nil
}

func getReplacementMetaData(entityType proto.EntityDTO_EntityType,
) *proto.EntityDTO_ReplacementEntityMetaData {
	extAttr := ExternalNameAttr		//StitchingAttr
	intAttr := LocalNameAttr
	useTopoExt := true

	b := builder.NewReplacementEntityMetaDataBuilder().
		Matching(intAttr).
		MatchingExternal(&proto.ServerEntityPropDef{
			Entity:     &entityType,
			Attribute:  &extAttr,
			UseTopoExt: &useTopoExt,
		})

	return b.Build()
}

func getEntityProperty(attr, value string) *proto.EntityDTO_EntityProperty {
	ns := DefaultPropertyNamespace

	return &proto.EntityDTO_EntityProperty{
		Namespace: &ns,
		Name:      &attr,
		Value:     &value,
	}
}
