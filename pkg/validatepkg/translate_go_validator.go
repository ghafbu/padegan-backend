package validatepkg

import (
	"strings"

	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

func TranslateGoValidator(lang string, ErrorList validator.ValidationErrors) ([]messagepkg.Response, []CustomValidate) {
	var _mess []messagepkg.Response
	var _data []CustomValidate
	var ItemCustom CustomValidate

	var TagParam = make([]string, 10)

	for _, OldItem := range ErrorList {
		ItemCustom.Name = OldItem.Field()
		// ItemCustom.Json = slugpkg.Make(OldItem.Field(), "_")
		ItemCustom.Json = strcase.ToSnake(OldItem.Field())
		ItemCustom.Tag = OldItem.Tag()
		ItemCustom.Kind = OldItem.Kind()
		TagParam = strings.Split(OldItem.Param(), "-")

		// ItemCustom.Message.Id = messagepkg.UtlBuildId(ItemCustom.Tag, ItemCustom.Name)
		ItemCustom.Message.Id = messagepkg.UtlBuildId(ItemCustom.Tag, ItemCustom.Json)

		// fmt.Println("ItemCustom.Json:", ItemCustom.Json, OldItem.Tag(), OldItem.Field(), strcase.ToSnake(OldItem.Field()))

		switch OldItem.Tag() {
		case "required":
			ItemCustom.Message.Text = messagepkg.Required(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_slugify":
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "eqfield":
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_repeat":
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		// case "omitempty":
		//	ItemCustom.Param = make([]string, 1)
		//	ItemCustom.Param[0] = TagParam[0]
		//	ItemCustom.Message.Text = messagepkg.Length(&lang, messagepkg.Keyword.Field[ItemCustom.Json], ItemCustom.Param[0]).Text
		case "gh_length":
			ItemCustom.Param = make([]string, 2)
			ItemCustom.Param[0] = TagParam[0]
			ItemCustom.Param[1] = TagParam[1]
			ItemCustom.Message.Text = messagepkg.LengthBetween(&lang, messagepkg.Keyword.Field[ItemCustom.Json], ItemCustom.Param[0], ItemCustom.Param[1]).Text
		case "gh_array_len":
			ItemCustom.Param = make([]string, 2)
			// ItemCustom.Param[0] = TagParam[0]
			// ItemCustom.Param[1] = TagParam[1]
			ItemCustom.Message.Text = messagepkg.ArrayLength(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "len":
			ItemCustom.Param = make([]string, 1)
			ItemCustom.Param[0] = TagParam[0]
			ItemCustom.Message.Text = messagepkg.Length(&lang, messagepkg.Keyword.Field[ItemCustom.Json], ItemCustom.Param[0]).Text
		case "gh_phone_by_country_code":
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_phone_format_e164":
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_sanitize":
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_printable":
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_printable_with_zwj":
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_word_limit":
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_char_abc":
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_char_abc_with_space_and_zwj":
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_email":
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_username":
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_object", "gh_object_media_pro", "gh_object_media_slim", "gt":
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "lowercase":
			ItemCustom.Message.Text = messagepkg.Lowercase(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "uppercase":
			ItemCustom.Message.Text = messagepkg.Uppercase(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "numeric":
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "min":
			ItemCustom.Param = make([]string, 1)
			ItemCustom.Param[0] = TagParam[0]
			ItemCustom.Message.Text = messagepkg.LengthMin(&lang, messagepkg.Keyword.Field[ItemCustom.Json], ItemCustom.Param[0]).Text
		case "max":
			ItemCustom.Param = make([]string, 1)
			ItemCustom.Param[0] = TagParam[0]
			ItemCustom.Message.Text = messagepkg.LengthMax(&lang, messagepkg.Keyword.Field[ItemCustom.Json], ItemCustom.Param[0]).Text
		case "email":
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "datetime":
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		// case "":
		//	ItemCustom.Message.Text = messagepkg.Required(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		// case "":
		//	ItemCustom.Message.Text = messagepkg.Required(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		// case "":
		//	ItemCustom.Message.Text = messagepkg.Required(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		// case "":
		//	ItemCustom.Message.Text = messagepkg.Required(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text

		case "oneof":
			ItemCustom.Param = strings.Split(OldItem.Param(), " ")
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_endswith":
			ItemCustom.Param = strings.Split(OldItem.Param(), " ")
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_g_ends_with":
			ItemCustom.Param = strings.Split(OldItem.Param(), " ")
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_starts_with_plus":
			ItemCustom.Param = strings.Split(OldItem.Param(), " ")
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_valid_and_secure_password":
			ItemCustom.Param = strings.Split(OldItem.Param(), " ")
			ItemCustom.Message.Text = messagepkg.IsValidAndSecurePassword(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_is_valid_country_code":
			ItemCustom.Param = strings.Split(OldItem.Param(), " ")
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text
		case "gh_seo_robots_option":
			ItemCustom.Param = strings.Split(OldItem.Param(), " ")
			ItemCustom.Message.Text = messagepkg.Match(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text

		default:
			ItemCustom.Param = strings.Split(OldItem.Param(), "-")
			ItemCustom.Message.Text = messagepkg.Required(&lang, messagepkg.Keyword.Field[ItemCustom.Json]).Text

		}

		// fmt.Println(OldItem.Tag())
		// _data = append(_data, MyCustom)
		// var a = CustomValidate{
		//	Field: fmt.Sprintf(
		//		"%v / %v / %v / %v / %v / %v / %v",
		//		OldItem.ActualTag(),
		//		OldItem.Kind(),
		//		ItemCustom.Name,
		//		OldItem.Namespace(),
		//		OldItem.Param(),
		//		OldItem.StructField(),
		//		OldItem.StructNamespace(),
		//	),
		//	Message: messagepkg.Response{},
		// }

		_mess = append(_mess, ItemCustom.Message)
		_data = append(_data, ItemCustom)
	}
	return _mess, _data
}
