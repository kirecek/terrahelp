package terraform

import (
	"fmt"
	"os"
	"testing"
)

// GitlabHTTPBackendConfig is Go copy of Gitlab's wrapper for terraform binary.
//
// The function returns BackendConfig map which can be passed to terraform.Options.BackendConfig
// in terratest package.
//
// https://gitlab.com/gitlab-org/terraform-images/-/blob/master/src/bin/gitlab-terraform.sh
// https://docs.gitlab.com/ee/user/infrastructure/terraform_state.html
func GitlabHTTPBackendConfig(t *testing.T) map[string]interface{} {
	username := getEnv("TF_USERNAME", os.Getenv("GITLAB_USER_LOGIN"))
	password := getEnv("TF_HTTP_PASSWORD", os.Getenv("TF_PASSWORD"))

	address := getEnv("TF_HTTP_ADDRESS", os.Getenv("TF_ADDRESS"))

	if password == "" {
		t.Log("TF_PASSWORD nor TF_HTTP_PASSWORD not set.")
		t.Log("\t - using default CI credentials 'gitlab-ci-token/$CI_JOB_TOKEN'.")

		username = "gitlab-ci-token"
		password = os.Getenv("CI_JOB_TOKEN")
	}

	if address == "" {
		stateName, ok := os.LookupEnv("TF_STATE_NAME")
		if !ok {
			t.Fatal("Clould not configure terraform gitlab backend - one of [TF_HTTP_ADDRESS, TF_ADDRESS, TF_STATE_NAME] env variable must be provided.")
		}

		address = fmt.Sprintf(
			"%s/projects/%s/terraform/state/%s",
			os.Getenv("CI_API_V4_URL"),
			os.Getenv("CI_PROJECT_ID"),
			stateName,
		)
	}

	return map[string]interface{}{
		"address":        address,
		"password":       password,
		"username":       getEnv("TF_HTTP_USERNAME", username),
		"lock_address":   getEnv("TF_HTTP_LOCK_ADDRESS", fmt.Sprintf("%s/lock", address)),
		"lock_method":    getEnv("TF_HTTP_LOCK_METHOD", "POST"),
		"unlock_address": getEnv("TF_HTTP_UNLOCK_ADDRESS", fmt.Sprintf("%s/lock", address)),
		"unlock_method":  getEnv("TF_HTTP_UNLOCK_METHOD", "DELETE"),
		"retry_wait_min": getEnv("TF_HTTP_RETRY_WAIT_MIN", "5"),
	}
}
