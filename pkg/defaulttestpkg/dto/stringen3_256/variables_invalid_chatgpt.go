package defaulttestpkg

var VariablesInvalidChatGPT = []DTOStringEn3To256Struct{
	{
		Id:       1,
		Content:  "JohnDoe",
		Category: "invalid",
		Scenario: "نام به درستی ارسال شده است",
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
		Content:  "Jo",
		Category: "invalid",
		Scenario: "نام کوتاه‌تر از حد مجاز است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       3,
		Content:  "JohnDoe123",
		Category: "invalid",
		Scenario: "نام شامل عدد انگلیسی است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       4,
		Content:  "۱۲۳جاندو",
		Category: "invalid",
		Scenario: "نام با عدد فارسی شروع شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       5,
		Content:  "John!Doe",
		Category: "invalid",
		Scenario: "نام شامل کاراکتر ویژه است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       6,
		Content:  " نام کاربر ",
		Category: "invalid",
		Scenario: "نام شامل فاصله است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       7,
		Content:  "NamE",
		Category: "invalid",
		Scenario: "نام ترکیب حروف بزرگ و کوچک دارد",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       8,
		Content:  "Nam3",
		Category: "invalid",
		Scenario: "نام شامل ترکیب عدد و حروف است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       9,
		Content:  "",
		Category: "invalid",
		Scenario: "نام ارسال نشده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       10,
		Content:  "a very long name that exceeds the 256 character limit...",
		Category: "invalid",
		Scenario: "نام بیشتر از ۲۵۶ کاراکتر است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       11,
		Content:  "JohnDoe123", // ترکیب حروف و اعداد انگلیسی
		Category: "invalid",
		Scenario: "رشته شامل حروف و اعداد انگلیسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود (فقط حروف انگلیسی مجاز است)",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       12,
		Content:  "John_Doe", // استفاده از کاراکتر زیرخط (_)
		Category: "invalid",
		Scenario: "رشته شامل کاراکتر زیرخط ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود (کاراکترهای خاص مجاز نیستند)",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       13,
		Content:  "John-Doe", // استفاده از خط تیره (-)
		Category: "invalid",
		Scenario: "رشته شامل خط تیره ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود (کاراکترهای خاص مجاز نیستند)",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       14,
		Content:  "John Doe", // استفاده از فاصله (Space)
		Category: "invalid",
		Scenario: "رشته شامل فاصله ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود (فاصله مجاز نیست)",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       15,
		Content:  "John\nDoe", // استفاده از کاراکتر newline (\n)
		Category: "invalid",
		Scenario: "رشته شامل کاراکتر newline ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود (کاراکترهای کنترلی مجاز نیستند)",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       16,
		Content:  "John\tDoe", // استفاده از کاراکتر tab (\t)
		Category: "invalid",
		Scenario: "رشته شامل کاراکتر tab ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود (کاراکترهای کنترلی مجاز نیستند)",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       17,
		Content:  "John😊Doe", // استفاده از ایموجی
		Category: "invalid",
		Scenario: "رشته شامل ایموجی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود (کاراکترهای غیرمجاز مانند ایموجی)",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       18,
		Content:  "J", // یک کاراکتر انگلیسی
		Category: "invalid",
		Scenario: "رشته با طول 1 کاراکتر انگلیسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود (حداقل طول 3 کاراکتر)",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       19,
		Content:  "JöhnDöe", // استفاده از کاراکترهای غیر انگلیسی (مثل ö)
		Category: "invalid",
		Scenario: "رشته شامل کاراکترهای غیر انگلیسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود (فقط حروف انگلیسی مجاز است)",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
}
