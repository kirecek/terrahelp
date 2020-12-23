package k8s

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// CreateKubeConfigFile creates a KUBECONFIG file with a given API configuration. This will fail a test
// if there is an error in creating the config file.
func CreateKubeConfigFile(t *testing.T, contextName, configPath, namespace, clusterName, server, caCertificate, token string) {
	err := CreateKubeConfigFileE(contextName, configPath, namespace, clusterName, server, caCertificate, token)
	assert.NoError(t, err)
}

// CreateKubeConfigFileE creates a KUBECONFIG file with a given API configuration.
func CreateKubeConfigFileE(contextName, configPath, namespace, clusterName, server, caCertificate, token string) error {
	if namespace == "" {
		namespace = "default"
	}

	clientConfig := clientcmdapi.NewConfig()

	clientConfig.Clusters[clusterName] = &clientcmdapi.Cluster{
		Server:                   server,
		CertificateAuthorityData: []byte(caCertificate),
	}

	clientConfig.Contexts[contextName] = &clientcmdapi.Context{
		Cluster:   clusterName,
		Namespace: namespace,
		AuthInfo:  namespace,
	}

	clientConfig.AuthInfos[namespace] = &clientcmdapi.AuthInfo{
		Token: token,
	}

	clientConfig.CurrentContext = contextName

	return clientcmd.WriteToFile(*clientConfig, configPath)
}
