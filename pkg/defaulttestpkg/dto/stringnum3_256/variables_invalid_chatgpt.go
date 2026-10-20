package defaulttestpkg

var VariablesInvalidChatGPT = []DTOStringNum3To256Struct{
	{
		Id:       1,
		Content:  "1234567890",
		Category: "invalid",
		Scenario: "اعداد به درستی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       2,
		Content:  "۱۲۳۴۵۶۷۸۹۰", // عدد فارسی
		Category: "invalid",
		Scenario: "اعداد فارسی به درستی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       3,
		Content:  "०१२३४५६७८९", // عدد هندی
		Category: "invalid",
		Scenario: "اعداد هندی به درستی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       4,
		Content:  "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890", // طول بیش از ۲۵۶
		Category: "invalid",
		Scenario: "طول عدد بیشتر از ۲۵۶ کاراکتر است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - طول بیشتر از ۲۵۶ کاراکتر غیرقابل قبول است",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       5,
		Content:  "12345", // طول کمتر از ۳
		Category: "invalid",
		Scenario: "طول عدد کمتر از ۳ کاراکتر است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - طول کمتر از ۳ کاراکتر غیرقابل قبول است",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       6,
		Content:  "1234abcd567", // شامل کاراکتر غیر عددی
		Category: "invalid",
		Scenario: "وجود کاراکتر غیرعدد در ورودی",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - ورودی باید فقط عدد باشد",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       7,
		Content:  "###123456###", // وجود کاراکتر خاص
		Category: "invalid",
		Scenario: "وجود کاراکترهای خاص در ورودی",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - ورودی نمی‌تواند شامل کاراکترهای خاص باشد",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       8,
		Content:  "١٢٣٤٥٦٧٨٩٠", // عدد عربی
		Category: "invalid",
		Scenario: "اعداد عربی به درستی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       9,
		Content:  "٠١٢٣٤٥٦٧٨٩", // عدد صفر
		Category: "invalid",
		Scenario: "اعداد صفر به درستی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       10,
		Content:  "۱۲۳۴۵۶7890", // عدد فارسی با ترکیب عدد انگلیسی
		Category: "invalid",
		Scenario: "ترکیب اعداد فارسی و انگلیسی به درستی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       11,
		Content:  "۱۲۳4567890", // عدد فارسی با ترکیب عدد انگلیسی
		Category: "invalid",
		Scenario: "ترکیب اعداد فارسی و انگلیسی در وسط ورودی",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       12,
		Content:  "123@456789", // ورودی حاوی کاراکتر خاص
		Category: "invalid",
		Scenario: "وجود کاراکتر خاص در ورودی",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - ورودی نمی‌تواند شامل کاراکترهای خاص باشد",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       13,
		Content:  "一二三四五六七八九零", // اعداد چینی
		Category: "invalid",
		Scenario: "اعداد چینی به درستی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       14,
		Content:  "12345678901", // طول درست
		Category: "invalid",
		Scenario: "عدد با طول درست و بدون خطا ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       15,
		Content:  "٢٣٤٥٦٧٨٩٠", // عدد فارسی
		Category: "invalid",
		Scenario: "عدد فارسی به درستی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       16,
		Content:  "12345678!@#$%^", // وجود کاراکترهای خاص
		Category: "invalid",
		Scenario: "وجود کاراکترهای خاص در عدد",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - ورودی نمی‌تواند شامل کاراکترهای خاص باشد",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       17,
		Content:  "٠١٢٣٤٥٦٧٨٩١٢٣٤٥٦٧٨٩١٢٣٤٥٦٧٨٩١٢٣٤٥٦٧٨٩١٢٣٤٥٦٧٨٩١٢٣٤٥٦٧٨٩١٢٣٤٥٦٧٨٩", // طول زیاد
		Category: "invalid",
		Scenario: "طول عدد بیشتر از ۲۵۶ است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - طول عدد بیشتر از ۲۵۶ غیرقابل قبول است",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       18,
		Content:  "1", // طول کمتر از حداقل
		Category: "invalid",
		Scenario: "عدد با طول کمتر از حداقل ۳ کاراکتر",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - طول کمتر از ۳ کاراکتر غیرقابل قبول است",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       19,
		Content:  "۱٢٣۴۵٦", // ترکیب اعداد فارسی و عربی
		Category: "invalid",
		Scenario: "ترکیب اعداد فارسی و عربی",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       20,
		Content:  "12345678901234567890123456789012345678901234567890", // مقدار مجاز
		Category: "invalid",
		Scenario: "مقدار مجاز (تعداد درست اعداد)",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       21,
		Content:  "۱۲۳۴۵۶۷۸۹۰", // عدد به فارسی
		Category: "invalid",
		Scenario: "عدد فارسی به درستی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       22,
		Content:  "12345678901@", // عدد با کاراکتر خاص
		Category: "invalid",
		Scenario: "عدد همراه با کاراکتر خاص ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تست ناموفق - زیرا عدد با کاراکتر خاص ارسال شده است.",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       23,
		Content:  "१२३४५६७८९०", // عدد به هندی
		Category: "invalid",
		Scenario: "عدد هندی به درستی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       24,
		Content:  "۱۲۳۴۵۶۷۸۹۰آ", // عدد با حرف فارسی
		Category: "invalid",
		Scenario: "عدد همراه با حرف فارسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تست ناموفق - زیرا عدد همراه با حرف فارسی ارسال شده است.",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       25,
		Content:  "٣٣٣٣٣٣٣٣٣٣", // عدد با اعداد عربی
		Category: "invalid",
		Scenario: "عدد عربی به درستی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       26,
		Content:  "١٢٣٤٥٦٧٨٩٠", // عدد با اعداد عربی (از شماره 1 تا 9 به صورت عربی)
		Category: "invalid",
		Scenario: "عدد عربی (از شماره 1 تا 9 به صورت عربی) به درستی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       27,
		Content:  "1234567890123456789012345678901234567890", // طول بسیار بلند (بیش از 256 کاراکتر)
		Category: "invalid",
		Scenario: "عدد طولانی تر از حد مجاز ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تست ناموفق - زیرا طول عدد بیش از حد مجاز است.",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       28,
		Content:  "٠١٢٣٤٥٦٧٨٩", // عدد با اعداد فارسی و صفرهای ابتداییش
		Category: "invalid",
		Scenario: "عدد با صفرهای ابتدایی به درستی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       29,
		Content:  "١٢٣٤٥٦٧٨٩٠abcdef", // عدد با حروف انگلیسی
		Category: "invalid",
		Scenario: "عدد همراه با حروف انگلیسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تست ناموفق - زیرا عدد همراه با حروف انگلیسی ارسال شده است.",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       30,
		Content:  "۱۲۳۴۵۶۷۸۹٠٠٠", // عدد با عدد فارسی و صفرهای اضافی
		Category: "invalid",
		Scenario: "عدد فارسی همراه با صفرهای اضافی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
}
