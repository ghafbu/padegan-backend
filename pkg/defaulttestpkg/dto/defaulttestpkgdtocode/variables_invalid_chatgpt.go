package defaulttestpkgdtocode

var variablesInvalidChatGPT = []myStruct{
	{
		Id:       1,
		Content:  "12345",
		Category: "invalid",
		Scenario: "کد عددی ۵ رقمی صحیح ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       2,
		Content:  "۱۲۳۴۵", // فارسی
		Category: "invalid",
		Scenario: "کد عددی فارسی ۵ رقمی ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       3,
		Content:  "123۴۵", // ترکیب انگلیسی و فارسی
		Category: "invalid",
		Scenario: "کد شامل اعداد انگلیسی و فارسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: شکست - ترکیب اعداد انگلیسی و فارسی مجاز نیست",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       4,
		Content:  "१२३४५", // هندی
		Category: "invalid",
		Scenario: "خالی ارسال شد",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       5,
		Content:  "abcde", // حروف انگلیسی
		Category: "invalid",
		Scenario: "کد حروف انگلیسی ۵ رقمی ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: شکست - فقط اعداد مجاز هستند",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       6,
		Content:  "!@#$%", // کاراکترهای خاص
		Category: "invalid",
		Scenario: "کد با کاراکترهای خاص ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: شکست - کاراکترهای خاص مجاز نیستند",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       7,
		Content:  "1234567890", // طول بیش از حد
		Category: "invalid",
		Scenario: "کد بیش از ۵ رقم ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: شکست - طول کد باید ۵ رقم باشد",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       8,
		Content:  "۱۲۳۴", // طول کم‌تر از ۵
		Category: "invalid",
		Scenario: "کد با طول کمتر از ۵ رقم ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: شکست - طول کد باید ۵ رقم باشد",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       9,
		Content:  "۱۲۳45", // ترکیب فارسی و انگلیسی
		Category: "invalid",
		Scenario: "کد با ترکیب فارسی و انگلیسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: شکست - ترکیب فارسی و انگلیسی مجاز نیست",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       10,
		Content:  "abcdef", // حروف غیرعدد
		Category: "invalid",
		Scenario: "کد با حروف غیرعدد ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: شکست - فقط اعداد مجاز هستند",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       11,
		Content:  "١٢٣۴۵", // اعداد عربی
		Category: "invalid",
		Scenario: "کد عددی عربی ۵ رقمی ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       12,
		Content:  "123a5", // وجود حروف در کد
		Category: "invalid",
		Scenario: "کد با وجود حروف ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: شکست - کد فقط باید شامل اعداد باشد",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       13,
		Content:  "¡¿@#&", // کاراکترهای غیرمجاز
		Category: "invalid",
		Scenario: "کد با کاراکترهای غیرمجاز ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: شکست - کاراکترهای خاص مجاز نیستند",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       14,
		Content:  "زبان", // کاراکترهای غیرعدد
		Category: "invalid",
		Scenario: "کد با کاراکترهای غیرعدد فارسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: شکست - کد باید فقط شامل اعداد باشد",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       15,
		Content:  "٥٤٣٢١", // اعداد عربی به فارسی
		Category: "invalid",
		Scenario: "کد عددی عربی ۵ رقمی ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       16,
		Content:  "123,45", // کاما در کد
		Category: "invalid",
		Scenario: "کد با کاما (,) ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: شکست - کاما مجاز نیست",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       17,
		Content:  "6AB56", // ترکیب عدد و حروف
		Category: "invalid",
		Scenario: "کد با ترکیب عدد و حروف ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: شکست - فقط اعداد مجاز هستند",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       18,
		Content:  "۱۲۳۴&", // کاراکتر خاص در فارسی
		Category: "invalid",
		Scenario: "کد با کاراکتر خاص در کد فارسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: شکست - کاراکترهای خاص مجاز نیستند",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       19,
		Content:  "123٤5", // عدد عربی در کد
		Category: "invalid",
		Scenario: "کد با ترکیب اعداد عربی و فارسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: شکست - ترکیب اعداد عربی و فارسی مجاز نیست",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       20,
		Content:  "00000", // همه صفر
		Category: "invalid",
		Scenario: "کد با صفرهای مکرر ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       21,
		Content:  "۱۲۳۴۵",
		Category: "invalid",
		Scenario: "کد با اعداد فارسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: در صورت ارسال اعداد فارسی، باید در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       22,
		Content:  "123@!#",
		Category: "invalid",
		Scenario: "کد با کاراکترهای خاص ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: ارسال کاراکترهای خاص در کد، باعث شکست تست و کد 400 باید برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       23,
		Content:  "abcde",
		Category: "invalid",
		Scenario: "کد با حروف انگلیسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: استفاده از حروف انگلیسی در کد باعث شکست تست و کد 400 باید برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       24,
		Content:  "۱۲۳۴۶۸۷۹۰",
		Category: "invalid",
		Scenario: "کد طول بیشتر از ۵ کاراکتر دارد",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: طول کد باید ۵ کاراکتر باشد و اگر بیشتر باشد، تست شکست می‌خورد و کد 400 برگشت داده می‌شود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       25,
		Content:  "۱۲۳۴۵۶",
		Category: "invalid",
		Scenario: "کد با طول غیرمجاز (۶ کاراکتر) ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: طول کد باید ۵ کاراکتر باشد، در غیر اینصورت تست باید شکست بخورد و کد 400 برگشت داده شود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       26,
		Content:  "١٢٣٤٥", // Persian numbers (Unicode form)
		Category: "invalid",
		Scenario: "کد با اعداد فارسی یونی‌کد شده ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: ارسال اعداد فارسی به صورت یونی‌کد باعث شکست تست و کد 400 باید برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       27,
		Content:  "١٢٣٤٥١",
		Category: "invalid",
		Scenario: "کد با عدد فارسی و یک عدد اضافه (طول بیشتر از ۵) ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: طول کد باید ۵ باشد، ارسال کد طولانی‌تر باعث شکست تست و کد 400 برگشت داده می‌شود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       28,
		Content:  "۱۲۳۴۵#%",
		Category: "invalid",
		Scenario: "کد با ترکیب اعداد فارسی و کاراکترهای خاص ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: کد باید فقط شامل اعداد باشد، ترکیب اعداد با کاراکترهای خاص باعث شکست تست و کد 400 برگشت داده می‌شود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       29,
		Content:  "12345لطفا",
		Category: "invalid",
		Scenario: "کد با عدد و یک کلمه فارسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: کد باید فقط شامل اعداد باشد، ترکیب با کلمه فارسی باعث شکست تست و کد 400 برگشت داده می‌شود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       30,
		Content:  "١٢٣۴۵",
		Category: "invalid",
		Scenario: "کد با اعداد فارسی و عدد غیر صحیح (۴ به جای ۴) ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: استفاده از اعداد فارسی و سایر کاراکترهای غیر متداول باید باعث شکست تست و کد 400 برگشت داده شود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       31,
		Content:  "123۴۵",
		Category: "invalid",
		Scenario: "کد با عدد فارسی به شکل ترکیب با اعداد انگلیسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: ترکیب اعداد فارسی و انگلیسی باعث شکست تست و کد 400 برگشت داده می‌شود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       32,
		Content:  "١٢٣۴۵",
		Category: "invalid",
		Scenario: "کد با اعداد فارسی و یک عدد اشتباه (۴ به جای ۵) ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: فقط باید اعداد صحیح 5 رقمی ارسال شوند، در غیر این صورت تست شکست می‌خورد و کد 400 برگشت داده می‌شود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       33,
		Content:  "54321",
		Category: "invalid",
		Scenario: "کد معکوس با اعداد انگلیسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       34,
		Content:  "١٢٣۴٥",
		Category: "invalid",
		Scenario: "کد با عدد فارسی به شکل ترکیب درست با اعداد انگلیسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: ارسال اعداد فارسی همراه با اعداد انگلیسی باعث شکست تست و کد 400 برگشت داده می‌شود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       35,
		Content:  "الفبای12345",
		Category: "invalid",
		Scenario: "کد با ترکیب حروف فارسی و اعداد ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: کد باید فقط از اعداد تشکیل شده باشد، ارسال حروف فارسی باعث شکست تست و کد 400 برگشت داده می‌شود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       36,
		Content:  "١٢٣۴۵۶",
		Category: "invalid",
		Scenario: "کد با ۶ عدد فارسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: کد باید دقیقاً ۵ عدد باشد، ارسال ۶ عدد باعث شکست تست و کد 400 برگشت داده می‌شود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       37,
		Content:  "123a45",
		Category: "invalid",
		Scenario: "کد با یک حرف همراه با عدد ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: ترکیب اعداد با حروف باعث شکست تست و کد 400 برگشت داده می‌شود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       38,
		Content:  "٩٩٩٩۹",
		Category: "invalid",
		Scenario: "کد با اعداد فارسی و عربی ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: ارسال اعداد فارسی و عربی باعث شکست تست و کد 400 برگشت داده می‌شود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       39,
		Content:  "١٢۳۴۵",
		Category: "invalid",
		Scenario: "کد با اعداد فارسی و عدد غیر متعارف 3 ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: استفاده از اعداد فارسی باید با فرم صحیح باشد، در غیر این صورت کد 400 برگشت داده می‌شود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       40,
		Content:  "12345@#",
		Category: "invalid",
		Scenario: "کد با عدد و کاراکترهای خاص ارسال شده است",
		Provider: "ChatGPT",
		Expected: expected{
			Description: "انتظار ما: کد باید فقط شامل اعداد باشد، ارسال کاراکترهای خاص باعث شکست تست و کد 400 برگشت داده می‌شود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
}
