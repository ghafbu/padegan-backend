package model

import (
	"context"
	"errors"

	"github.com/ghafbu/padegan-backend/lib/consts"

	modelerrors "github.com/ghafbu/padegan-backend/lib/consts"

	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/api/dto/response"
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
)

func (myModel *UserModel) UsersBatchGet(
	paramCtx context.Context,
	paramRequest *request.UsersBatchGet,
) ([]response.UsersPublic, int64, error) {
	var (
		thisPublicResponse           []response.UsersPublic
		thisUsers                    []User
		thisOrderBy, thisOrderMethod string
	)

	// DB Nil check
	if err := myModel.checkConnection(); err != nil {
		return thisPublicResponse, 0, err
	}

	// fetch
	thisFetch := myModel.postgres.Conn().Model(&User{})

	// sort
	thisOrderBy = ghafpkg.GHIfNilSetDefault[string](&paramRequest.OrderBy, consts.DefaultModelOrderByID)
	thisOrderMethod = ghafpkg.GHIfNilSetDefault[string](&paramRequest.OrderMethod,
		consts.DefaultModelOrderMethodASC)

	// Set Order
	thisFetch.Order(thisOrderBy + " " + thisOrderMethod)

	// find by
	if paramRequest.Mobiles != nil {
		thisFetch.Where("mobile in (?)", *paramRequest.Mobiles)
	}

	if paramRequest.Phones != nil {
		thisFetch.Where("phone in (?)", *paramRequest.Phones)
	}

	if paramRequest.NationalCodes != nil {
		thisFetch.Where("national_code in (?)", *paramRequest.NationalCodes)
	}

	if paramRequest.Emails != nil {
		thisFetch.Where("email in (?)", *paramRequest.Emails)
	}

	if paramRequest.Usernames != nil {
		thisFetch.Where("username in (?)", *paramRequest.Usernames)
	}

	if paramRequest.Statuses != nil {
		thisFetch.Where("status in (?)", *paramRequest.Statuses)
	}

	// the end finding
	if paramRequest.Ids != nil {
		thisFetch.
			Find(&thisUsers, paramRequest.Ids)
	} else {
		thisFetch.
			Find(&thisUsers)
	}

	if thisFetch.RowsAffected == 0 {
		return thisPublicResponse, 0, errors.New(modelerrors.ErrorKeyModelRecordNotFound)
	}

	if thisFetch.Error != nil {
		return thisPublicResponse, thisFetch.RowsAffected, thisFetch.Error
	}

	for _, adminMember := range thisUsers {
		thisUserPublic := response.UsersPublic{
			ID:                adminMember.ID,
			Name:              adminMember.Name,
			Family:            adminMember.Family,
			Mobile:            adminMember.Mobile,
			MobileCountryCode: adminMember.MobileCountryCode,
			HasPassword:       adminMember.HasPassword,
			Gender:            adminMember.Gender,
			IsEmailVerified:   adminMember.IsEmailVerified,
			OrgThemeMode:      uint(adminMember.OrgThemeMode),
			Language:          uint(adminMember.Language),
			ProfileStatus:     uint(adminMember.ProfileStatus),
			Status:            uint(adminMember.Status),
			Creator:           adminMember.Creator,
			CreatedAt:         adminMember.CreatedAt,
			UpdatedAt:         adminMember.UpdatedAt,
			// BirthAddressID:      nil,
			// BirthLocationID:     nil,
			// CurrentAddressID:    nil,
			// ResidenceLocationID: nil,
			// DeletedAt:           nil,
		}
		if adminMember.Phone != nil {
			thisUserPublic.Phone = adminMember.Phone
		}
		if adminMember.PhoneCountryCode != nil {
			thisUserPublic.PhoneCountryCode = adminMember.PhoneCountryCode
		}
		if adminMember.Birthdate != nil {
			thisUserPublic.Birthdate = adminMember.Birthdate
			ghafpkg.TimeGenerateStructDetailsWithPointer(
				adminMember.Birthdate,
				&thisUserPublic.BirthdateDetails,
			)
		}
		if adminMember.NationalCode != nil {
			thisUserPublic.NationalCode = adminMember.NationalCode
		}

		if adminMember.Email != nil {
			thisUserPublic.Email = adminMember.Email
		}
		if adminMember.EmailVerifiedAt != nil {
			thisUserPublic.EmailVerifiedAt = adminMember.EmailVerifiedAt
		}
		if adminMember.Username != nil {
			thisUserPublic.Username = adminMember.Username
		}
		if adminMember.ProfileImage != nil {
			thisUserPublic.ProfileImage = adminMember.ProfileImage
		}

		ghafpkg.TimeGenerateStructFullDetails(adminMember.CreatedAt, &thisUserPublic.CreatedAtDetails)
		ghafpkg.TimeGenerateStructFullDetails(adminMember.UpdatedAt, &thisUserPublic.UpdatedAtDetails)
		// ghafpkg.TimeGenerateStructFullDetails(methods.ConvertDeletedAtToTime(adminMember.DeletedAt),
		// 	&thisUserPublic.DeletedAtDetails)

		thisPublicResponse = append(thisPublicResponse, thisUserPublic)
	}

	return thisPublicResponse, thisFetch.RowsAffected, nil
}
