package conf

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/turbonomic/turbo-go-sdk/pkg/service"
	"io/ioutil"
)

const (
	defaultProbeCategory = "Cloud Native"
	defaultTargetType    = "Faas-Istio"
)

type FaasIstioTurboServiceSpec struct {
	*service.TurboCommunicationConfig `json:"communicationConfig,omitempty"`
	*FaasIstioTurboTargetConf         //`json:"faasIstioTurboTargetConfig,omitempty"`
}

type FaasIstioTurboTargetConf struct {
	ProbeCategory string `json:"probeCategory,omitempty"`
	TargetType    string `json:"targetType,omitempty"`
	TargetAddress string
	Kubeconfig    string `json:"kubeconfig,omitempty"`
}

func NewFaasIstioTurboServiceSpec(configFilePath string) (*FaasIstioTurboServiceSpec, error) {

	glog.Infof("Read configuration from %s", configFilePath)
	tapSpec, err := readConfig(configFilePath)

	if err != nil {
		return nil, err
	}

	if tapSpec.TurboCommunicationConfig == nil {
		return nil, fmt.Errorf("Unable to read the turbo communication config from %s", configFilePath)
	}

	tapSpec.FaasIstioTurboTargetConf = &FaasIstioTurboTargetConf{
		ProbeCategory: defaultProbeCategory,
		TargetType:    defaultTargetType,
	}

	return tapSpec, nil
}

func readConfig(path string) (*FaasIstioTurboServiceSpec, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		glog.Errorf("File error: %v\n", err)
		return nil, err
	}
	glog.Infoln(string(file))

	var spec FaasIstioTurboServiceSpec
	err = json.Unmarshal(file, &spec)

	if err != nil {
		glog.Errorf("Unmarshall error :%v\n", err)
		return nil, err
	}
	glog.Infof("Results: %+v\n", spec)

	return &spec, nil
}
