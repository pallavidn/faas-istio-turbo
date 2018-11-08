package conf

import (
	"flag"
)

const (
	defaultDiscoveryIntervalSec = 600
	DefaultConfPath             = "/etc/faas-istioturbo/turbo.config"
)

type FaasIstioTurboArgs struct {
	DiscoveryIntervalSec *int
	TurboConf            string
	KubeConf             string
}

func NewFaasIstioTurboArgs(fs *flag.FlagSet) *FaasIstioTurboArgs {
	p := &FaasIstioTurboArgs{}

	p.DiscoveryIntervalSec = fs.Int("discovery-interval-sec", defaultDiscoveryIntervalSec, "The discovery interval in seconds")
	fs.StringVar(&p.TurboConf, "turboconfig", p.TurboConf, "Path to the turbo config file.")
	fs.StringVar(&p.KubeConf, "kubeconfig", p.KubeConf, "Path to kubeconfig file with authorization and master location information.")

	return p
}
