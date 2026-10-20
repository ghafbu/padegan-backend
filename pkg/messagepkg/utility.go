package messagepkg

import (
	"fmt"

	"github.com/ghafbu/padegan-backend/pkg/slugpkg"
)

func UtlBuildId(paramKey string, paramField string) string {
	return slugpkg.Make(
		fmt.Sprintf("%s_%s", paramKey, paramField),
		"_",
	)
}

func UtlBuildIdText(paramText string) string {
	return slugpkg.Make(paramText, "_")
}
