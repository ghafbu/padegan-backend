package logics

import (
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/phonepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func PhoneDetails(
	paramRequest interface{},
	paramPhone *string,
	paramLang string,
	paramFieldName string,
) (
	phonepkg.PhoneDetails,
	responsepkg.EntityV1UseCase,
	bool,
) {
	var thisMessage []messagepkg.Response

	// if phone is empty
	if paramPhone == nil {
		thisMessage = append(thisMessage, messagepkg.Wrong(&paramLang, messagepkg.Keyword.Field[paramFieldName]))
		return phonepkg.PhoneDetails{}, responsepkg.Code400(
			thisMessage,
			paramRequest,
		), false
	}

	// phone details
	thisPhoneDetails, err := phonepkg.GetPhoneDetails(*paramPhone)
	if err != nil {
		thisMessage = append(thisMessage, messagepkg.Wrong(&paramLang, messagepkg.Keyword.Field[paramFieldName]))
		return thisPhoneDetails, responsepkg.Code400(
			thisMessage,
			paramRequest,
		), false
	}

	return thisPhoneDetails, responsepkg.EntityV1UseCase{}, true
}
