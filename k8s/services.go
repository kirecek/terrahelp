package k8s

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetServiceBySelector returns a kubernetes service matched by the given selector.
func GetServiceBySelector(t *testing.T, options *k8s.KubectlOptions, selector string) corev1.Service {
	services := k8s.ListServices(t, options, metav1.ListOptions{LabelSelector: selector})

	require.Len(t, services, 1, fmt.Sprintf("Service selector '%s' is ambiguous", selector))
	svc := services[0]

	k8s.WaitUntilServiceAvailable(t, options, svc.Name, 10, 5*time.Second)

	return svc
}
