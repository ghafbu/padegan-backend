package defaulttestpkg

import "strings"

var VariablesInvalidDeepSeek = []DTOStringEn3To256Struct{
	{
		Id:       1,
		Content:  "abc",
		Category: "invalid",
		Scenario: "رشته به درستی ارسال شده است",
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
		Content:  "ab",
		Category: "invalid",
		Scenario: "رشته به درستی ارسال نشده است",
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
		Content:  "۱۲۳", // فارسی
		Category: "invalid",
		Scenario: "رشته با اعداد فارسی ارسال شده است",
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
		Content:  "123", // انگلیسی
		Category: "invalid",
		Scenario: "رشته با اعداد انگلیسی ارسال شده است",
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
		Content:  "نام طولانی که بیش از 256 کاراکتر است " + strings.Repeat("a", 257),
		Category: "invalid",
		Scenario: "رشته با طول بیش از حد مجاز ارسال شده است",
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
		Content:  "نام",
		Category: "invalid",
		Scenario: "رشته با طول کمتر از حد مجاز ارسال شده است",
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
		Content:  "نام@خاص",
		Category: "invalid",
		Scenario: "رشته شامل کاراکترهای خاص ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       8,
		Content:  "نام معمولی",
		Category: "invalid",
		Scenario: "رشته به درستی ارسال شده است",
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
		Content:  "", // خالی
		Category: "invalid",
		Scenario: "رشته خالی ارسال شده است",
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
		Content:  "   ", // فاصله خالی
		Category: "invalid",
		Scenario: "رشته شامل فقط فاصله خالی ارسال شده است",
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
		Content:  "नाम", // هندی
		Category: "invalid",
		Scenario: "رشته به زبان هندی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       12,
		Content:  "名字", // چینی
		Category: "invalid",
		Scenario: "رشته به زبان چینی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       13,
		Content:  "名前", // ژاپنی
		Category: "invalid",
		Scenario: "رشته به زبان ژاپنی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       14,
		Content:  "이름", // کره‌ای
		Category: "invalid",
		Scenario: "رشته به زبان کره‌ای ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       15,
		Content:  "اسم", // عربی
		Category: "invalid",
		Scenario: "رشته به زبان عربی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       16,
		Content:  "نام طولانی که بیش از 256 کاراکتر است " + strings.Repeat("a", 256),
		Category: "invalid",
		Scenario: "رشته با طول دقیقا 256 کاراکتر ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       17,
		Content:  "نام طولانی که بیش از 256 کاراکتر است " + strings.Repeat("a", 255),
		Category: "invalid",
		Scenario: "رشته با طول 255 کاراکتر ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       18,
		Content:  "نام طولانی که بیش از 256 کاراکتر است " + strings.Repeat("a", 3),
		Category: "invalid",
		Scenario: "رشته با طول دقیقا 3 کاراکتر ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       19,
		Content:  "نام طولانی که بیش از 256 کاراکتر است " + strings.Repeat("a", 2),
		Category: "invalid",
		Scenario: "رشته با طول کمتر از 3 کاراکتر ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       20,
		Content:  "نام طولانی که بیش از 256 کاراکتر است " + strings.Repeat("a", 1),
		Category: "invalid",
		Scenario: "رشته با طول 1 کاراکتر ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: خطا - در خروجی کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
}
