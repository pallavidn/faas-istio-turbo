package discovery

import (
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"

	"fmt"
	v1alpha3 "github.com/pallavidn/faas-istio/pkg/apis/istio/v1alpha3"
	istiov1alpha3 "github.com/pallavidn/faas-istio/pkg/client/clientset/versioned/typed/istio/v1alpha3"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/golang/glog"
	restclient "k8s.io/client-go/rest"
)

type ClientApp struct {
	AppLabel            string
	GatewayHost         string
	GatewayEndpoint     string
	FunctionEndpoint    string
	VirtualServiceName  string
	KubernetesNamespace string
}

func DiscoverIstio(kubeClientConfig *restclient.Config) ([]*proto.EntityDTO, error) {
	namespace := apiv1.NamespaceAll

	istioClient, err := istiov1alpha3.NewForConfig(kubeClientConfig)
	if err != nil {
		fmt.Errorf("Error creating knative istio client: %++v", err)
	}

	fmt.Printf("Got faas istio client\n")
	// Get istio services
	virtualSvcList, err := istioClient.VirtualServices(namespace).List(metav1.ListOptions{})
	if err != nil {
		fmt.Errorf("Error while getting istio virtual services %++v", err)
	}
	fmt.Printf("Got istio virtual services %d\n", len(virtualSvcList.Items))

	var clientApps []*ClientApp
	for _, svc := range virtualSvcList.Items {
		fmt.Printf("*********** Virtual Service %s\n", svc.Name)
		vsSpec := svc.Spec
		//fmt.Printf("Hosts : %s\n", vsSpec.Hosts)
		httproutes := vsSpec.Http

		for _, httproute := range httproutes {
			clientApp := ParseClientApp(svc, httproute)
			if clientApp != nil {
				clientApps = append(clientApps, clientApp)
			}
		}
	}

	fmt.Printf("Building DTOs for %d client apps\n", len(clientApps))
	istiodtoBuilder := FaasIstioDTOBuilder{}
	var discoveryResult []*proto.EntityDTO
	for _, clientApp := range clientApps {

		dtoBuilder, err := istiodtoBuilder.buildFunctionDto(clientApp)
		if err != nil {
			glog.Errorf("%s", err)
			fmt.Printf("Error while building entity : %v\n", err)
		}
		if dtoBuilder == nil {
			fmt.Printf("%v\n", err)
			continue
		}
		dto, err := dtoBuilder.Create()
		if err != nil {
			fmt.Printf("builder error : %v\n", err)
		}
		fmt.Printf("Function DTO %++v\n", dto)
		discoveryResult = append(discoveryResult, dto)
	}
	fmt.Printf("DONE Building %d DTOs\n", len(discoveryResult))
	return discoveryResult, nil
}

func ParseClientApp(virtualSvc v1alpha3.VirtualService, httpRoute v1alpha3.HTTPRoute) *ClientApp {
	var sourceApp string
	var gatewayHost string
	var gatewayEndpoint, functionEndpoint string

	for _, match := range httpRoute.Match {
		// gateway endpoint
		if match.Uri != nil {
			gatewayEndpoint = match.Uri.Exact[1:]
		}
		// client identifier
		//fmt.Printf("source labels %s\n", match.SourceLabels)
		sourceLabelMap := match.SourceLabels
		sourceApp, _ = sourceLabelMap["app"]
	}

	// serverless function endpoint
	if httpRoute.Rewrite != nil {
		functionEndpoint = httpRoute.Rewrite.Uri[1:]
	}
	//
	for _, route := range httpRoute.Route {
		gatewayHost = route.Destination.Host
	}

	if sourceApp != "" && gatewayHost != "" {
		fmt.Printf("Parsed sourceApp:%s, gatewayEndpoint:%s, functionEndpoint:%s, gatewayHost:%s\n",
			sourceApp, gatewayEndpoint, functionEndpoint, gatewayHost)

		clientApp := &ClientApp{
			AppLabel:            sourceApp,
			GatewayHost:         gatewayHost,
			GatewayEndpoint:    gatewayEndpoint,
			FunctionEndpoint:        functionEndpoint,
			VirtualServiceName:  virtualSvc.Name,
			KubernetesNamespace: virtualSvc.Namespace,
		}
		return clientApp
	}
	return nil
}

//for key, val := range sourceLabelMap {
//	if key == "app" {
//		sourceApp = val
//		break
//	}
//	fmt.Printf("%s = %++v\n", key, val)
//}
