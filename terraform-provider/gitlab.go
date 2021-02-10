package terraform_provider

import (
	"fmt"
	"os"
)

// GitlabBackendConfig is Go copy of Gitlab's wrapper for terraform binary.
//
// The function returns BackendConfig map which can be passed to terraform.Options.BackendConfig
// in terratest package.
//
// https://gitlab.com/gitlab-org/terraform-images/-/blob/master/src/bin/gitlab-terraform.sh
// https://docs.gitlab.com/ee/user/infrastructure/terraform_state.html
func GitlabHTTPBackendConfig() map[string]interface{} {
	username := getEnvWithDefault("TF_USERNAME", os.Getenv("GITLAB_USER_LOGIN"))
	password := os.Getenv("TF_PASSWORD")
	address := os.Getenv("TF_ADDRESS")

	// If TF_PASSWORD is unset then default to gitlab-ci-token/CI_JOB_TOKEN
	if password == "" {
		username = "gitlab-ci-token"
		password = os.Getenv("CI_JOB_TOKEN")
	}

	// let's build address from CI variables if address is not provided
	if address == "" {
		address = fmt.Sprintf(
			"%s/projects/%s/terraform/state/%s",
			os.Getenv("CI_API_V4_URL"),
			os.Getenv("CI_PROJECT_ID"),
			getEnvWithDefault("TF_STATE_NAME", "terratest"),
		)
	}

	// return backend config
	return map[string]interface{}{
		"address":        getEnvWithDefault("TF_HTTP_ADDRESS", address),
		"lock_address":   getEnvWithDefault("TF_HTTP_LOCK_ADDRESS", fmt.Sprintf("%s/lock", address)),
		"lock_method":    getEnvWithDefault("TF_HTTP_LOCK_METHOD", "POST"),
		"unlock_address": getEnvWithDefault("TF_HTTP_UNLOCK_ADDRESS", fmt.Sprintf("%s/lock", address)),
		"unlock_method":  getEnvWithDefault("TF_HTTP_UNLOCK_METHOD", "DELETE"),
		"username":       getEnvWithDefault("TF_HTTP_USERNAME", username),
		"password":       getEnvWithDefault("TF_HTTP_PASSWORD", password),
		"retry_wait_min": getEnvWithDefault("TF_HTTP_RETRY_WAIT_MIN", "5"),
	}
}
