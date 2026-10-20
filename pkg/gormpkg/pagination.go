package gormpkg

import (
	"gorm.io/gorm"
)

func Paginate(paramPage int, paramItemPerPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		//// Deprecated Logic ---------
		//// Use [GHIfNilOrDefaultTypeSetDefaultForINT] before using [Paginate]
		//// EX: thisPage = ghafpkg.GHIfNilOrDefaultTypeSetDefaultForINT(paramRequest.Page, 1)
		//if *paramPage <= 0 {
		//	*paramPage = 1
		//}
		//
		//// Deprecated Logic ---------
		//// Use [GHIfNilOrDefaultTypeSetDefaultForINT] before using [Paginate]
		//// EX: thisItemPerPage = ghafpkg.GHIfNilOrDefaultTypeSetDefaultForINT(paramRequest.ItemPerPage, consts.ItemPerPage10)
		//if paramItemPerPage == 0 {
		//	paramItemPerPage = ghafpkg.FetchItemPerPage
		//}

		offset := (paramPage - 1) * paramItemPerPage
		return db.Offset(offset).Limit(paramItemPerPage)
	}
}
