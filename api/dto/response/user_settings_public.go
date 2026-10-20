package response

import (
	"time"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
)

type UserSettingsPublic struct {
	StatusSite       bool                        `json:"status_site"`         //وضعیت سایت
	StatusOrgPanel   bool                        `json:"status_org_panel"`    //وضعیت پنل سازمان ها
	StatusSiteSignUp bool                        `json:"status_site_sign_up"` //وضعیت ثبت نام
	StatusOrgSite    bool                        `json:"status_org_site"`     //وضعیت سایت سازمان ها
	CreatedAt        time.Time                   `json:"-"`
	CreatedAtDetails ghafpkg.DateDetailsWithTime `json:"created_at_details" gorm:"-"`
	UpdatedAt        time.Time                   `json:"-"`
	UpdatedAtDetails ghafpkg.DateDetailsWithTime `json:"updated_at_details" gorm:"-"`
	DeletedAt        *string                     `gorm:"index"`
}
