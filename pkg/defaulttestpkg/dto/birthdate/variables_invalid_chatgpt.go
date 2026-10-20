package defaulttestpkg

var VariablesInvalidChatGPT = []DTOBirthdateStruct{
	{
		Id:       1,
		Content:  "1970/01/01",
		Category: "invalid",
		Scenario: "تاریخ تولد با فرمت صحیح و به درستی ارسال شده است",
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
		Content:  "۱۹۷۰/۰۱/۰۱", // عدد فارسی
		Category: "invalid",
		Scenario: "تاریخ تولد با عدد فارسی ارسال شده است",
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
		Content:  "١٩٧٠/٠١/٠١", // عدد هندی
		Category: "invalid",
		Scenario: "تاریخ تولد با عدد هندی ارسال شده است",
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
		Content:  "1970-01-01", // تاریخ با فرمت درست
		Category: "valid",
		Scenario: "تاریخ تولد با فرمت صحیح ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       5,
		Content:  "31/12/1970", // فرمت تاریخ اشتباه
		Category: "invalid",
		Scenario: "تاریخ تولد با فرمت اشتباه ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید با فرمت 2006-01-02 باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       6,
		Content:  "1970-31-12", // ترتیب روز و ماه اشتباه
		Category: "invalid",
		Scenario: "تاریخ تولد با ترتیب نادرست روز و ماه ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید با فرمت 2006-01-02 باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       7,
		Content:  "2020-02-30", // تاریخ اشتباه
		Category: "invalid",
		Scenario: "تاریخ تولد با تاریخ نادرست ارسال شده است (ماه 2 بیشتر از 29 روز ندارد)",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       8,
		Content:  "2020-04-31", // تاریخ اشتباه
		Category: "invalid",
		Scenario: "تاریخ تولد با تاریخ نادرست ارسال شده است (ماه 4 بیشتر از 30 روز ندارد)",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       9,
		Content:  "2000-02-29", // سال کبیسه
		Category: "valid",
		Scenario: "تاریخ تولد با تاریخ صحیح در سال کبیسه ارسال شده است",
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
		Content:  "2021-02-29", // سال غیر کبیسه
		Category: "invalid",
		Scenario: "تاریخ تولد با تاریخ اشتباه در سال غیر کبیسه ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       11,
		Content:  "   ", // تاریخ خالی
		Category: "invalid",
		Scenario: "تاریخ تولد خالی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید پر شده باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       12,
		Content:  "2021-13-01", // ماه اشتباه
		Category: "invalid",
		Scenario: "تاریخ تولد با ماه اشتباه ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       13,
		Content:  "2021-00-01", // ماه صفر
		Category: "invalid",
		Scenario: "تاریخ تولد با ماه صفر ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       14,
		Content:  "2022-07-04", // تاریخ صحیح به زبان انگلیسی
		Category: "valid",
		Scenario: "تاریخ تولد به فرمت صحیح انگلیسی ارسال شده است",
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
		Content:  "2022-7-4", // فرمت اشتباه انگلیسی
		Category: "invalid",
		Scenario: "تاریخ تولد با فرمت اشتباه انگلیسی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید با فرمت 2006-01-02 باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       16,
		Content:  "2022-08-15", // تاریخ صحیح
		Category: "valid",
		Scenario: "تاریخ تولد به فرمت صحیح ارسال شده است",
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
		Content:  "2020-05-25", // تاریخ صحیح
		Category: "valid",
		Scenario: "تاریخ تولد به فرمت صحیح و معتبر ارسال شده است",
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
		Content:  "2023-04-31", // روز اشتباه
		Category: "invalid",
		Scenario: "تاریخ تولد با روز اشتباه ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       19,
		Content:  "2022-11-31", // روز اشتباه
		Category: "invalid",
		Scenario: "تاریخ تولد با روز اشتباه ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       20,
		Content:  "2021-09-01", // تاریخ صحیح
		Category: "valid",
		Scenario: "تاریخ تولد به فرمت صحیح و معتبر ارسال شده است",
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
		Content:  "2025-02-29", // تاریخ اشتباه در سال غیر کبیسه
		Category: "invalid",
		Scenario: "تاریخ تولد با تاریخ اشتباه در سال غیر کبیسه ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       22,
		Content:  "2024-13-01", // ماه اشتباه
		Category: "invalid",
		Scenario: "تاریخ تولد با ماه غیرمعتبر ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       23,
		Content:  "2022-11-00", // روز صفر
		Category: "invalid",
		Scenario: "تاریخ تولد با روز صفر ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       24,
		Content:  "2022-02-31", // روز اشتباه در ماه فوریه
		Category: "invalid",
		Scenario: "تاریخ تولد با روز غیرمعتبر در ماه فوریه ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       25,
		Content:  "2025-04-01", // تاریخ صحیح
		Category: "valid",
		Scenario: "تاریخ تولد به فرمت صحیح ارسال شده است",
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
		Content:  "1999-12-31", // تاریخ قدیمی
		Category: "valid",
		Scenario: "تاریخ تولد به فرمت صحیح و معتبر ارسال شده است",
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
		Content:  "2023-01-01", // سال جدید
		Category: "valid",
		Scenario: "تاریخ تولد در ابتدای سال جدید ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       28,
		Content:  "2022-07-15", // تاریخ معتبر
		Category: "valid",
		Scenario: "تاریخ تولد به فرمت صحیح و معتبر ارسال شده است",
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
		Content:  "1996-08-24", // تاریخ معتبر در یک سال قدیمی
		Category: "valid",
		Scenario: "تاریخ تولد با فرمت صحیح در یک سال قدیمی ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       30,
		Content:  "2022-02-29", // تاریخ اشتباه در سال غیر کبیسه
		Category: "invalid",
		Scenario: "تاریخ تولد با تاریخ اشتباه در سال غیر کبیسه ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       31,
		Content:  "1998-09-31", // روز اشتباه
		Category: "invalid",
		Scenario: "تاریخ تولد با روز اشتباه ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       32,
		Content:  "2023-03-32", // روز بیشتر از 31
		Category: "invalid",
		Scenario: "تاریخ تولد با روز بیشتر از 31 ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       33,
		Content:  "2020-11-00", // روز صفر
		Category: "invalid",
		Scenario: "تاریخ تولد با روز صفر ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       34,
		Content:  "2023-06-31", // روز اشتباه در ماه ژوئن
		Category: "invalid",
		Scenario: "تاریخ تولد با روز اشتباه در ماه ژوئن ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       35,
		Content:  "2025-10-03", // تاریخ صحیح
		Category: "valid",
		Scenario: "تاریخ تولد به فرمت صحیح ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       36,
		Content:  "2020-05-32", // روز بیشتر از 31
		Category: "invalid",
		Scenario: "تاریخ تولد با روز بیشتر از 31 ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       37,
		Content:  "2021-03-31", // تاریخ صحیح
		Category: "valid",
		Scenario: "تاریخ تولد به فرمت صحیح ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       38,
		Content:  "2022-10-32", // روز اشتباه
		Category: "invalid",
		Scenario: "تاریخ تولد با روز اشتباه در ماه اکتبر ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       39,
		Content:  "2023-12-25", // تاریخ معتبر
		Category: "valid",
		Scenario: "تاریخ تولد به فرمت صحیح ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: اجرای سالم - در خروجی کد 200 برگشت داده بشود",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       40,
		Content:  "2021-02-29", // تاریخ اشتباه در سال غیر کبیسه
		Category: "invalid",
		Scenario: "تاریخ تولد با تاریخ اشتباه در سال غیر کبیسه ارسال شده است",
		Provider: "ChatGPT",
		Expected: Expected{
			Description: "انتظار ما: تاریخ تولد باید معتبر باشد - کد 400 برگشت داده بشود",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
}
