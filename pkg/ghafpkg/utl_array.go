package ghafpkg

func ArraySearch(paramArray []string, paramTarget string) bool {
	for _, item := range paramArray {
		if item == paramTarget {
			return true
		}
	}
	return false
}
