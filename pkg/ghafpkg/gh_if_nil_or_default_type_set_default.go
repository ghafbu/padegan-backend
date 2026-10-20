package ghafpkg

func GHIfNilOrDefaultTypeSetDefaultForINT(paramData *int, paramDefault int) int {
	if paramData == nil || *paramData == 0 {
		return paramDefault
	}
	return *paramData
}

func GHIfNilOrDefaultTypeSetDefaultForUINT(paramData *uint, paramDefault uint) uint {
	if paramData == nil || *paramData == 0 {
		return paramDefault
	}
	return *paramData
}

func GHIfNilOrDefaultTypeSetDefaultForUINT64(paramData *uint64, paramDefault uint64) uint64 {
	if paramData == nil || *paramData == 0 {
		return paramDefault
	}
	return *paramData
}

func GHIfNilOrDefaultTypeSetDefaultForSTRING(paramData *string, paramDefault string) string {
	if paramData == nil || *paramData == "" {
		return paramDefault
	}
	return *paramData
}
