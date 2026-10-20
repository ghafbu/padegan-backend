package ghafpkg

import "math"

func GetTotalPages(paramTotalItems int64, paramItemPerPage int) int {
	var paperCountByFloat = float64(paramTotalItems) / float64(paramItemPerPage)
	var thisEnd = int(math.Ceil(paperCountByFloat))

	return thisEnd
}
