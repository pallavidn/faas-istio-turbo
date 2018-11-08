package action

import (
	"fmt"
	v1alpha3 "github.com/pallavidn/faas-istio/pkg/apis/istio/v1alpha3"
	istiov1alpha3 "github.com/pallavidn/faas-istio/pkg/client/clientset/versioned/typed/istio/v1alpha3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"github.com/pallavidn/faas-istio/pkg/discovery"
)

type ClientMoveActionExecutor struct {
	kubeConfig *rest.Config
}

func NewClientMoveActionExecutor(kubeConfig *rest.Config) *ClientMoveActionExecutor {
	return &ClientMoveActionExecutor{
		kubeConfig: kubeConfig,
	}
}
func (ae *ClientMoveActionExecutor) Execute(input *TurboActionExecutorInput) (*TurboActionExecutorOutput, error) {

	istioClient, err := istiov1alpha3.NewForConfig(ae.kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("Error creating knative istio client: %++v", err)
	}

	fmt.Printf("Got knative istio client\n")

	// Action Item
	actionItemDTO := input.ActionItem
	currEntity := actionItemDTO.GetCurrentSE()
	currEntityProps := currEntity.GetEntityProperties()

	newEntity := actionItemDTO.GetNewSE()
	newEntityProps := newEntity.GetEntityProperties()

	// Get the properties required to fetch the Istio VirtualService object
	var namespace, virtualservice, funcName, client string
	for _, prop := range currEntityProps {
		if *prop.Name == string(discovery.NsAttr) {
			namespace = *prop.Value
		}
		if *prop.Name == string(discovery.VsAttr) {
			virtualservice = *prop.Value
		}
		if *prop.Name == string(discovery.FuncAttr) {
			funcName = *prop.Value
		}
		if *prop.Name == string(discovery.ClientAttr) {
			client = *prop.Value
		}
	}

	var destFuncName string
	for _, prop := range newEntityProps {
		if *prop.Name == string(discovery.FuncAttr) {
			destFuncName = *prop.Value
		}
	}

	//namespace = "default"
	//virtualservice = "hello-kong"	//"test-vs"
	//client = "foo"
	//funcName = "hello-lambda"
	//destFuncName = "hello-knative"

	fmt.Printf("namespace: %s virtualservice: %s client: %s\n", namespace, virtualservice, client)
	fmt.Printf("To route client from function %s to function %s \n", funcName, destFuncName)

	var desiredSvc *v1alpha3.VirtualService
	desiredSvc, _ = istioClient.VirtualServices(namespace).Get(virtualservice, metav1.GetOptions{})
	if err != nil {
		fmt.Errorf("Error while getting %s istio virtual services %++v", virtualservice, err)
	}
	fmt.Printf("#### FOUND virtual service %++v\n", desiredSvc)

	httproutes := desiredSvc.Spec.Http
	fmt.Printf("%d number of routes\n", len(httproutes))

	var svcRoutes []v1alpha3.HTTPRoute
	for _, httproute := range httproutes {
		curr_rewrite := httproute.Rewrite
		//functionName := httproute.Rewrite.Uri[1:]
		//fmt.Printf("---> rewrite %s\n", functionName)

		clientApp := discovery.ParseClientApp(*desiredSvc, httproute)
		if clientApp == nil {
			continue
		}

		// Modify client app function endpoint
		if clientApp.FunctionEndpoint == funcName && clientApp.AppLabel == client {
			curr_rewrite.Uri = "/" + destFuncName
			fmt.Printf("	changed current rewrite uri to %s\n", curr_rewrite.Uri)
			//httproute.Rewrite = curr_rewrite //TODO
			//fmt.Printf("#### NEW %++v\n", httproute.Rewrite) //TODO
		}
		svcRoutes = append(svcRoutes, httproute)
	}

	desiredSvc.Spec.Http = svcRoutes

	httproutes1 := desiredSvc.Spec.Http
	fmt.Printf("%d number of routes\n", len(httproutes1))
	for _, httproute := range httproutes1 {
		fmt.Printf("%s: %d number of matches\n", httproute.Rewrite.Uri, len(httproute.Match))
		//functionName := httproute.Rewrite.Uri[1:]
		//fmt.Printf("---> rewrite %s\n", functionName)
		//fmt.Printf("#### NEW %++v\n", httproute.Rewrite)
	}

	_, err = istioClient.VirtualServices(namespace).Update(desiredSvc)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	return &TurboActionExecutorOutput{Succeeded: true}, nil
}

