package pathpkg

import (
	"log"
	"path/filepath"
)

func PathCreateGlobalEnv(paramPathRoot, paramEnvMode string) string {
	if paramEnvMode == "development" {
		return filepath.Join(paramPathRoot, "deploy/development", ".env")
	}

	if paramEnvMode == "production" {
		return filepath.Join(paramPathRoot, "deploy/production", ".env")
	}

	log.Printf("Warning: Unsupported environment mode '%s'", paramEnvMode)

	return ""
}
