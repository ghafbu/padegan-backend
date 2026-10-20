package model

import (
	"context"

	"github.com/ghafbu/padegan-backend/api/dto/response"
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
)

func (myModel *UserSettingModel) UserSettingsGet(_ context.Context) (response.UserSettingsPublic, int64, error) {
	var (
		thisSetting  UserSetting
		thisResponse response.UserSettingsPublic
	)

	// DB Nil check
	if err := myModel.checkConnection(); err != nil {
		return thisResponse, 0, err
	}

	// fetch
	thisFetch := myModel.postgres.Conn().Model(&UserSetting{})

	// find by
	// thisId, err := strconv.Atoi(paramId)
	// if err != nil || thisId == 0 {
	// 	return thisResponse, 0, errors.New(consts.UserSettingsErrorKeyModelIDBadRequest)
	// }
	// thisFetch.Where("id=?", thisId).Find(&thisSetting)

	thisFetch.Find(&thisSetting)

	if thisFetch.Error != nil {
		return thisResponse, thisFetch.RowsAffected, thisFetch.Error
	}

	if thisFetch.RowsAffected > 0 {
		thisResponse.StatusSite = thisSetting.StatusSite
		thisResponse.StatusOrgPanel = thisSetting.StatusOrgPanel
		thisResponse.StatusSiteSignUp = thisSetting.StatusSiteSignUp
		thisResponse.StatusOrgSite = thisSetting.StatusOrgSite
		thisResponse.CreatedAt = thisSetting.CreatedAt
		thisResponse.UpdatedAt = thisSetting.UpdatedAt
		ghafpkg.TimeGenerateStructFullDetails(thisSetting.CreatedAt, &thisResponse.CreatedAtDetails)
		ghafpkg.TimeGenerateStructFullDetails(thisSetting.UpdatedAt, &thisResponse.UpdatedAtDetails)
	}

	return thisResponse, thisFetch.RowsAffected, nil
}
