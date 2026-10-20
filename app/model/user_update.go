package model

import (
	"context"
	"errors"
	"time"

	"github.com/ghafbu/padegan-backend/lib/consts"
	modelerrors "github.com/ghafbu/padegan-backend/lib/consts"

	"github.com/ghafbu/padegan-backend/api/dto/params"
	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/api/dto/response"
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/ghafbu/padegan-backend/pkg/phonepkg"
	"gorm.io/gorm/clause"
)

func (myModel *UserModel) UsersUpdate(
	paramCtx context.Context,
	paramReq *request.UsersUpdate,
	paramPhoneDetails phonepkg.PhoneDetails,
	paramParams params.UsersUpdate,
) (response.UsersPublic, int64, error) {
	var thisResponse response.UsersPublic

	// DB Nil check
	if err := myModel.checkConnection(); err != nil {
		return thisResponse, 0, err
	}

	thisClone := map[string]interface{}{
		"name":   paramReq.Name,
		"family": paramReq.Family,
		// "mobile":              paramReq.Mobile,
		// "mobile_country_code": paramMobileDetails.PhoneE164,
		"email":          paramReq.Email,
		"national_code":  paramReq.NationalCode,
		"username":       paramReq.Username,
		"gender":         paramReq.Gender,
		"profile_image":  paramReq.ProfileImage,
		"org_theme_mode": paramReq.OrgThemeMode,
		"language":       paramReq.Language,
		"profile_status": paramReq.ProfileStatus,
		"status":         paramReq.Status,
		"updated_at":     time.Now(),
	}

	if paramReq.BirthDate != nil {
		var thisBirthdate, err = ghafpkg.TimeConvertStringToTimeByFormat(*paramReq.BirthDate, time.DateOnly)
		if err != nil {
			return thisResponse, 0, errors.New(consts.ErrorKeyModelBirthdatePars)
		}
		thisClone["birthdate"] = thisBirthdate
	}

	if paramReq.Phone != nil {
		thisClone["phone"] = paramPhoneDetails.PhoneE164
		thisClone["phone_country_code"] = paramPhoneDetails.CountryCodeString
	} else {
		thisClone["phone"] = nil
		thisClone["phone_country_code"] = nil
	}

	thisFetch := myModel.postgres.Conn().Model(&thisClone).
		Where("id = ?", paramParams.Id).
		Clauses(clause.Returning{}).
		Updates(thisClone).
		Scan(&thisResponse)

	if thisFetch.RowsAffected == 0 {
		return thisResponse, 0, errors.New(modelerrors.ErrorKeyModelRecordNotFound)
	}

	if thisFetch.Error != nil {
		return thisResponse, thisFetch.RowsAffected, thisFetch.Error
	}

	if thisResponse.Birthdate != nil {
		ghafpkg.TimeGenerateStructDetailsWithPointer(thisResponse.Birthdate, &thisResponse.BirthdateDetails)
	}

	ghafpkg.TimeGenerateStructFullDetails(thisResponse.CreatedAt, &thisResponse.CreatedAtDetails)
	ghafpkg.TimeGenerateStructFullDetails(thisResponse.UpdatedAt, &thisResponse.UpdatedAtDetails)

	return thisResponse, thisFetch.RowsAffected, nil
}
