package istio

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/k8s"
	versionedclient "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/tools/clientcmd"
)

// NewIstioClientFromOptions initializes and returns an Istio client. The function will
// fail a tests if initialization gets an error.
func NewIstioClientFromOptions(t *testing.T, options *k8s.KubectlOptions) *versionedclient.Clientset {
	client, err := NewIstioClientFromOptionsE(t, options)
	if err != nil {
		t.Fatalf("Unable to initialize Istio client: %s", err)
	}
	return client
}

// NewIstioClientFromOptionsE initializes and returns an Istio client.
func NewIstioClientFromOptionsE(t *testing.T, options *k8s.KubectlOptions) (*versionedclient.Clientset, error) {
	configPath, err := options.GetConfigPath(t)
	if err != nil {
		return err
	}

	restConfig, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		return err
	}

	client, err := versionedclient.NewForConfig(restConfig)
	if err != nil {
		return err
	}

	return client, nil
}
