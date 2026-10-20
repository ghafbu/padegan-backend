package pathpkg

import "path/filepath"

func PathCreateService(paramRootPath string, paramServiceName string) string {
	return filepath.Join(paramRootPath, "internal", paramServiceName)
}
