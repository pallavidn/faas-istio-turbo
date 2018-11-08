package discovery

import (
	"github.com/pallavidn/faas-istio/pkg/apis/istio"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
	//v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//types "k8s.io/apimachinery/pkg/types"
	//watch "k8s.io/apimachinery/pkg/watch"
	//"github.com/istio/api/networking/v1alpha3"
)

type IstioRestClient struct {
	restClient rest.Interface
}

// virtualServices implements VirtualServiceInterface
type virtualServices struct {
	client rest.Interface
	ns     string
}

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: istio.GroupName, Version: "v1alpha3"}
var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var ParameterCodec = runtime.NewParameterCodec(Scheme)
//
//
//// VirtualServiceInterface has methods to work with VirtualService resources.
//type VirtualServiceInterface interface {
//	Create(*v1alpha3.VirtualService) (*v1alpha3.VirtualService, error)
//	Update(*v1alpha3.VirtualService) (*v1alpha3.VirtualService, error)
//	Delete(name string, options *v1.DeleteOptions) error
//	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
//	Get(name string, options v1.GetOptions) (*v1alpha3.VirtualService, error)
//	List(opts v1.ListOptions) (*v1alpha3.VirtualServiceList, error)
//	Watch(opts v1.ListOptions) (watch.Interface, error)
//	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha3.VirtualService, err error)
//}
//
//
//// Get takes name of the virtualService, and returns the corresponding virtualService object, and an error if there is any.
//func (c *virtualServices) Get(name string, options v1.GetOptions) (result *v1alpha3.VirtualService, err error) {
//	result = &v1alpha3.VirtualService{}
//	err = c.client.Get().
//		Namespace(c.ns).
//		Resource("virtualservices").
//		Name(name).
//		VersionedParams(&options, ParameterCodec).
//		Do().
//		Into(result)
//	return
//}

// newVirtualServices returns a VirtualServices
func newVirtualServices(c *IstioRestClient, namespace string) *virtualServices {
	return &virtualServices{
		client: c.restClient, //RESTClient(),
		ns:     namespace,
	}
}

// NewForConfig creates a new NetworkingV1alpha3Client for the given config.
func NewForConfig(c *rest.Config) (*IstioRestClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &IstioRestClient{client}, nil
}

func setConfigDefaults(config *rest.Config) error {
	gv := SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}
