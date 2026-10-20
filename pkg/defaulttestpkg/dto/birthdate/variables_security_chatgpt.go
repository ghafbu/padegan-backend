package defaulttestpkg

var VariablesSecurityChatGPT = []DTOBirthdateStruct{
	{
		Id:       1,
		Content:  "2023/15/30",
		Category: "invalid",
		Scenario: "تاریخ تولد با فرمت اشتباه ارسال شده است (روز و ماه اشتباه)",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای ناموفق - باید خطا در خروجی مشاهده شود.",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       2,
		Content:  "15-01-2023",
		Category: "invalid",
		Scenario: "تاریخ تولد با فرمت اشتباه ارسال شده است (فرمت مختلف)",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای ناموفق - باید خطا در خروجی مشاهده شود.",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       3,
		Content:  "1970-01-01",
		Category: "valid",
		Scenario: "تاریخ تولد صحیح ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای موفق - کد 200 باید بازگردانده شود.",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       4,
		Content:  "۱۳۷۰/۱۰/۱۰", // تاریخ فارسی
		Category: "valid",
		Scenario: "تاریخ تولد به فارسی صحیح ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای موفق - کد 200 باید بازگردانده شود.",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       5,
		Content:  "04-15-2023", // تاریخ با فرمت ماه/روز/سال
		Category: "invalid",
		Scenario: "تاریخ تولد با فرمت اشتباه (ماه/روز/سال) ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای ناموفق - باید خطا در خروجی مشاهده شود.",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       6,
		Content:  "2001-12-31; DROP TABLE Users;", // SQL Injection
		Category: "invalid",
		Scenario: "حمله SQL Injection با درج دستور Drop Table",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای ناموفق - سیستم باید از این حمله جلوگیری کند.",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       7,
		Content:  "2001-12-31<script>alert('XSS')</script>", // XSS Attack
		Category: "invalid",
		Scenario: "حمله XSS با درج اسکریپت جاوا اسکریپت در تاریخ تولد",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای ناموفق - سیستم باید از حمله XSS جلوگیری کند.",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       8,
		Content:  "۲۰۲۳/۰۲/۳۲", // تاریخ اشتباه
		Category: "invalid",
		Scenario: "تاریخ تولد با روز اشتباه ارسال شده است (۳۲ روز)",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای ناموفق - باید خطا در خروجی مشاهده شود.",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       9,
		Content:  "1900-01-01", // تاریخ قدیمی
		Category: "valid",
		Scenario: "تاریخ تولد بسیار قدیمی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای موفق - کد 200 باید بازگردانده شود.",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       10,
		Content:  "2023-01-01<script>evilCode()</script>", // XSS with Code Injection
		Category: "invalid",
		Scenario: "حمله XSS و تزریق کد در تاریخ تولد",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای ناموفق - سیستم باید از حمله XSS جلوگیری کند.",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       11,
		Content:  "۱۳۷۵/۰۸/۲۵%20OR%201=1", // SQL Injection with URL encoding
		Category: "invalid",
		Scenario: "حمله SQL Injection با استفاده از URL encoding",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای ناموفق - باید از SQL Injection جلوگیری شود.",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       12,
		Content:  "12/31/2023", // تاریخ نادرست با فرمت آمریکایی
		Category: "invalid",
		Scenario: "تاریخ تولد نادرست با فرمت آمریکایی",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای ناموفق - باید خطا در خروجی مشاهده شود.",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       13,
		Content:  "2023-12-25<script>alert('xss')</script>", // XSS معیوب
		Category: "invalid",
		Scenario: "حمله XSS با اسکریپت جاوا اسکریپت در تاریخ تولد",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای ناموفق - جلوگیری از حملات XSS",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       14,
		Content:  "2024-02-29", // تاریخ معکوس سال کبیسه
		Category: "invalid",
		Scenario: "تاریخ تولد اشتباه با روز کبیسه ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای ناموفق - تاریخ اشتباه برای سال غیر کبیسه",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       15,
		Content:  "2024-01-01%20<script>malicious()</script>", // تزریق اسکریپت در تاریخ
		Category: "invalid",
		Scenario: "حمله XSS با تزریق اسکریپت در تاریخ تولد",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: جلوگیری از تزریق اسکریپت و XSS",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       16,
		Content:  "1990-11-30; DROP DATABASE;", // SQL Injection معکوس
		Category: "invalid",
		Scenario: "حمله SQL Injection به سیستم",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: جلوگیری از دستورات مخرب در سیستم",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       17,
		Content:  "2023/10/15' OR 1=1 --", // SQL Injection
		Category: "invalid",
		Scenario: "حمله SQL Injection با دستور OR 1=1",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: سیستم باید از حملات SQL Injection جلوگیری کند.",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       18,
		Content:  "2025-02-30", // تاریخ نادرست
		Category: "invalid",
		Scenario: "تاریخ تولد با روز نادرست ارسال شده است (۳۰ فوریه)",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: باید خطا در پاسخ باز گردد.",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       19,
		Content:  "2021-05-18<script>evilFunction()</script>", // تزریق کد JavaScript
		Category: "invalid",
		Scenario: "حمله تزریق کد جاوا اسکریپت",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: جلوگیری از کدهای مخرب جاوا اسکریپت",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       20,
		Content:  "2023-12-01<script>maliciousCode()</script>", // تزریق کد
		Category: "invalid",
		Scenario: "حمله تزریق کد در تاریخ تولد",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: جلوگیری از اجرای کد مخرب",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
}
