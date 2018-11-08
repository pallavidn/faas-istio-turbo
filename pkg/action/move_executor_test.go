package action

import (
	"testing"
	"github.com/golang/glog"
	"os"
	"k8s.io/client-go/tools/clientcmd"
)

func TestMoveAction(t *testing.T) {

	kubeConfigFile := "/Users/pallavinayak/kubeconfig/the-shire/kubeconfig"

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigFile)
	if err != nil {
		glog.Errorf("Fatal error: failed to get kubeconfig:  %s", err)
		os.Exit(1)
	}

	ae := NewClientMoveActionExecutor(kubeConfig)
	ae.Execute(&TurboActionExecutorInput{})

}
