package myjwt

import (
	_interface "github.com/ghafbu/padegan-backend/lib/interface"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
	"github.com/golang-jwt/jwt/v5"
)

func parse[Type interface{}](
	paramRequest interface{},
	paramLang string,
	paramSecret string,
	paramToken string,
	paramOwnerName string,
) (
	responsepkg.EntityV1UseCase,
	*_interface.CustomClaims[Type],
	bool,
) {
	var thisMessage []messagepkg.Response

	token, err := jwt.ParseWithClaims(
		paramToken,
		&_interface.CustomClaims[Type]{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(paramSecret), nil // secret is generateCustomTokenSecret
		})
	if err != nil {
		thisMessage = append(thisMessage, messagepkg.StatusCode500(
			&paramLang, "-logics_token_parse"),
		)
		return responsepkg.Code500(
			thisMessage,
			responsepkg.Code500CustomResponseData{
				Request:          paramRequest,
				RequestOwnerName: paramOwnerName,
				Err:              err.Error(),
			},
		), nil, false
	}

	if claims, ok := token.Claims.(*_interface.CustomClaims[Type]); ok && token.Valid {
		return responsepkg.EntityV1UseCase{}, claims, true
	}

	thisMessage = append(thisMessage, messagepkg.StatusCode500(
		&paramLang, "-logics_token_parse"),
	)
	return responsepkg.Code500(
		thisMessage,
		responsepkg.Code500CustomResponseData{
			Request:          paramRequest,
			RequestOwnerName: paramOwnerName,
			Err:              err.Error(),
		},
	), nil, false

	// thisParseToken, thisErr := tokenpkg.TokenParse[Type](
	// 	paramSecret,
	// 	paramToken,
	// )
	// if thisErr != nil {
	// 	thisMessage = append(thisMessage, messagepkg.StatusCode500(
	// 		&paramLang, "-logics_token_parse"),
	// 	)
	// 	return responsepkg.Code500(
	// 		thisMessage,
	// 		responsepkg.Code500CustomResponseData{
	// 			Request:          paramRequest,
	// 			RequestOwnerName: paramOwnerName,
	// 			Err:              thisErr.Error(),
	// 		},
	// 	), nil, false
	// }
	//
	// return responsepkg.EntityV1UseCase{}, thisParseToken, true
}
