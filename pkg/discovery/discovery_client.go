package discovery

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/pallavidn/faas-istio/pkg/conf"
	"github.com/pallavidn/faas-istio/pkg/registration"
	"github.com/turbonomic/turbo-go-sdk/pkg/probe"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"k8s.io/client-go/rest"
)

// Implements the TurboDiscoveryClient interface
type FaasIstioDiscoveryClient struct {
	kubeConfig   *rest.Config
	targetConfig *conf.FaasIstioTurboTargetConf
}

func NewDiscoveryClient(targetConfig *conf.FaasIstioTurboTargetConf, kubeConfig *rest.Config) *FaasIstioDiscoveryClient {
	glog.V(2).Infof("New Discovery client for kubernetes host: %s", kubeConfig.Host)
	return &FaasIstioDiscoveryClient{
		targetConfig: targetConfig,
		kubeConfig:   kubeConfig,
	}
}

// Get the Account Values to create VMTTarget in the turbo server corresponding to this client
func (d *FaasIstioDiscoveryClient) GetAccountValues() *probe.TurboTargetInfo {
	targetId := registration.TargetIdentifierField
	targetConf := d.targetConfig
	targetIdVal := &proto.AccountValue{
		Key:         &targetId,
		StringValue: &targetConf.TargetAddress,
	}

	accountValues := []*proto.AccountValue{
		targetIdVal,
	}

	targetInfo := probe.NewTurboTargetInfoBuilder(targetConf.ProbeCategory, targetConf.TargetType,
		registration.TargetIdentifierField, accountValues).Create()

	return targetInfo
}

// Validate the Target
func (d *FaasIstioDiscoveryClient) Validate(accountValues []*proto.AccountValue) (*proto.ValidationResponse, error) {
	glog.V(2).Infof("Validating Faas Istio target %s", accountValues)
	fmt.Printf("Validating Faas Istio target: %s\n", accountValues)
	// TODO: Add logic for validation
	validationResponse := &proto.ValidationResponse{}

	// Validation fails if no exporter responses
	return validationResponse, nil
}

// Discover the Target Topology
func (d *FaasIstioDiscoveryClient) Discover(accountValues []*proto.AccountValue) (*proto.DiscoveryResponse, error) {
	glog.V(2).Infof("Discovering Faas Istio target %s", accountValues)
	fmt.Printf("Discovering Faas Istio target %s\n", accountValues)
	var entities []*proto.EntityDTO

	var discoveryResponse *proto.DiscoveryResponse
	entities, err := DiscoverIstio(d.kubeConfig)
	if err != nil {
		fmt.Printf("Discovery failure %++v\n", err)
		return d.failDiscovery(), nil
	}

	discoveryResponse = &proto.DiscoveryResponse{
		EntityDTO: entities,
	}
	fmt.Printf("DONE Discovering Faas Istio target: %s\n", d.kubeConfig.Host)
	return discoveryResponse, nil
}

func (d *FaasIstioDiscoveryClient) failDiscovery() *proto.DiscoveryResponse {
	description := fmt.Sprintf("FaasIstioTurbo probe discovery failed")
	glog.Errorf(description)
	severity := proto.ErrorDTO_CRITICAL
	errorDTO := &proto.ErrorDTO{
		Severity:    &severity,
		Description: &description,
	}
	discoveryResponse := &proto.DiscoveryResponse{
		ErrorDTO: []*proto.ErrorDTO{errorDTO},
	}
	return discoveryResponse
}

func (d *FaasIstioDiscoveryClient) failValidation() *proto.ValidationResponse {
	description := fmt.Sprintf("FaasIstioTurbo probe validation failed")
	glog.Errorf(description)
	severity := proto.ErrorDTO_CRITICAL
	errorDto := &proto.ErrorDTO{
		Severity:    &severity,
		Description: &description,
	}

	validationResponse := &proto.ValidationResponse{
		ErrorDTO: []*proto.ErrorDTO{errorDto},
	}
	return validationResponse
}
