/*

LICENSE

*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"net/http"

	v1alpha1 "github.com/portworx/kdmp/pkg/apis/kdmp/v1alpha1"
	"github.com/portworx/kdmp/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type KdmpV1alpha1Interface interface {
	RESTClient() rest.Interface
	BackupLocationMaintenancesGetter
	DataExportsGetter
	ResourceBackupsGetter
	ResourceExportsGetter
	VolumeBackupsGetter
	VolumeBackupDeletesGetter
}

// KdmpV1alpha1Client is used to interact with features provided by the kdmp.portworx.com group.
type KdmpV1alpha1Client struct {
	restClient rest.Interface
}

func (c *KdmpV1alpha1Client) BackupLocationMaintenances(namespace string) BackupLocationMaintenanceInterface {
	return newBackupLocationMaintenances(c, namespace)
}

func (c *KdmpV1alpha1Client) DataExports(namespace string) DataExportInterface {
	return newDataExports(c, namespace)
}

func (c *KdmpV1alpha1Client) ResourceBackups(namespace string) ResourceBackupInterface {
	return newResourceBackups(c, namespace)
}

func (c *KdmpV1alpha1Client) ResourceExports(namespace string) ResourceExportInterface {
	return newResourceExports(c, namespace)
}

func (c *KdmpV1alpha1Client) VolumeBackups(namespace string) VolumeBackupInterface {
	return newVolumeBackups(c, namespace)
}

func (c *KdmpV1alpha1Client) VolumeBackupDeletes(namespace string) VolumeBackupDeleteInterface {
	return newVolumeBackupDeletes(c, namespace)
}

// NewForConfig creates a new KdmpV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*KdmpV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new KdmpV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*KdmpV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &KdmpV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new KdmpV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *KdmpV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new KdmpV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *KdmpV1alpha1Client {
	return &KdmpV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *KdmpV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
