package pkg

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/pallavidn/faas-istio/pkg/conf"
	"github.com/pallavidn/faas-istio/pkg/discovery"
	"github.com/pallavidn/faas-istio/pkg/registration"
	"github.com/turbonomic/turbo-go-sdk/pkg/probe"
	"github.com/turbonomic/turbo-go-sdk/pkg/service"
	"hash/fnv"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"os/signal"
	"syscall"
	"github.com/pallavidn/faas-istio/pkg/action"
)

type disconnectFromTurboFunc func()

type FaasIstioTAPService struct {
	tapService *service.TAPService
}

func NewFaasIstioTAPService(args *conf.FaasIstioTurboArgs) (*FaasIstioTAPService, error) {
	tapService, err := createTAPService(args)

	if err != nil {
		glog.Errorf("Error while building turbo TAP service on target %v", err)
		return nil, err
	}

	return &FaasIstioTAPService{tapService}, nil
}

func (p *FaasIstioTAPService) Start() {
	glog.V(0).Infof("Starting Faas Istio TAP service...")

	// Disconnect from Turbo server when kongturbo is shutdown
	handleExit(func() { p.tapService.DisconnectFromTurbo() })

	// Connect to the Turbo server
	p.tapService.ConnectToTurbo()

	select {}
}

func createTAPService(args *conf.FaasIstioTurboArgs) (*service.TAPService, error) {
	confPath := args.TurboConf

	conf, err := conf.NewFaasIstioTurboServiceSpec(confPath)
	if err != nil {
		glog.Errorf("Error while parsing the service config file %s: %v", confPath, err)
		os.Exit(1)
	}

	glog.V(3).Infof("Read service configuration from %s: %++v", confPath, conf)

	communicator := conf.TurboCommunicationConfig
	targetConf := conf.FaasIstioTurboTargetConf

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", args.KubeConf)
	if err != nil {
		glog.Errorf("Fatal error: failed to get kubeconfig:  %s", err)
		os.Exit(1)
	}

	targetConf.TargetAddress = targetConf.TargetType + "-" + kubeConfig.Host //conf.KnativeTurboTargetConf.Kubeconfig

	registrationClient := &registration.FaasIstioTurboRegistrationClient{}
	discoveryClient := discovery.NewDiscoveryClient(targetConf, kubeConfig)

	// Action Execution Client
	actionHandler := action.NewActionHandler(kubeConfig)

	targetType := targetConf.TargetType + "-" + fmt.Sprint(hash(targetConf.TargetAddress))

	return service.NewTAPServiceBuilder().
		WithTurboCommunicator(communicator).
		WithTurboProbe(probe.NewProbeBuilder(targetType, targetConf.ProbeCategory).
			WithDiscoveryOptions(probe.FullRediscoveryIntervalSecondsOption(int32(*args.DiscoveryIntervalSec))).
			RegisteredBy(registrationClient).
			WithActionPolicies(registrationClient).
			WithEntityMetadata(registrationClient).
			DiscoversTarget(targetConf.TargetAddress, discoveryClient).
			ExecutesActionsBy(actionHandler)).
		Create()
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// TODO: Move the handle to turbo-sdk-probe as it should be common logic for similar probes
// handleExit disconnects the tap service from Turbo service when kongturbo is terminated
func handleExit(disconnectFunc disconnectFromTurboFunc) {
	glog.V(4).Infof("*** Handling Knativeturbo Termination ***")
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGHUP)

	go func() {
		select {
		case sig := <-sigChan:
			// Close the mediation container including the endpoints. It avoids the
			// invalid endpoints remaining in the server side. See OM-28801.
			glog.V(2).Infof("Signal %s received. Disconnecting from Turbo server...\n", sig)
			disconnectFunc()
		}
	}()
}
