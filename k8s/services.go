package k8s

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
)

// GetServiceBySelector returns a kubernetes service matched by the given selector.
func GetServiceBySelector(t *testing.T, kubectlOptions *k8s.KubectlOptions, selector string) corev1.Service {
	services := k8s.ListServices(t, kubectlOptions, metav1.ListOptions{LabelSelector: selector})

	require.Len(t, services, 1, fmt.Sprintf("Service selector '%s' is ambiguous", selector))
	svc := services[0]

	k8s.WaitUntilServiceAvailable(t, kubectlOptions, svc.Name, 10, 5*time.Second)

	return svc
}
