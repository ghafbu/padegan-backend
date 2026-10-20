package messagepkg

import (
	"fmt"
)

func Create(lang *string, KeywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_create", KeywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s با موفقیت ثبت شد.", KeywordSection.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("%s was created successfully.", KeywordSection.English)
	}

	return clone
}

func Update(lang *string, KeywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_update", KeywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s با موفقیت بروزرسانی شد.", KeywordSection.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("%s updated successfully.", KeywordSection.English)
	}

	return clone
}

func Updates(lang *string, KeywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_updates", KeywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s با موفقیت بروزرسانی شدند.", KeywordSection.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("%s updated successfully.", KeywordSection.English)
	}

	return clone
}

func UpdateSpecial(lang *string, KeywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_update", KeywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s با موفقیت بروزرسانی شد.", KeywordSection.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("%s updated successfully.", KeywordSection.English)
	}

	return clone
}

func Get(lang *string, KeywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_get", KeywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s مورد نظر یافت شد.", KeywordSection.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("%s found successfully.", KeywordSection.English)
	}

	return clone
}

func BatchGet(lang *string, KeywordSection Language, paramCount int64) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_batch_get", KeywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("نتیجه واکشی: %d مورد.", paramCount)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("Fetch result: %d items.", paramCount)
	}

	return clone
}

func BatchUpdate(lang *string, KeywordSection Language, paramCount int64) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_batch_update", KeywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("نتیجه بروز رسانی: %d مورد.", paramCount)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("Update result: %d items.", paramCount)
	}

	return clone
}

func Search(lang *string, KeywordSection Language, paramCount int64) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_search", KeywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("نتیجه جستجو: %d مورد.", paramCount)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("Search result: %d items.", paramCount)
	}

	return clone
}

func Delete(lang *string, KeywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_delete", KeywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s با موفقیت حذف شد.", KeywordSection.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("%s deleted successfully.", KeywordSection.English)
	}

	return clone
}

func DeleteMany(lang *string, KeywordSection Language) Response {
	var clone = Response{}
	clone.Id = fmt.Sprintf("event-delete-many-%s", KeywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s با موفقیت حذف شدند.", KeywordSection.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("%s deleted successfully.", KeywordSection.English)
	}

	return clone
}

func NotFound(paramLang *string) Response {
	var clone = Response{}
	clone.Id = UtlBuildIdText("event_not_found")

	if *paramLang == "farsi" || *paramLang == "" {
		clone.Text = fmt.Sprintf("404 - یافت نشده.")
	} else if *paramLang == "english" {
		clone.Text = fmt.Sprintf("404 - No Found.")
	}

	return clone
}

func NotFoundSection(paramLang *string, KeywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_not_found_section", KeywordSection.English)

	if *paramLang == "farsi" || *paramLang == "" {
		clone.Text = fmt.Sprintf("%s یافت نشده.", KeywordSection.Farsi)
	} else if *paramLang == "english" {
		clone.Text = fmt.Sprintf("%s not found.", KeywordSection.Farsi)
	}

	return clone
}

func NotFoundSections(paramLang *string, KeywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_not_found_sections", KeywordSection.English)

	if *paramLang == "farsi" || *paramLang == "" {
		clone.Text = fmt.Sprintf("%s یافت نشدند.", KeywordSection.Farsi)
	} else if *paramLang == "english" {
		clone.Text = fmt.Sprintf("%s not found.", KeywordSection.Farsi)
	}

	return clone
}

func NotFoundField(paramLang *string, KeywordField Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_not_found_field", KeywordField.English)

	if *paramLang == "farsi" || *paramLang == "" {
		clone.Text = fmt.Sprintf("%s یافت نشده.", KeywordField.Farsi)
	} else if *paramLang == "english" {
		clone.Text = fmt.Sprintf("%s not found.", KeywordField.Farsi)
	}

	return clone
}

func NotFoundFile(paramLang *string) Response {
	var clone = Response{}
	clone.Id = UtlBuildIdText("event_not_found_file")

	if *paramLang == "farsi" || *paramLang == "" {
		clone.Text = fmt.Sprintf("فایل‌ یافت نشده.")
	} else if *paramLang == "english" {
		clone.Text = fmt.Sprintf("File not found.")
	}

	return clone
}

func NotFoundFiles(paramLang *string) Response {
	var clone = Response{}
	clone.Id = UtlBuildIdText("event_not_found_files")

	if *paramLang == "farsi" || *paramLang == "" {
		clone.Text = fmt.Sprintf("فایل‌ها یافت نشدند.")
	} else if *paramLang == "english" {
		clone.Text = fmt.Sprintf("Files not found.")
	}

	return clone
}

func NotUploadFile(paramLang *string) Response {
	var clone = Response{}
	clone.Id = UtlBuildIdText("event_not_upload_file")

	if *paramLang == "farsi" || *paramLang == "" {
		clone.Text = fmt.Sprintf("فایل ای آپلود نشده.")
	} else if *paramLang == "english" {
		clone.Text = fmt.Sprintf("file not uploaded")
	}

	return clone
}

