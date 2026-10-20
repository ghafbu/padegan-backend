package messagepkg

import (
	"fmt"

	"github.com/ghafbu/padegan-backend/pkg/slugpkg"
)

func StatusCode500(lang *string, section string) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("status-code-500", section)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("متاسفانه سرور با مشکل مواجه شده است - %s", section)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("Internal Server Error - %s.", section)
	}

	return clone
}

func StatusCode500Service(
	paramUniqueID string,
	paramServiceName string,
	paramLayerName string,
	paramDescription string,
) Response {
	// Variables
	var clone = Response{}
	clone.Id = UtlBuildId("status-code-500", paramUniqueID)

	// Tag Log
	var thisLogPreset = "UniqueID:" + paramUniqueID + "Service:" + paramServiceName + " / Layer:" + paramLayerName + " / Description:" + paramDescription
	clone.Text = thisLogPreset

	return clone
}

func StatusCode500Unique(
	paramDescription string,
	paramUniqueID string,
) Response {
	// Variables
	var clone = Response{}
	clone.Id = UtlBuildId("status-code-500", paramUniqueID)

	thisDescription := slugpkg.Make(paramDescription, "_")
	// Tag Log
	var thisLogPreset = "UniqueID:" + paramUniqueID + " / description: " + thisDescription
	clone.Text = thisLogPreset

	return clone
}

func StatusCode404Unique(
	paramDescription string,
	paramUniqueID string,
) Response {
	// Variables
	var clone = Response{}
	clone.Id = UtlBuildId("status-code-404", paramUniqueID)

	thisDescription := slugpkg.Make(paramDescription, "_")
	// Tag Log
	var thisLogPreset = "UniqueID:" + paramUniqueID + " / description: " + thisDescription
	clone.Text = thisLogPreset

	return clone
}
