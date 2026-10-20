package ghafpkg

import "strings"

func FetchMakerOrderMethod(paramTarget *string, paramOrderByDefault string) {
	//order
	//if *paramTarget == "" {
	//	*paramTarget = "desc"
	//}

	*paramTarget = strings.ToLower(*paramTarget)
	switch *paramTarget {
	case "asc":
	case "desc":
	case "":
		if paramOrderByDefault == "sort" {
			*paramTarget = "desc"
		} else {
			// id
			*paramTarget = "desc"
		}
	case "1":
		*paramTarget = "asc"
	case "-1":
		*paramTarget = "desc"
	case "ascending":
		*paramTarget = "asc"
	case "descending":
		*paramTarget = "desc"
	default:
		*paramTarget = "desc"
	}
}
