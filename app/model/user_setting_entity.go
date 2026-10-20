package model

import (
	"time"

	"gorm.io/gorm"
)

type UserSetting struct {
	ID               uint64         `gorm:"type:bigint;primaryKey;autoIncrement"`
	StatusSite       bool           `gorm:"type:boolean;not null"` // وضعیت سایت
	StatusOrgPanel   bool           `gorm:"type:boolean;not null"` // وضعیت پنل سازمان ها
	StatusSiteSignUp bool           `gorm:"type:boolean;not null"` // وضعیت ثبت نام
	StatusOrgSite    bool           `gorm:"type:boolean;not null"` // وضعیت سایت سازمان ها
	CreatedAt        time.Time      `gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
