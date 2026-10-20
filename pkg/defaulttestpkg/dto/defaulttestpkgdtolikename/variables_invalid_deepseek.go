package defaulttestpkgdtolikename

import "strings"

//var variablesInvalidDeepSeek = []myStruct{
//	{
//		Id:       69,
//		Content:  "𐐧𐐨𐑅𐐻𐐮𐑍", // Deseret script
//		Category: "invalid",
//		Scenario: "سیستم نوشتاری دیزرت (مورمونی)",
//		Provider: "DeepSeek",
//		Expected: expected{
//			Description: "اجرای سالم - کد 200",
//			Status:      true,
//			StatusCode:  200,
//			Message:     nil,
//		},
//	},
//}

var variablesInvalidDeepSeek = []myStruct{
	{
		Id:       1,
		Content:  "Ali",
		Category: "invalid",
		Scenario: "نام انگلیسی معتبر با 3 کاراکتر",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       2,
		Content:  "اب",
		Category: "invalid",
		Scenario: "رشته فارسی کوتاه (2 کاراکتر)",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       3,
		Content:  "",
		Category: "invalid",
		Scenario: "رشته خالی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       4,
		Content:  strings.Repeat("a", 257),
		Category: "invalid",
		Scenario: "رشته با 257 کاراکتر",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       5,
		Content:  "John Michael Alexander William Smith ok",
		Category: "invalid",
		Scenario: "6 کلمه در رشته انگلیسی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       6,
		Content:  "محمد رضا علی",
		Category: "invalid",
		Scenario: "3 کلمه فارسی معتبر",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       7,
		Content:  "Li Wei 张伟",
		Category: "invalid",
		Scenario: "ترکیب چینی و انگلیسی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       8,
		Content:  "山田太郎_Yamada",
		Category: "invalid",
		Scenario: "ترکیب ژاپنی و انگلیسی با کاراکترهای خاص",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       9,
		Content:  "😊Emma_Watson",
		Category: "invalid",
		Scenario: "حاوی ایموجی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       10,
		Content:  "Иван Петров",
		Category: "invalid",
		Scenario: "روسی با کاراکترهای سیریلیک",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       11,
		Content:  `अनिल कुमार`,
		Category: "invalid",
		Scenario: "هندی با دیواناگری",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       12,
		Content:  "Jean-François",
		Category: "invalid",
		Scenario: "کاراکتر خاص فرانسوی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       13,
		Content:  "  Trim  Spaces  ",
		Category: "invalid",
		Scenario: "فاصله های اضافه در ابتدا و انتها",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200 (بعد از Trim)",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       14,
		Content:  "A",
		Category: "invalid",
		Scenario: "تک کاراکتر لاتین",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       15,
		Content:  strings.Repeat("あ", 255),
		Category: "invalid",
		Scenario: "رشته ژاپنی با 255 کاراکتر",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       16,
		Content:  "12345",
		Category: "invalid",
		Scenario: "اعداد انگلیسی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       17,
		Content:  "١٢٣٤٥",
		Category: "invalid",
		Scenario: "اعداد عربی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       18,
		Content:  "John@Doe#",
		Category: "invalid",
		Scenario: "کاراکترهای ویژه انگلیسی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       19,
		Content:  "控制代码测试",
		Category: "invalid",
		Scenario: "کاراکترهای چینی معتبر",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       20,
		Content:  "\x00\x01\x02",
		Category: "invalid",
		Scenario: "کاراکترهای غیرقابل چاپ",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       21,
		Content:  "محمد رضا علی حسنی", // 5 کلمه فارسی
		Category: "invalid",
		Scenario: "رشته فارسی با حداکثر تعداد کلمات مجاز (5 کلمه)",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       22,
		Content:  strings.Repeat("あ", 255), // ژاپنی 255 کاراکتری
		Category: "invalid",
		Scenario: "رشته ژاپنی با حداکثر طول مجاز (255 کاراکتر)",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       23,
		Content:  "A\u200B\u200B\u200B", // کاراکترهای Zero-Width Space
		Category: "invalid",
		Scenario: "کاراکترهای نامرئی تکراری در متن",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       24,
		Content:  "नमस्ते_你好_مرحبا", // ترکیب هندی، چینی و عربی
		Category: "invalid",
		Scenario: "ترکیب سه زبان مختلف در یک رشته",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       25,
		Content:  "\u202ACyber\u202C", // Right-to-Left Override
		Category: "invalid",
		Scenario: "کاراکترهای کنترل جهت نوشتار RTL",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       26,
		Content:  "L̷̔̑̈́ŏ̴̓̎r̸͌̔͝e̷̕̚͝ṁ̵̌", // Zalgo Text پیشرفته
		Category: "invalid",
		Scenario: "متن زالگو با کاراکترهای ترکیبی چندلایه",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       27,
		Content:  "١٢٣٤٥٦", // اعداد عربی-هندی
		Category: "invalid",
		Scenario: "اعداد شرقی با طول 6 کاراکتر",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       28,
		Content:  "𐌀𐌁𐌂𐌃𐌄𐌅", // ایتالیایی باستان
		Category: "invalid",
		Scenario: "خط ایتالیایی باستان با کاراکترهای پیوسته",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       29,
		Content:  "AͦͦͦBͦͦͦCͦͦͦ", // تکرار نشانگرهای ترکیبی
		Category: "invalid",
		Scenario: "ترکیب نشانگرهای تکراری روی حروف",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       30,
		Content:  "ភាសាខ្មែរ", // خمری (Cambodian)
		Category: "invalid",
		Scenario: "نوشتار خمری (کامبوجی)",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       31,
		Content:  "أيُّهَا", // عربی با همزه و حرکات
		Category: "invalid",
		Scenario: "عربی با نشانه‌های ترکیبی پیشرفته",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       32,
		Content:  "𝄞\uFE0E", // Variation Selector
		Category: "invalid",
		Scenario: "نماد موسیقی با انتخابگر واریانت",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       33,
		Content:  "᚛ᚇᚔᚈᚐ᚜", // Ogham متن بسته‌بندی شده
		Category: "invalid",
		Scenario: "متن اوگام با محدوده‌های مرزی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       34,
		Content:  "A\u035C\u035D\u035EB", // Combining Double Marks
		Category: "invalid",
		Scenario: "نشانگرهای ترکیبی دوبل",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       35,
		Content:  "꯱꯲꯳꯴꯵", // Meetei Mayek اعداد
		Category: "invalid",
		Scenario: "اعداد میتئی مایک (هند شمال شرقی)",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       36,
		Content:  "𓂀𓃀𓄀𓅀𓆀", // هیروگلیف مصری
		Category: "invalid",
		Scenario: "هیروگلیف‌های مصری ترکیبی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       37,
		Content:  "𝌆\u0301\u0302", // ترکیب نماد و نشانگر
		Category: "invalid",
		Scenario: "نماد هگزاگرام با نشانگرهای ترکیبی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       38,
		Content:  "𐎭𐎮𐎯𐎰𐎱", // خط میخی اکدی
		Category: "invalid",
		Scenario: "خط میخی اکدی با کاراکترهای پیچیده",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       39,
		Content:  "A\u1DFF\u1DFE", // Combining Diacritical Marks Supplement
		Category: "invalid",
		Scenario: "نشانگرهای ترکیبی پیشرفته",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       40,
		Content:  "𝓐𝓑𝓒𝓓𝓔𝓕", // ریاضی اسکریپت
		Category: "invalid",
		Scenario: "کاراکترهای اسکریپت ریاضی پیوسته",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       41,
		Content:  "\u202Bمحمد\u202C", // RTL embedding
		Category: "invalid",
		Scenario: "کاراکترهای کنترل جهت نوشتار در متن فارسی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       42,
		Content:  "Áéíóú", // ترکیب کاراکترهای اکسان دار
		Category: "invalid",
		Scenario: "کاراکترهای ترکیبی با علامت‌های سجاوندی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       43,
		Content:  "𐏐𐏑𐏒𐏓𐏔", // Old Persian numbers
		Category: "invalid",
		Scenario: "کاراکترهای عددی فارسی باستان",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       44,
		Content:  "ᚠᚡᚢᚣᚤ", // Runic characters
		Category: "invalid",
		Scenario: "کاراکترهای رونی (خط قدیمی ژرمنی)",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       45,
		Content:  "ᠤᠯᠠᠭᠠᠨᠳᠠᠢ", // Mongolian with multiple suffixes
		Category: "invalid",
		Scenario: "رشته مغولی با پسوندهای چندگانه",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       46,
		Content:  "𝓐𝓑𝓒𝓓𝓔", // Mathematical bold script
		Category: "invalid",
		Scenario: "کاراکترهای اسکریپت ریاضی پررنگ",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       47,
		Content:  "⡍⠔⠊⠇ ⠞⠑⠭⠞", // Braille pattern
		Category: "invalid",
		Scenario: "متن بریل",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 400",
			Status:      true,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       48,
		Content:  "㋐㋑㋒㋓㋔", // Circled Katakana
		Category: "invalid",
		Scenario: "کاراکترهای کاتاکانای محصورشده",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       49,
		Content:  "అనుచరుల", // Telugu script
		Category: "invalid",
		Scenario: "نوشتار تلگو (هند جنوبی)",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       50,
		Content:  "꧅ꦱꦸꦏꦱ꧅", // Javanese punctuation
		Category: "invalid",
		Scenario: "نشانه‌گذاری جاوه‌ای",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       51,
		Content:  "𑁍𑁯𑂠𑂰𑂫𑂰𑂩", // Brahmi script
		Category: "invalid",
		Scenario: "خط براهمی (قدیمی هندی)",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       52,
		Content:  "ᜀᜃᜓᜂᜅ᜔", // Tagalog script
		Category: "invalid",
		Scenario: "نوشتار تاگالوگ (فیلیپین)",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       53,
		Content:  "𐒰𐒻𐒷𐒻𐒼", // Osmanya script
		Category: "invalid",
		Scenario: "خط عثمانی (سومالی)",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       54,
		Content:  "𐎀𐎁𐎂𐎃𐎄", // Ugaritic cuneiform
		Category: "invalid",
		Scenario: "خط میخی اوگاریتی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       55,
		Content:  "𒍑𒍑𒍑𒍑𒍑", // Repeated cuneiform
		Category: "invalid",
		Scenario: "تکرار کاراکترهای میخی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       56,
		Content:  "𝄀𝄁𝄂𝄃𝄄", // Musical symbols
		Category: "invalid",
		Scenario: "نمادهای موسیقی یونی‌کد",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       57,
		Content:  "󠀡󠀢󠀣", // Tag characters
		Category: "invalid",
		Scenario: "کاراکترهای تگ یونی‌کد",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       58,
		Content:  "𓀀𓀁𓀂𓀃𓀄", // Egyptian hieroglyphs
		Category: "invalid",
		Scenario: "هیروگلیف‌های مصری متوالی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       59,
		Content:  "\uE0041\uE0042\uE0043", // Language tag characters
		Category: "invalid",
		Scenario: "کاراکترهای تگ زبان",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       60,
		Content:  "ＨＥＬＬＯ", // Fullwidth Latin
		Category: "invalid",
		Scenario: "حروف لاتین فول‌ویلد",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       61,
		Content:  "A\u200eB\u200fC", // ترکیب جهت‌های متناقض
		Category: "invalid",
		Scenario: "ترکیب کاراکترهای کنترل جهت RTL و LTR",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       62,
		Content:  "Ḩ̵̛͖̭͋͜͟E̸̝͈͌̅̎͢͞L̴̡͇̻̘̽̿͑͜L̵̳̰̦͊̕͟͝O̷̳̤̺̝̾̔̊", // Zalgo Text
		Category: "invalid",
		Scenario: "متن زالگو با کاراکترهای ترکیبی بیش از حد",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       63,
		Content:  "ρ⍺у𝓅𝒶у", // Homoglyph حمله
		Category: "invalid",
		Scenario: "کاراکترهای همگلیف برای فیشینگ",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       64,
		Content:  "A\u034F\u035F\u0360B", // Combining graphemes
		Category: "invalid",
		Scenario: "ترکیب نشانگرهای گرافیمی پیچیده",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       65,
		Content:  "ᴮᴵᴿᵀᴴᴰᴬᵞ", // Small caps با اعداد
		Category: "invalid",
		Scenario: "حروف کوچک سرمایه با اعداد یونی‌کد",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       66,
		Content:  "𝔸+𝔹=ℂ", // نمادهای ریاضی
		Category: "invalid",
		Scenario: "معادله ریاضی با کاراکترهای یونی‌کد",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       67,
		Content:  "\uFDEF", // Arabic ligature نادر
		Category: "invalid",
		Scenario: "لیگاتور عربی نادر یونی‌کد",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       68,
		Content:  "🇨🇦🇺🇸\u200D🇲🇽", // ترکیب پرچم‌ها با ZWJ
		Category: "invalid",
		Scenario: "ترکیب غیرممکن پرچم‌ها",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       69,
		Content:  "𐐧𐐨𐑅𐐻𐐮𐑍", // Deseret script
		Category: "invalid",
		Scenario: "سیستم نوشتاری دیزرت (مورمونی)",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       70,
		Content:  "A\u0F3B\u0F3DB", // Tibetan digits
		Category: "invalid",
		Scenario: "ترکیب تبتانی و لاتین",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	//{
	//	Id:       71,
	//	Content:  "\uD800\uDC00", // Surrogate pair (𐀀)
	//	Category: "invalid",
	//	Scenario: "جفت جایگزین یونی‌کد",
	//	Provider: "DeepSeek",
	//	Expected: expected{
	//		Description: "اجرای سالم - کد 200",
	//		Status:      true,
	//		StatusCode:  200,
	//		Message:     nil,
	//	},
	//},
	{
		Id:       72,
		Content:  "A͊̚̕͢͜Bͦ̒ͣ͠C͌ͥͦ͏", // Stacked diacritics
		Category: "invalid",
		Scenario: "کاراکترهای ترکیبی چندلایه",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       73,
		Content:  "ᏣᎳᎩ ᎦᏬᏂᎯᏍᏗ", // Cherokee script
		Category: "invalid",
		Scenario: "زبان چروکی (سرخپوستی)",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       74,
		Content:  "A\u2061B\u2062C\u2063", // Invisible operators
		Category: "invalid",
		Scenario: "عملگرهای نامرئی ریاضی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       75,
		Content:  "𝕏ⅅⅆⅇⅈⅉ", // Mathematical symbols
		Category: "invalid",
		Scenario: "نمادهای ریاضی پیشرفته",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       76,
		Content:  "ᠠᠰᠢᠷᠠᠬᠤ", // Mongolian with suffixes
		Category: "invalid",
		Scenario: "مغولی با پسوندهای دستوری",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       77,
		Content:  "キ\u3099ャ\u3099", // Combining kana
		Category: "invalid",
		Scenario: "کاراکترهای کانا ترکیبی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       78,
		Content:  "\uE0100\uE0101\uE0102", // Variation selectors
		Category: "invalid",
		Scenario: "انتخابگرهای واریانت یونی‌کد",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       79,
		Content:  "ȺⱥȾⱦ", // Latin extensions
		Category: "invalid",
		Scenario: "حروف توسعه یافته لاتین",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       80,
		Content:  "𝌆𝌇𝌈𝌉", // Hexagram symbols
		Category: "invalid",
		Scenario: "نمادهای هگزاگرام آی چینگ",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       81,
		Content:  "𐎠𐎡𐎢𐎣𐎤𐎥", // Old Persian Cuneiform
		Category: "invalid",
		Scenario: "خط میخی فارسی باستان با کاراکترهای پیوسته",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       82,
		Content:  "A\u036F\u034F\u035F", // Combining marks overflow
		Category: "invalid",
		Scenario: "تراکم بیش از حد نشانگرهای ترکیبی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       83,
		Content:  "ᛖᛏᚢ ᛟᚠ ᚦᛖ ᚱᚢᚾᛖᛋ", // Elder Futhark Runes
		Category: "invalid",
		Scenario: "خط رونی فوتارک قدیمی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       84,
		Content:  "𝌖𝌗𝌘𝌙𝌚", //易经六线形 ترکیبی
		Category: "invalid",
		Scenario: "ترکیب نمادهای هگزاگرام آی چینگ",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       85,
		Content:  "𖠀𖡀𖢀𖣀𖤀", // Anatolian Hieroglyphs
		Category: "invalid",
		Scenario: "هیروگلیف آناتولی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       86,
		Content:  "A\u0E3F\u0E4F\u0E5F", // Thai combining marks
		Category: "invalid",
		Scenario: "نشانگرهای ترکیبی تایلندی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       87,
		Content:  "𝄫𝄬𝄭𝄮𝄯", // Byzantine Musical Symbols
		Category: "invalid",
		Scenario: "نمادهای موسیقی بیزانسی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       88,
		Content:  "𒁱𒈬𒀭𒆠𒋛", // Sumero-Akkadian Cuneiform
		Category: "invalid",
		Scenario: "خط میخی سومری-اکدی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       89,
		Content:  "𐌀𐌂𐌃𐌄𐌅𐌆𐌇", // Old Italic
		Category: "invalid",
		Scenario: "خط ایتالیایی باستان",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       90,
		Content:  "A⃰B⃰C⃰", // Combining asterisks
		Category: "invalid",
		Scenario: "ستاره‌های ترکیبی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       91,
		Content:  "ꜢꜣꜥꜦꜨ", // Egyptian transliteration
		Category: "invalid",
		Scenario: "حروف انتقال مصری",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       92,
		Content:  "𝣵𝣲𝣱𝣰𝣴", // SignWriting
		Category: "invalid",
		Scenario: "نمادهای نوشتاری زبان اشاره",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       93,
		Content:  "ᎤᏔᎾᏅᏍᎬ", // Cherokee advanced
		Category: "invalid",
		Scenario: "جمله کامل چروکی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       94,
		Content:  "A\uFFA0\uFFB0\uFFC0", // Halfwidth forms
		Category: "invalid",
		Scenario: "کاراکترهای نیم‌عرض",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       95,
		Content:  "𑁍𑁎𑁏𑁐𑁑", // Brahmi numerals
		Category: "invalid",
		Scenario: "اعداد براهمی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       96,
		Content:  "𝈆𝈇𝈈𝈉𝈊", // Ancient Greek Music
		Category: "invalid",
		Scenario: "نمادهای موسیقی یونان باستان",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
	{
		Id:       97,
		Content:  "𑀓𑀔𑀕𑀖𑀗", // Brahmi letters
		Category: "invalid",
		Scenario: "حروف براهمی کلاسیک",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       98,
		Content:  "A\u1ABE\u1ABF\u1AC0", // Combining diacritics extended
		Category: "invalid",
		Scenario: "نشانگرهای ترکیبی توسعه یافته",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       99,
		Content:  "𓎡𓎢𓎣𓎤𓎥", // Egyptian hieroglyph sequence
		Category: "invalid",
		Scenario: "دنباله هیروگلیف مصری",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "اجرای سالم - کد 200",
			Status:      true,
			StatusCode:  200,
			Message:     nil,
		},
	},
	{
		Id:       100,
		Content:  "𝍠𝍡𝍢𝍣𝍤", // Counting Rod Numerals
		Category: "invalid",
		Scenario: "اعداد چوب شمارش چینی",
		Provider: "DeepSeek",
		Expected: expected{
			Description: "خطا - کد 400",
			Status:      false,
			StatusCode:  400,
			Message:     nil,
		},
	},
}
