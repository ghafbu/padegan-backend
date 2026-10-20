package model

import (
	"gorm.io/gorm"
	"time"
)

type OrgThemeMode uint

const (
	OrgThemeModeSystem OrgThemeMode = iota + 1
	OrgThemeModeLight
	OrgThemeModeDark
)

type Language uint

const (
	LanguagePersian Language = iota + 1
	LanguageEnglish
)

type StatusUser uint

const (
	Approved StatusUser = iota + 1
	Pending
	Suspended
)

type ProfileStatus uint

const (
	ProfileStatusAvailable ProfileStatus = iota + 1
	ProfileStatusBusy
)

type User struct {
	ID                uint64         `gorm:"type:bigint;primaryKey;autoIncrement"`
	Name              string         `gorm:"type:varchar(255);not null;index"`
	Family            string         `gorm:"type:varchar(255);not null;index"`
	Mobile            string         `gorm:"type:varchar(15);not null;unique;index"` // +989112253390 E164 FORMAT
	MobileCountryCode string         `gorm:"type:varchar(7);not null;index"`
	Email             *string        `gorm:"type:varchar(255);null;unique;index"`
	IsEmailVerified   bool           `gorm:"type:boolean;not null;default:false"`
	EmailVerifiedAt   *time.Time     `gorm:"type:timestamp;null"`
	Username          *string        `gorm:"type:varchar(255);null;unique;index"`
	Password          *string        `gorm:"type:varchar(1024);null"`
	HasPassword       bool           `gorm:"type:boolean;not null;default:false"`
	Salt              *[]byte        `gorm:"type:bytea;null"`
	Phone             *string        `gorm:"type:varchar(15);null;index"` // +981144342 E164 FORMAT
	PhoneCountryCode  *string        `gorm:"type:varchar(7);null;index"`
	NationalCode      *string        `gorm:"type:varchar(10);null;unique;index"`
	Birthdate         *time.Time     `gorm:"type:date;null;index"`
	Gender            uint           `gorm:"type:integer;not null;index"` // 1=male & 2=female
	Language          Language       `gorm:"type:integer;not null;index"` // 1=persian & 2=english ,....
	OrgThemeMode      OrgThemeMode   `gorm:"type:integer;not null;index"` // 1=system & 2=light & 3=dark & 4=blue & etc..
	ProfileStatus     ProfileStatus  `gorm:"type:integer;not null;index"` // 1=Available & 2=Busy
	ProfileImage      *string        `gorm:"type:varchar(1024);null;index"`
	Status            StatusUser     `gorm:"type:integer;not null;index"` // 1=pending & 2=approved & 3=rejected
	Creator           uint64         `gorm:"bigint,not null;index'"`
	CreatedAt         time.Time      `gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
