package model

import (
	"context"
	"time"

	"github.com/ghafbu/padegan-backend/api/dto/request"

	"github.com/ghafbu/padegan-backend/api/dto/response"
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
)

func (myModel *UserSettingModel) UserSettingsUpdate(
	_ context.Context,
	paramRequest *request.UserSettingsUpdate,
) (response.UserSettingsPublic, error) {
	var (
		thisResponse response.UserSettingsPublic
		adminSetting UserSetting
	)

	// DB Nil check
	if err := myModel.checkConnection(); err != nil {
		return thisResponse, err
	}

	thisFirst := myModel.postgres.Conn().Model(&UserSetting{}).FirstOrCreate(&adminSetting)

	if thisFirst.Error != nil {
		return response.UserSettingsPublic{}, thisFirst.Error
	}

	adminSetting.StatusSite = paramRequest.StatusSite
	adminSetting.StatusOrgSite = paramRequest.StatusOrgSite
	adminSetting.StatusSiteSignUp = paramRequest.StatusSiteSignUp
	adminSetting.StatusOrgPanel = paramRequest.StatusOrgPanel
	adminSetting.UpdatedAt = time.Now()

	thisFetch := myModel.postgres.Conn().
		Save(&adminSetting).
		Scan(&thisResponse)

	if thisFetch.Error != nil {
		return response.UserSettingsPublic{}, thisFetch.Error
	}

	if thisFetch.RowsAffected > 0 {
		ghafpkg.TimeGenerateStructFullDetails(thisResponse.CreatedAt, &thisResponse.CreatedAtDetails)
		ghafpkg.TimeGenerateStructFullDetails(thisResponse.UpdatedAt, &thisResponse.UpdatedAtDetails)
	}

	return thisResponse, nil
}
