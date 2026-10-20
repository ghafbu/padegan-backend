package ghafpkg

func FetchMakerOrder(paramOrderBy *string, paramOrderMethod *string, paramOrderAccess *[]string, paramOrderByDefault string) {
	//orderBy
	FetchMakerOrderBy(paramOrderBy, paramOrderAccess, paramOrderByDefault)
	//orderMethod
	FetchMakerOrderMethod(paramOrderMethod, paramOrderByDefault)
}
