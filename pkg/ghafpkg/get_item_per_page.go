package ghafpkg

// GetItemPerPage ItemPerPage 20 | 50 | 100 ....
func GetItemPerPage(paramRequestItemPerPage int, defaultItemPerPage int) int {
	var thisItemPerPage int
	if paramRequestItemPerPage == 0 {
		thisItemPerPage = defaultItemPerPage
	} else {
		thisItemPerPage = paramRequestItemPerPage
	}

	return thisItemPerPage
}
