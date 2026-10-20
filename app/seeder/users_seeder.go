package seeder

import (
	"fmt"
	"time"

	"github.com/ghafbu/padegan-backend/pkg/databasepkg/gorm"
	"github.com/ghafbu/padegan-backend/pkg/databasepkg/gorm/faker"

	"github.com/ghafbu/padegan-backend/app/model"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
)

//func UsersSeedersCreate(paramCount uint64) ([]model.User, error) {
//	var thisSeeders []model.User
//
//	var i uint64
//	for i = 1; i <= paramCount; i++ {
//		//thisRandSource := rand.NewSource(time.Now().UnixNano())
//		//thisRandGenerator := rand.New(thisRandSource)
//
//		thisCreateAt, err := ghafpkg.GHRandomDate(
//			time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
//			time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
//		)
//
//		if err != nil {
//			return nil, fmt.Errorf("error generating created-at date for index [%d]: %w", i, err)
//		}
//
//		thisSeeders = append(thisSeeders, model.User{
//			CreatedAt: thisCreateAt,
//			UpdatedAt: time.Time{},
//		})
//	}
//
//	return thisSeeders, nil
//}

func UsersSeeders(paramCount uint64, paramPostgres gorm.DB) error {
	var thisSeeders []model.User

	var i uint64
	for i = 1; i <= paramCount; i++ {
		//thisRandSource := rand.NewSource(time.Now().UnixNano())
		//thisRandGenerator := rand.New(thisRandSource)

		thisCreateAt, err := ghafpkg.GHRandomDate(
			time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		)

		if err != nil {
			return fmt.Errorf("error generating created-at date for index [%d]: %w", i, err)
		}

		thisSeeders = append(thisSeeders, model.User{
			Name:              "محمد",
			Family:            "قاسمی",
			Mobile:            "+989112203390",
			MobileCountryCode: "+98",

			CreatedAt: thisCreateAt,
			UpdatedAt: time.Time{},
		})
	}

	var DatabaseContentCount int64
	if err := paramPostgres.Conn().
		Model(model.User{}).
		Count(&DatabaseContentCount).Error; err != nil {
		return err
	}

	if DatabaseContentCount == 0 {
		// Init Faker
		thisFaker := faker.Init(paramPostgres)
		// Faker - drop
		if err := thisFaker.TruncateWithModel(&model.User{}); err != nil {
			return err
		}

		for _, thisSeeder := range thisSeeders {
			if _, err := thisFaker.Create(&thisSeeder); err != nil {
				return err
			}
		}
	}

	return nil
}
