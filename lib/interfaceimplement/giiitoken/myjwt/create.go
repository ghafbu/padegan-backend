package myjwt

import (
	"time"

	_interface "github.com/ghafbu/padegan-backend/lib/interface"
	"github.com/golang-jwt/jwt/v5"

	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func create[Type interface{}](
	paramRequest interface{},
	paramLang string,
	paramSecret string,
	paramValue Type,
	paramExpireTime time.Duration,
	paramOwnerName string,
) (
	responsepkg.EntityV1UseCase,
	string,
	bool,
) {
	var thisMessage []messagepkg.Response

	claims := _interface.CustomClaims[Type]{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(paramExpireTime)),
		},
		Data: paramValue,
		// any more property for response to user (name, family, role, etc...)
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(paramSecret))
	if err != nil {
		thisMessage = append(thisMessage, messagepkg.StatusCode500(
			&paramLang, "-logics_token_create"),
		)
		return responsepkg.Code500(
			thisMessage,
			responsepkg.Code500CustomResponseData{
				Request:          paramRequest,
				RequestOwnerName: paramOwnerName,
				Err:              err.Error(),
			},
		), "", false
	}

	return responsepkg.EntityV1UseCase{}, tokenString, true

	// thisToken, thisErr := tokenpkg.TokenCreate(
	// 	paramSecret,
	// 	paramValue,
	// 	paramExpireTime,
	// )
	// if thisErr != nil {
	// 	thisMessage = append(thisMessage, messagepkg.StatusCode500(
	// 		&paramLang, "-logics_token_create"),
	// 	)
	// 	return responsepkg.Code500(
	// 		thisMessage,
	// 		responsepkg.Code500CustomResponseData{
	// 			Request:          paramRequest,
	// 			RequestOwnerName: paramOwnerName,
	// 			Err:              thisErr.Error(),
	// 		},
	// 	), "", false
	// }
	//
	// return responsepkg.EntityV1UseCase{}, thisToken, true
}
