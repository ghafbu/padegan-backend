package ghafpkg

func FetchMakerOrderBy(paramOrderBy *string, paramOrderAccess *[]string, paramOrderByDefault string) {
	if paramOrderByDefault == "" {
		paramOrderByDefault = "id"
	}
	if !ArraySearch(*paramOrderAccess, *paramOrderBy) {
		*paramOrderBy = paramOrderByDefault
	}
}
