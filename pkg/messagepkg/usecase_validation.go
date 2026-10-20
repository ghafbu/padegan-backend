package messagepkg

import (
	"fmt"
)

func Required(lang *string, KeywordField Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("required", KeywordField.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s اجباری است.", KeywordField.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("%s is required.", KeywordField.English)
	}

	return clone
}

func Exist(lang *string, KeywordField Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("exist", KeywordField.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s وارد شده قبلا در سیستم موجود می باشد.", KeywordField.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("The %s is already available in the system. Try another.", KeywordField.English)
	}

	return clone
}

func NotExist(lang *string, KeywordField Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("not_exist", KeywordField.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s وارد شده قبلا در سیستم موجود نمی باشد.", KeywordField.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("The %s is not already available in the system. Try another.", KeywordField.English)
	}

	return clone
}

func Exists(lang *string, KeywordField Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("exists", KeywordField.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s وارد شده قبلا در سیستم موجود می باشند.", KeywordField.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("%s are already available in the system. Try another.", KeywordField.English)
	}

	return clone
}

func NotExists(lang *string, KeywordField Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("not_exists", KeywordField.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s وارد شده قبلا در سیستم موجود نمی باشند.", KeywordField.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("%s are not already available in the system. Try another.", KeywordField.English)
	}

	return clone
}

func Match(lang *string, KeywordField Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("match", KeywordField.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s به درستی وارد نشده است.", KeywordField.Farsi)

	} else if *lang == "english" {
		clone.Text = fmt.Sprintf(" %s was not entered correctly.", KeywordField.English)
	}

	return clone
}

func Lowercase(lang *string, KeywordField Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("lowercase", KeywordField.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s باید با حروف کوچک باشند.", KeywordField.Farsi)

	} else if *lang == "english" {
		clone.Text = fmt.Sprintf(" %s must be in lowercase.", KeywordField.English)
	}

	return clone
}

func Uppercase(lang *string, KeywordField Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("uppercase", KeywordField.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s باید با حروف بزرگ باشند.", KeywordField.Farsi)

	} else if *lang == "english" {
		clone.Text = fmt.Sprintf(" %s must be in uppercase.", KeywordField.English)
	}

	return clone
}

func NotEqual(lang *string, KeywordField1 Language, KeywordField2 Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId(
		"not_equal",
		fmt.Sprintf("%s %s", KeywordField1.English, KeywordField2.English),
	)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s و %s با یکدیگری برابر نیستند.", KeywordField1.Farsi, KeywordField2.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf(" %s and field %s are not equal to each other.", KeywordField1.English, KeywordField2.English)
	}

	return clone
}

func Length(lang *string, KeywordField Language, length string) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("length", KeywordField.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s باید %s کاراکتر باشد.", KeywordField.Farsi, length)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf(" %s must be %s characters.", KeywordField.English, length)
	}

	return clone
}

func LengthBetween(lang *string, KeywordField Language, minimum string, maximum string) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("length_between", KeywordField.English)

	// old
	//if *lang == "farsi" || *lang == "" {
	//	clone.Text = fmt.Sprintf("%s باید حداقل %s و حداکثر %s کارکتر باشد.", KeywordField.Farsi, minimum, maximum)
	//} else if *lang == "english" {
	//	clone.Text = fmt.Sprintf("The %s must be at least %s and at most %s characters.", KeywordField.English, minimum, maximum)
	//}

	// new
	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s باید حداقل %s کارکتر باشد.", KeywordField.Farsi, minimum)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("The %s must be at least %s characters.", KeywordField.English, minimum)
	}

	return clone
}

func LengthMax(lang *string, KeywordField Language, maximum string) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("length_max", KeywordField.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s باید حداکثر %s کاراکتر باشد.", KeywordField.Farsi, maximum)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("The %s KeywordField must be at most %s characters.", KeywordField.English, maximum)
	}

	return clone
}

func LengthMin(lang *string, KeywordField Language, minimum string) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("length_min", KeywordField.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s باید حداقل %s کاراکتر باشد.", KeywordField.Farsi, minimum)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf(" %s must be at least %s characters.", KeywordField.English, minimum)
	}

	return clone
}

func ArrayLength(lang *string, KeywordField Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("array_length", KeywordField.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("تعداد %s به درستی ارسال نشده است.", KeywordField.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("The number of %s was not submitted correctly.", KeywordField.English)
	}

	return clone
}

func NotStandardRequest(lang *string, KeywordField Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("not_standard_request", KeywordField.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("درخواست شما به شکل استاندارد ارسال نشده است.")
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("Your request has not been submitted in the standard format.")
	}

	return clone
}

func Wrong(lang *string, keywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event-wrong", keywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s وارد شده نادرست است.", keywordSection.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("%s Entered is incorrect", keywordSection.English)
	}

	return clone
}

func NotLogin(lang *string, keywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event-not-login", keywordSection.English)

	if *lang == "farsi" || *lang == "" {
		// clone.Text = fmt.Sprintf("لطفا اقدام به لاگین کنید.", keywordSection.Farsi)
		clone.Text = "لطفا اقدام به لاگین کنید."
	} else if *lang == "english" {
		// clone.Text = fmt.Sprintf("please log in", keywordSection.English)
		clone.Text = "please log in"
	}

	return clone
}

func UnAuthorized(lang *string, keywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event-un-authorized", keywordSection.English)

	if *lang == "farsi" || *lang == "" {
		// clone.Text = fmt.Sprintf("لطفا اقدام به لاگین کنید.", keywordSection.Farsi)
		clone.Text = "احراز هویت ناموفق بوده یا هنوز ارائه نشده است"
	} else if *lang == "english" {
		// clone.Text = fmt.Sprintf("please log in", keywordSection.English)
		clone.Text = "authentication is failed or not yet been provided"
	}

	return clone
}

func NoAccess(lang *string, keywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event-", keywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = "شما دسترسی به این بخش را ندارید."
	} else if *lang == "english" {
		clone.Text = "You do not have access to this section"
	}

	return clone
}

func IsValidAndSecurePassword(lang *string, KeywordField Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("match", KeywordField.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s ناامن، سعی کنید نویسه‌های خاص بیشتری اضافه کنید، از حروف بزرگ، استفاده از اعداد یا رمز عبور طولانی‌تر استفاده کنید.", KeywordField.Farsi)

	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("insecure %s, try including more special characters, using uppercase letters, using numbers or using a longer password", KeywordField.English)
	}

	return clone
}
