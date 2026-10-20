package errors

import (
	"github.com/ghafbu/padegan-backend/lib/consts"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func ModelAction(paramError error, paramRequest interface{}, paramLang string, paramServiceName string, paramModelActionName string) responsepkg.EntityV1UseCase {
	var thisMessage []messagepkg.Response

	//get model error
	var thisErrorHappened = paramError.Error()

	var thisLogPreset = "ServiceName:" + paramServiceName + " Layer:Model ModelActionName:" + paramModelActionName

	//db exist thisError
	if thisErrorHappened == "db_exist_mobile" ||
		thisErrorHappened == "db_exist_national_code" ||
		thisErrorHappened == "db_exist_email" ||
		thisErrorHappened == "db_exist_username" {
		thisMessage = append(thisMessage, messagepkg.StatusCode500(&paramLang, thisLogPreset+" db exist"))
		return responsepkg.Code500(
			thisMessage,
			thisErrorHappened,
		)
	}

	//exist email
	if thisErrorHappened == "exist_email" {
		thisMessage = append(thisMessage, messagepkg.Exist(&paramLang, messagepkg.Keyword.Field["email"]))
		return responsepkg.Code409(
			thisMessage,
			paramRequest,
		)
	}

	//exist
	if thisErrorHappened == "exist_title" {
		thisMessage = append(thisMessage, messagepkg.Exist(&paramLang, messagepkg.Keyword.Field["title"]))
		return responsepkg.Code409(
			thisMessage,
			paramRequest,
		)
	}

	//exist
	//if thisErrorHappened == "exist_slug" {
	//	thisMessage = append(thisMessage, messagepkg.Exist(&paramLang, messagepkg.Keyword.Field["slug"]))
	//	return responsepkg.Code409(
	//		thisMessage,
	//		paramRequest,
	//	)
	//}

	//exist username
	if thisErrorHappened == "exist_username" {
		thisMessage = append(thisMessage, messagepkg.Exist(&paramLang, messagepkg.Keyword.Field["username"]))
		return responsepkg.Code409(
			thisMessage,
			paramRequest,
		)
	}

	//exist nationalCode
	if thisErrorHappened == "exist_national_code" {
		thisMessage = append(thisMessage, messagepkg.Exist(&paramLang, messagepkg.Keyword.Field["national_code"]))
		return responsepkg.Code409(
			thisMessage,
			paramRequest,
		)
	}

	// not found record
	if thisErrorHappened == consts.ErrorKeyModelRecordNotFound {
		thisMessage = append(thisMessage, messagepkg.NotFound(&paramLang))
		return responsepkg.Code404(
			thisMessage,
			paramRequest,
		)
	}

	//birthdate parse error
	if thisErrorHappened == "date_parse_error" {
		thisMessage = append(thisMessage, messagepkg.StatusCode500(&paramLang, thisLogPreset+" date parse error"))
		return responsepkg.Code500(
			thisMessage,
			thisErrorHappened,
		)
	}

	//password hash generate error
	if thisErrorHappened == "password_hash_generate_error" {
		thisMessage = append(thisMessage, messagepkg.StatusCode500(&paramLang, thisLogPreset+" password hash generate error"))
		return responsepkg.Code500(
			thisMessage,
			thisErrorHappened,
		)
	}

	// Id not standard
	//if thisErrorHappened == "id-bad-request" {
	//	//return forLogger
	//	thisMessage = append(thisMessage, messagepkg.Match(&paramLang, messagepkg.Keyword.Field["id"]))
	//	return responsepkg.Code400(
	//		thisMessage,
	//		paramRequest,
	//	)
	//}

	//Anonymous Error
	thisMessage = append(thisMessage, messagepkg.StatusCode500(&paramLang, thisLogPreset+" anonymous Error"))
	return responsepkg.Code500(
		thisMessage,
		thisErrorHappened,
	)
}
