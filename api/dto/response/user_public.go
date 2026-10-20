package response

import (
	"time"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
)

type UsersPublic struct {
	ID                uint64                      `json:"id"`
	Name              string                      `json:"name"`
	Family            string                      `json:"family"`
	Mobile            string                      `json:"mobile"`
	MobileCountryCode string                      `json:"mobile_country_code"`
	Phone             *string                     `json:"phone"`
	PhoneCountryCode  *string                     `json:"phone_country_code"`
	Birthdate         *time.Time                  `json:"-"`
	BirthdateDetails  ghafpkg.DateDetails         `json:"birth_date_details" gorm:"-"`
	NationalCode      *string                     `json:"national_code"`
	Email             *string                     `json:"email"`
	IsEmailVerified   bool                        `json:"is_email_verified"`
	EmailVerifiedAt   *time.Time                  `json:"email_verified_at"`
	Username          *string                     `json:"username"`
	HasPassword       bool                        `json:"has_password"`
	Gender            uint                        `json:"gender"`
	ProfileImage      *string                     `json:"profile_image"`
	OrgThemeMode      uint                        `json:"org_theme_mode"`
	Language          uint                        `json:"language"`
	ProfileStatus     uint                        `json:"profile_status"`
	Status            uint                        `json:"status"`
	Creator           uint64                      `gorm:"creator"`
	CreatedAt         time.Time                   `json:"-"`
	CreatedAtDetails  ghafpkg.DateDetailsWithTime `json:"created_at_details" gorm:"-"`
	UpdatedAt         time.Time                   `json:"-"`
	UpdatedAtDetails  ghafpkg.DateDetailsWithTime `json:"updated_at_details" gorm:"-"`
	// DeletedAt         *string                     `gorm:"index" json:"deleted_at"`
	// BirthAddressID      *uint64                     `json:"birth_address_id"`
	// BirthLocationID     *uint64                     `json:"birth_location_id"`
	// CurrentAddressID    *uint64                     `json:"current_address_id"`
	// ResidenceLocationID *uint64                     `json:"residence_location_id"`
}
