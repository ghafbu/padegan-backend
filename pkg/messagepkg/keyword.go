package messagepkg

type KeywordStruct struct {
	Field   map[string]Language
	Section map[string]Language
}

// Keyword : terms for name => snake_case
// Keyword : example 1 > fieldName = Family / convert to family
// Keyword : example 2 > fieldName = Family lastname / convert to family_lastname
var Keyword = KeywordStruct{
	Field: map[string]Language{
		"id": {
			Farsi:   "شناسه",
			English: "id",
		},
		"title": {
			Farsi:   "عنوان",
			English: "title",
		},
		"slug": {
			Farsi:   "اسلاگ",
			English: "slug",
		},
		"created_at": {
			Farsi:   "تاریخ ثبت",
			English: "created_at",
		},
		"created_at_start": {
			Farsi:   "تاریخ ثبت (از)",
			English: "created_at_start",
		},
		"created_at_end": {
			Farsi:   "تاریخ ثبت (تا)",
			English: "created_at_end",
		},
		"updated_at": {
			Farsi:   "تاریخ بروزرسانی",
			English: "updated_at",
		},
		"updated_at_start": {
			Farsi:   "تاریخ بروزرسانی (از)",
			English: "updated_at_start",
		},
		"updated_at_end": {
			Farsi:   "تاریخ بروزرسانی (تا)",
			English: "updated_at_end",
		},
		"deleted_at": {
			Farsi:   "تاریخ حذف",
			English: "deleted_at",
		},
		"deleted_at_start": {
			Farsi:   "تاریخ حذف (از)",
			English: "deleted_at_start",
		},
		"deleted_at_end": {
			Farsi:   "تاریخ حذف (تا)",
			English: "deleted_at_end",
		},
		"page": {
			Farsi:   "صفحه",
			English: "page",
		},
		"item_per_page": {
			Farsi:   "تعداد نمایش",
			English: "item per page",
		},
		"order_by": {
			Farsi:   "ستون مرتب سازی",
			English: "order by",
		},
		"order_method": {
			Farsi:   "نوع مرتب سازی",
			English: "order method",
		},
		"description": {
			Farsi:   "توضیحات",
			English: "description",
		},
		"keyword": {
			Farsi:   "کلمه کلیدی",
			English: "keyword",
		},
		"keywords": {
			Farsi:   "کلمات کلیدی",
			English: "keywords",
		},
		"creator": {
			Farsi:   "ایجاد کننده",
			English: "creator",
		},
		"account": {
			Farsi:   "اکانت",
			English: "account",
		},
		"users": {
			Farsi:   "کاربران",
			English: "users",
		},
		"admin": {
			Farsi:   "مدیر",
			English: "admin",
		},
		"admins": {
			Farsi:   "مدیران",
			English: "admins",
		},
		"user": {
			Farsi:   "کاربر",
			English: "user",
		},
		"name": {
			Farsi:   "نام",
			English: "name",
		},
		"family": {
			Farsi:   "نام خانوادگی",
			English: "family",
		},
		"mobile": {
			Farsi:   "موبایل",
			English: "mobile",
		},
		"mobile_e164": {
			Farsi:   "موبایل",
			English: "mobile",
		},
		"phone": {
			Farsi:   "تلفن ثابت",
			English: "phone",
		},
		"gender": {
			Farsi:   "جنسیت",
			English: "gender",
		},
		"birthdate": {
			Farsi:   "تاریخ تولد",
			English: "birth date",
		},
		"national_code": {
			Farsi:   "کد ملی",
			English: "national code",
		},
		"job": {
			Farsi:   "شغل",
			English: "job",
		},
		"email": {
			Farsi:   "ایمیل",
			English: "email",
		},
		"username": {
			Farsi:   "نام کاربری",
			English: "username",
		},
		"password": {
			Farsi:   "رمزعبور",
			English: "password",
		},
		"status": {
			Farsi:   "وضعیت",
			English: "status",
		},
		"code": {
			Farsi:   "کد",
			English: "code",
		},
		"you": {
			Farsi:   "شما",
			English: "you",
		},
		"ids": {
			Farsi:   "شناسه ها",
			English: "id",
		},
		"mobiles": {
			Farsi:   "موبایل ها",
			English: "mobiles",
		},
		"phones": {
			Farsi:   "تلفن ثابت ها",
			English: "mobiles",
		},
		"national_codes": {
			Farsi:   "کد ملی ها",
			English: "national codes",
		},
		"jobs": {
			Farsi:   "شغل ها",
			English: "jobs",
		},
		"emails": {
			Farsi:   "ایمیل ها",
			English: "emails",
		},
		"usernames": {
			Farsi:   "نام کاربری ها",
			English: "usernames",
		},
		"statuses": {
			Farsi:   "وضعیت‌ ها",
			English: "statuses",
		},
		"target": {
			Farsi:   "هدف",
			English: "target",
		},
		"un_authorize": {
			Farsi:   "غیر مجاز",
			English: "un authorize",
		},
		"no_access": {
			Farsi:   "عدم دسترسی",
			English: "no access",
		},
		"organization": {
			Farsi:   "سازمان",
			English: "organization",
		},
		"organizations": {
			Farsi:   "سازمان ها",
			English: "organizations",
		},
		"sequence_number": {
			Farsi:   "شماره ترتیب",
			English: "sequence number",
		},
		"sort": {
			Farsi:   "مرتب سازی",
			English: "sort",
		},
		"mobile_country_code": {
			Farsi:   "کد منطقه موبایل",
			English: "mobile_country_code",
		},
		"module": {
			Farsi:   "ماٰژول",
			English: "module",
		},
		"modules": {
			Farsi:   "ماٰژول ها",
			English: "modules",
		},
		"sms": {
			Farsi:   "پیام رسان",
			English: "sms",
		},
		"sms_s": {
			Farsi:   "پیام رسان ها",
			English: "sms",
		},
		"category_org_id": {
			Farsi:   "دسته بندی سازمان ها",
			English: "category organization",
		},
		"tin": {
			Farsi:   "شماره اقتصادی",
			English: "Tax code",
		},
		"image": {
			Farsi:   "تصویر",
			English: "image",
		},
		"icon": {
			Farsi:   "آیکون",
			English: "icon",
		},
		"fax": {
			Farsi:   "شماره فکس",
			English: "fax",
		},
		"language": {
			Farsi:   "زبان",
			English: "language",
		},
		"number": {
			Farsi:   "شماره",
			English: "number",
		},
	},

	// Section
	Section: map[string]Language{
		"user": {
			Farsi:   "کاربر",
			English: "user",
		},
		"users": {
			Farsi:   "کاربران",
			English: "users",
		},
		"user_authentication": {
			Farsi:   "احرازهویت کاربر",
			English: "user authentication",
		},
		"users_authentication": {
			Farsi:   "احراز هویت کاربران",
			English: "users authentication",
		},
		"sms_config": {
			Farsi:   "تنظیمات پیام رسان",
			English: "config sms",
		},
		"sms": {
			Farsi:   "پیام رسان",
			English: "sms",
		},
		"organization": {
			Farsi:   "سازمان",
			English: "organization",
		},
		"organizations": {
			Farsi:   "سازمان‌ها",
			English: "organizations",
		},
		"organization_base_info": {
			Farsi:   "اطلاعات سازمان",
			English: "organization base information",
		},
		"organization_base_infos": {
			Farsi:   "اطلاعات سازمان",
			English: "organization base information",
		},
		"organization_blog_category": {
			Farsi:   "دسته بندی وبلاگ",
			English: "organization blog categories",
		},
		"organization_blog_categories": {
			Farsi:   "دسته بندی وبلاگ",
			English: "organization blog category",
		},
	},
}

func UpdateKeywords(newKeywords *KeywordStruct) {
	for key, value := range newKeywords.Field {
		Keyword.Field[key] = value
	}

	for key, value := range newKeywords.Section {
		Keyword.Section[key] = value
	}
}