func UploadedFileSuccess(paramLang *string) Response {
	var clone = Response{}
	clone.Id = UtlBuildIdText("event_uploaded_file_successful")

	if *paramLang == "farsi" || *paramLang == "" {
		clone.Text = fmt.Sprintf("آپلود با موفقیت انجام شد.")
	} else if *paramLang == "english" {
		clone.Text = fmt.Sprintf("The upload was successful.")
	}

	return clone
}

func NotFoundPage(paramLang *string) Response {
	var clone = Response{}
	clone.Id = UtlBuildIdText("event_not_found_page")

	if *paramLang == "farsi" || *paramLang == "" {
		clone.Text = fmt.Sprintf("صفحه یافت نشده.")
	} else if *paramLang == "english" {
		clone.Text = fmt.Sprintf("Page not found.")
	}

	return clone
}

func Status(lang *string, KeywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_status", KeywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("وضعیت %s با موفقیت تغییر داده شد.", KeywordSection.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("Status of %s successfully changed.", KeywordSection.English)
	}

	return clone
}

func Logout(lang *string, KeywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_logout", KeywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("خروج %s با موفقیت انجام شد.", KeywordSection.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("The %s logout was successful", KeywordSection.English)
	}

	return clone
}

func Send(lang *string, keywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_send", keywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s برای شما ارسال شد.", keywordSection.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("%s sent to you", keywordSection.English)
	}

	return clone
}

func Block(lang *string) Response {
	var clone = Response{}
	clone.Id = UtlBuildIdText("event_block")

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("دسترسی شما محدود شده است.")
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("Your Access is Restricted.")
	}

	return clone
}

func Verify(lang *string, KeywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_verify", KeywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s با موفقیت تایید شد.", KeywordSection.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("%s successfully verified", KeywordSection.English)
	}

	return clone
}

func Pending(lang *string, KeywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_pending", KeywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s  شما در انتظار تایید است.", KeywordSection.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("your %s is pending confirmation", KeywordSection.English)
	}

	return clone
}

func Suspended(lang *string, KeywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("event_suspended", KeywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("%s شما در حالت تعلیق است", KeywordSection.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("your %s is suspended", KeywordSection.English)
	}

	return clone
}

func WrongFTP(lang *string, KeywordSection Language) Response {
	var clone = Response{}
	// clone.Id = UtlBuildId("event_wrong_ftp", KeywordSection.English)
	clone.Id = UtlBuildId(fmt.Sprintf("event_wrong_ftp_%s", KeywordSection.English), KeywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("متاسفانه در %s مشکلی رخ داده است.", KeywordSection.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("Unfortunately, there is a problem with the %s", KeywordSection.English)
	}

	return clone
}

func MicroserviceFailed(lang *string, KeywordSection Language) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("microservice_failed_", KeywordSection.English)

	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("در ارتباط با سرویس %s مشکلی رخ داده است. دوباره تلاش کنید", KeywordSection.Farsi)
	} else if *lang == "english" {
		clone.Text = fmt.Sprintf("There is a problem with the %s service. Try again", KeywordSection.English)
	}

	return clone
}

func MicroserviceHealthTrue(lang *string, paramServiceName string) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("microservice_health_", paramServiceName)

	clone.Text = fmt.Sprintf("Service[%s] is Healthy.", paramServiceName)

	return clone
}

func JustPrintSuccessFul(paramLang *string) Response {
	var clone = Response{}
	clone.Id = UtlBuildIdText("event_not_found_page")

	if *paramLang == "farsi" || *paramLang == "" {
		clone.Text = fmt.Sprintf("موفقیت آمیز بود.")
	} else if *paramLang == "english" {
		clone.Text = fmt.Sprintf("successFul.")
	}

	return clone
}

func DefaultNumberMinError(lang *string, keyword Language) Response {
	var clone = Response{}

	clone.Id = UtlBuildId("default_number_min_required", keyword.English)
	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("حداقل یک %s باید پیش‌فرض باشد.", keyword.Farsi)
	} else {
		clone.Text = fmt.Sprintf("At least one %s must be default.", keyword.English)
	}

	return clone
}

func DefaultNumberMaxError(lang *string, keyword Language) Response {
	var clone = Response{}

	clone.Id = UtlBuildId("default_number_max_allowed", keyword.English)
	if *lang == "farsi" || *lang == "" {
		clone.Text = fmt.Sprintf("فقط یک %s می‌تواند پیش‌فرض باشد.", keyword.Farsi)
	} else {
		clone.Text = fmt.Sprintf("Only one %s can be default.", keyword.English)
	}

	return clone
}

func Success(lang *string) Response {
	var clone = Response{}
	clone.Id = UtlBuildId("operation_success", "operation_success")

	if *lang == "farsi" || *lang == "" {
		clone.Text = "عملیات با موفقیت انجام شد."
	} else if *lang == "english" {
		clone.Text = "Operation completed successfully."
	}

	return clone
}
