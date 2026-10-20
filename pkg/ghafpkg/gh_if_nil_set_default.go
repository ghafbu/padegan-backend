package ghafpkg

func GHIfNilSetDefault[T any](paramData *T, paramDefault T) T {
	if paramData == nil {
		return paramDefault
	}
	return *paramData
}
