package terraform_provider

import "os"

func getEnvWithDefault(name, def string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return def
}
