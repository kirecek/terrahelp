package k8s

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetPodBySelector returns Pod object specified by selector string. The getter return only the first matched pod
// and always waits till pod is available.
func GetPodBySelector(t *testing.T, options *k8s.KubectlOptions, selector string) corev1.Pod {
	var pods []corev1.Pod
	retry.DoWithRetry(t, fmt.Sprintf("Waiting for pod matching filter '%s'", selector), 5, 5*time.Second, func() (string, error) {
		var err error
		pods, err = k8s.ListPodsE(t, options, metav1.ListOptions{LabelSelector: selector})
		if err != nil {
			return "", err
		}
		if len(pods) == 0 {
			return "", fmt.Errorf("Pod matching filter '%s' not yet created", selector)
		}
		return "", nil
	})

	require.NotEmpty(t, pods, fmt.Sprintf("Could not find any pod by the given label selector: %s", selector))

	pod := pods[0]

	image := pods.Spec.Containers[0].Image
	for _, p := range pods {
		require.Equal(t, p.Spec.Containers[0].Image, image, fmt.Sprintf("Ambiguous slector %s. Slector must target pods only from one deploy/sts/ds.", selector))
	}

	k8s.WaitUntilPodAvailable(t, options, pod.Name, 10, 5*time.Second)

	return pod
}
