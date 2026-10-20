package model

import (
	"context"
	"errors"
	"time"

	"github.com/ghafbu/padegan-backend/api/dto/request"

	"github.com/ghafbu/padegan-backend/api/dto/response"
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/ghafbu/padegan-backend/pkg/phonepkg"
	"gorm.io/gorm/clause"
)

func (myModel *UserModel) UsersCreate(
	paramCtx context.Context,
	paramRequest *request.UsersCreate,
	paramMobileDetails phonepkg.PhoneDetails,
	paramPhoneDetails phonepkg.PhoneDetails,
) (response.UsersPublic, error) {
	var thisResponse response.UsersPublic

	// DB Nil check
	if err := myModel.checkConnection(); err != nil {
		return thisResponse, err
	}

	// Hash password
	if paramRequest.Password != nil {
		var hashPassword, thisErrPass = ghafpkg.BcryptPassword(*paramRequest.Password)
		if thisErrPass != nil {
			return thisResponse, errors.New("password_hash_generate_error")
		}
		paramRequest.Password = &hashPassword
	}

	// Clone
	thisClone := User{
		Name:              paramRequest.Name,
		Family:            paramRequest.Family,
		Mobile:            paramMobileDetails.PhoneE164,         // ExistsBy
		MobileCountryCode: paramMobileDetails.CountryCodeString, // ExistsBy
		NationalCode:      paramRequest.NationalCode,            // ExistsBy
		Email:             paramRequest.Email,                   // ExistsBy
		Username:          paramRequest.Username,                // ExistsBy
		Password:          paramRequest.Password,
		Gender:            paramRequest.Gender,
		ProfileImage:      paramRequest.ProfileImage,
		Creator:           paramRequest.Creator,
		// Phone:               &paramPhoneDetails.PhoneE164,
		// PhoneCountryCode:    &paramPhoneDetails.CountryCodeString,
		// BirthLocationID:     paramRequest.BirthLocationID,
		// ResidenceLocationID: paramRequest.ResidenceLocationID,
		// BirthAddressID:      paramRequest.BirthAddressID,
		// CurrentAddressID:    paramRequest.CurrentAddressID,
		// OrgThemeMode:        OrgThemeModeSystem,
		// ProfileStatus:       ProfileStatusAvailable,
		// StatusUser:              Approved,
		// CreatedAt:         time.Time{},
		// UpdatedAt:         time.Time{},
	}

	// Phone
	if paramRequest.Phone != nil { // Check, because empty "" string save to DB
		thisClone.Phone = &paramPhoneDetails.PhoneE164
		thisClone.PhoneCountryCode = &paramPhoneDetails.CountryCodeString
	}

	// Check language or set default language
	if paramRequest.Language != nil {
		thisClone.Language = Language(*paramRequest.Language)
	} else {
		thisClone.Language = LanguagePersian
	}

	// Check org theme mode or set default org them mode
	if paramRequest.OrgThemeMode != nil {
		thisClone.OrgThemeMode = OrgThemeMode(*paramRequest.OrgThemeMode)
	} else {
		thisClone.OrgThemeMode = OrgThemeModeSystem
	}

	// Check profile status or set default profile status
	if paramRequest.ProfileStatus != nil {
		thisClone.ProfileStatus = ProfileStatus(*paramRequest.ProfileStatus)
	} else {
		thisClone.ProfileStatus = ProfileStatusAvailable
	}

	// Check status or set default status
	if paramRequest.Status != nil {
		thisClone.Status = StatusUser(*paramRequest.Status)
	} else {
		thisClone.Status = Approved
	}

	// Check has password
	thisClone.HasPassword = func() bool {
		if paramRequest.Password != nil {
			return true
		}

		return false
	}()

	// Birthdate
	if paramRequest.BirthDate != nil {
		var thisBirthDate, err = ghafpkg.TimeConvertStringToTimeByFormat(*paramRequest.BirthDate, time.DateOnly)
		if err != nil {
			return thisResponse, errors.New("date_parse_error")
		}

		thisClone.Birthdate = &thisBirthDate
	}

	// DB action
	thisFetch := myModel.postgres.Conn().
		Clauses(clause.Returning{}).
		Create(&thisClone).
		Scan(&thisResponse)

	// DB error
	if thisFetch.Error != nil {
		return thisResponse, thisFetch.Error
	}

	if thisResponse.Birthdate != nil {
		ghafpkg.TimeGenerateStructDetailsWithPointer(thisClone.Birthdate, &thisResponse.BirthdateDetails)
	}

	ghafpkg.TimeGenerateStructFullDetails(thisClone.CreatedAt, &thisResponse.CreatedAtDetails)
	ghafpkg.TimeGenerateStructFullDetails(thisClone.UpdatedAt, &thisResponse.UpdatedAtDetails)

	return thisResponse, nil
}
