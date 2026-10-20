package pathpkg

import (
	"log"
	"path/filepath"
)

func PathCreateEnv(paramEnvMode string, paramServicePath string) string {
	if paramEnvMode == "development" {
		return filepath.Join(paramServicePath, "v1/deploy/development", ".env")
	}

	if paramEnvMode == "production" {
		return filepath.Join(paramServicePath, "v1/deploy/production", ".env")
	}

	log.Printf("Warning: Unsupported environment mode '%s'", paramEnvMode)

	return ""
}