//namespace := apiv1.NamespaceAll
// Get istio services
//virtualSvcList, err := istioClient.VirtualServices(apiv1.NamespaceAll).List(metav1.ListOptions{})
//if err != nil {
//	fmt.Errorf("Error while getting istio virtual services %++v", err)
//}
//fmt.Printf("Got istio virtual services %d\n", len(virtualSvcList.Items))

//// Get istio services
//virtualSvcList, err := istioClient.VirtualServices(namespace).List(metav1.ListOptions{})
//if err != nil {
//	fmt.Errorf("Error while getting istio virtual services %++v", err)
//}
//fmt.Printf("Got istio virtual services %d\n", len(virtualSvcList.Items))
//
//for _, svc := range virtualSvcList.Items {
//	fmt.Printf("*********** Virtual Service %s\n", svc.Name)
//	if svc.Name == virtualservice {
//				fmt.Printf("Found service %s\n", svc.Name)
//				desiredSvc = svc
//		break
//	}
//}
//
//fmt.Printf("#### OLD %++v\n", desiredSvc)

//svcName := "test-vs"
//funcRoute := "test"
//clientLabel := "client1"
//gatewayHost := "kong-proxy.kong.svc.cluster.local"
//funcName := "blahblah"	//"function1"
//
////var desiredMatch v1alpha3.HTTPMatchRequest
//
//for _, svc := range virtualSvcList.Items {
//	fmt.Printf("*********** Virtual Service %s\n", svc.Name)
//	if svc.Name != svcName {
//		continue
//	}
//	vsSpec := svc.Spec
//	httproutes := vsSpec.Http
//	for _, httproute := range httproutes {
//		var sourceApp string
//		var destinationHost string
//		var functionVappName, functionName string
//		for _, route := range httproute.Route {
//			destinationHost = route.Destination.Host
//			//}
//		}
//		// Host
//		if gatewayHost != destinationHost {
//			continue
//		}
//		// serverless function url
//		if httproute.Rewrite != nil {
//			fmt.Printf("Rewrite uri %s\n", httproute.Rewrite.Uri)
//			functionName = httproute.Rewrite.Uri[1:]
//		}
//		if funcName != functionName {
//			continue
//		}
//
//		// client label
//		for _, match := range httproute.Match {
//			// find client label
//			fmt.Printf("source labels %s\n", match.SourceLabels)
//			sourceLabelMap := match.SourceLabels
//			for key, val := range sourceLabelMap {
//				if key == "app" {
//					sourceApp = val
//				}
//			}
//			if sourceApp != clientLabel {
//				continue
//			}
//			fmt.Printf("Found client label %s\n", sourceApp)
//			//
//			if match.Uri != nil {
//				functionVappName = match.Uri.Exact[1:]
//			}
//			if functionVappName == funcRoute {
//				//desiredMatch = match
//				break
//			}
//		}
//
//		fmt.Printf("Found client and route %s::%s-->%s\n", sourceApp, functionVappName, functionName)
//		desiredSvc = svc
//		break
//	}
//}

////var svcRoutes []v1alpha3.HTTPRoute
//for _, httproute := range httproutes {
//	fmt.Printf("%s: %d number of matches\n", httproute.Rewrite, len(httproute.Match))
//	curr_rewrite := httproute.Rewrite
//	functionName := httproute.Rewrite.Uri[1:]
//	//fmt.Printf("rewrite %s\n", functionName)
//	if funcName == functionName {
//		//fmt.Printf("	current rewrite uri %s\n", curr_rewrite.Uri)
//		curr_rewrite.Uri = "/" + "blah"
//		fmt.Printf("	changed current rewrite uri %s\n", curr_rewrite.Uri)
//		httproute.Rewrite = curr_rewrite
//		fmt.Printf("#### NEW %++v\n", httproute.Rewrite)
//	}
//	svcRoutes = append(svcRoutes, httproute)
//}
