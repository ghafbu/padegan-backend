package defaulttestpkgdtobirthdate

var variablesInvalidDeepSeek = []myStruct{
	{
		Id:       1,
		Content:  "۱۳۹۵-۰۸-۲۵", // فارسی numbers
		Category: "invalid",
		Scenario: "استفاده از اعداد فارسی در تاریخ",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       2,
		Content:  "२०२४-१२-०१", // Hindi numbers
		Category: "invalid",
		Scenario: "استفاده از اعداد هندی در تاریخ",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       3,
		Content:  "2024/12/01", // Slash separator
		Category: "invalid",
		Scenario: "استفاده از اسلش به جای خط تیره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       4,
		Content:  "1403-15-99", // Invalid month/day
		Category: "invalid",
		Scenario: "ماه و روز نامعتبر در تقویم شمسی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       5,
		Content:  "2024-02-30", // Invalid Gregorian date
		Category: "invalid",
		Scenario: "تاریخ نامعتبر میلادی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       6,
		Content:  "0000-00-00", // All zeros
		Category: "invalid",
		Scenario: "مقادیر صفر در تمام بخشها",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       7,
		Content:  "2024-13-01", // Month > 12
		Category: "invalid",
		Scenario: "ماه بیشتر از 12",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       8,
		Content:  "2024-12-32", // Day > 31
		Category: "invalid",
		Scenario: "روز بیشتر از 31",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       9,
		Content:  "2024-12-1", // Short length (9 chars)
		Category: "invalid",
		Scenario: "طول رشته کمتر از 10 کاراکتر",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       10,
		Content:  "2024-12-012", // Long length (11 chars)
		Category: "invalid",
		Scenario: "طول رشته بیشتر از 10 کاراکتر",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       11,
		Content:  "۱۴۰۳-۰۲-۳۰", // Persian numbers with valid format
		Category: "invalid",
		Scenario: "اعداد فارسی با فرمت صحیح",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       12,
		Content:  "2024-AB-CD", // Alphabetic characters
		Category: "invalid",
		Scenario: "استفاده از حروف لاتین به جای اعداد",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       13,
		Content:  "2024-۱۲-۰۱", // Mixed English/Persian
		Category: "invalid",
		Scenario: "ترکیب اعداد انگلیسی و فارسی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       14,
		Content:  "2024.12.01", // Dot separator
		Category: "invalid",
		Scenario: "استفاده از نقطه به جای خط تیره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       15,
		Content:  "  2024-12-01  ", // Whitespace
		Category: "invalid",
		Scenario: "فاصله خالی در ابتدا و انتها",
		Provider: "DeepSeek",
		Expected: responseCode200("2024-12-01", nil),
	},
	{
		Id:       16,
		Content:  "1445-09-30", // Hijri valid date
		Category: "valid",
		Scenario: "تاریخ معتبر هجری قمری",
		Provider: "DeepSeek",
		Expected: responseCode200("", nil),
	},
	{
		Id:       17,
		Content:  "1989-XX-YY", // Invalid characters
		Category: "invalid",
		Scenario: "کاراکترهای غیرعددی نامشخص",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       18,
		Content:  "1403-02-31", // Invalid Persian date
		Category: "invalid",
		Scenario: "تاریخ نامعتبر شمسی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       19,
		Content:  "二零二四-十二-零一", // Chinese numbers
		Category: "invalid",
		Scenario: "استفاده از اعداد چینی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       20,
		Content:  "2024-12-01T00:00:00", // ISO format
		Category: "invalid",
		Scenario: "فرمت زمان همراه با تاریخ",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       21,
		Content:  "1700-02-29", // Invalid leap year (Gregorian)
		Category: "invalid",
		Scenario: "سال کبیسه نامعتبر میلادی قبل از اصلاح تقویم",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       22,
		Content:  "۱۴۰۳-۱٣-٢۵", // Mixed Persian/Arabic numerals
		Category: "invalid",
		Scenario: "ترکیب اعداد فارسی و عربی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       23,
		Content:  "2024-12-01 ", // Trailing space
		Category: "invalid",
		Scenario: "فاصله خالی در انتهای رشته",
		Provider: "DeepSeek",
		Expected: responseCode200("2024-12-01", nil),
	},
	{
		Id:       24,
		Content:  "۳۰۵۰-۰۱-۰۱", // Future Persian date
		Category: "invalid",
		Scenario: "تاریخ شمسی بسیار آینده",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       25,
		Content:  "0001-01-01", // Minimal valid date
		Category: "valid",
		Scenario: "قدیمیترین تاریخ معتبر ممکن",
		Provider: "DeepSeek",
		Expected: responseCode200("", nil),
	},
	{
		Id:       26,
		Content:  "2024-12-01-", // Extra hyphen at end
		Category: "invalid",
		Scenario: "خط تیره اضافی در انتها",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       27,
		Content:  "1991.08.25", // Dots with valid date
		Category: "invalid",
		Scenario: "نقطه به عنوان جداکننده با تاریخ معتبر",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       28,
		Content:  "۱۴۰۳/۰۷/۰۶", // Persian numbers with slash
		Category: "invalid",
		Scenario: "اعداد فارسی با جداکننده اسلش",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       29,
		Content:  "2024\u201312\u201301", // Using en-dash unicode
		Category: "invalid",
		Scenario: "استفاده از خط تیره یونیکد (en-dash)",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       30,
		Content:  "2567-05-15", // Buddhist calendar date
		Category: "valid",
		Scenario: "تاریخ معتبر بودایی",
		Provider: "DeepSeek",
		Expected: responseCode200("", nil),
	},
	{
		Id:       31,
		Content:  "13AB-CD-EF", // Hexadecimal-like format
		Category: "invalid",
		Scenario: "فرمت شبه هگزادسیمال",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       32,
		Content:  "2024-12-01T12:00", // ISO datetime without timezone
		Category: "invalid",
		Scenario: "فرمت تاریخ-زمان بدون تایم زون",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       33,
		Content:  "२०२४-१२-०१", // Devanagari (Hindi) numbers
		Category: "invalid",
		Scenario: "اعداد دواناگری (هندی)",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       34,
		Content:  "2024=12=01", // Equal sign separator
		Category: "invalid",
		Scenario: "جداکننده علامت مساوی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       35,
		Content:  "01-12-2024", // Day-month-year format
		Category: "invalid",
		Scenario: "فرت روز-ماه-سال",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       36,
		Content:  "2024-12-01.", // Trailing dot
		Category: "invalid",
		Scenario: "نقطه در انتهای رشته",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       37,
		Content:  "2024\uFF0D12\uFF0D01", // Full-width hyphens
		Category: "invalid",
		Scenario: "خط تیره با عرض کامل (یونیکد)",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       38,
		Content:  "۱۴۰۳-۰۷-۳۲", // Invalid Persian day
		Category: "invalid",
		Scenario: "روز نامعتبر شمسی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       39,
		Content:  "2024-12-001", // Three-digit day
		Category: "invalid",
		Scenario: "روز سه رقمی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       40,
		Content:  "2024-۱۲-۰۱", // Mixed English-Persian numbers
		Category: "invalid",
		Scenario: "ترکیب اعداد انگلیسی و فارسی در یک تاریخ",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       41,
		Content:  "۰۹۸۷-۰۶-۰۴", // Eastern Arabic numerals
		Category: "invalid",
		Scenario: "استفاده از اعداد عربی شرقی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       42,
		Content:  "2024\u200C12\u200C01", // Zero-width non-joiner
		Category: "invalid",
		Scenario: "کاراکترهای غیرمرئی بین اعداد",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       43,
		Content:  "令和6-12-01", // Japanese era format
		Category: "invalid",
		Scenario: "فرمت تاریخ ژاپنی با حروف کانجی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       44,
		Content:  "2024-12-01\u202E", // RTL override character
		Category: "invalid",
		Scenario: "کاراکتر جهت‌دهی راست به چپ در انتها",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       45,
		Content:  "1٬2٬3٬4-12-01", // Arabic thousands separator
		Category: "invalid",
		Scenario: "استفاده از جداکننده هزارگان عربی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       46,
		Content:  "2024-12-01?", // Special character suffix
		Category: "invalid",
		Scenario: "کاراکتر ویژه در انتهای رشته",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       47,
		Content:  "۱۴۰۳/۰۷/۳۱", // Persian with slash (valid date)
		Category: "invalid",
		Scenario: "فرمت نامعتبر با تاریخ شمسی واقعی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       48,
		Content:  "2024-Ad-01", // Month as abbreviation
		Category: "invalid",
		Scenario: "ماه به صورت حروف اختصاری",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       49,
		Content:  "+2024-12-01", // ISO year extension
		Category: "invalid",
		Scenario: "سال با پیشوند مثبت",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       50,
		Content:  "2024-12-1a", // Alphanumeric suffix
		Category: "invalid",
		Scenario: "ترکیب عدد و حرف در روز",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       51,
		Content:  "१४०३-०७-३१", // Devanagari (Sanskrit)
		Category: "invalid",
		Scenario: "اعداد سانسکریت",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       52,
		Content:  "2024-12-01 ", // Leading space
		Category: "invalid",
		Scenario: "فاصله خالی در ابتدای رشته",
		Provider: "DeepSeek",
		Expected: responseCode200("2024-12-01", nil),
	},
	{
		Id:       53,
		Content:  "2024\u201312\u201301", // Em-dash separator
		Category: "invalid",
		Scenario: "استفاده از em-dash یونیکد",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       54,
		Content:  "2024_12_01", // Underscore separator
		Category: "invalid",
		Scenario: "جداکننده زیرخطی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       55,
		Content:  "2024-闰-01", // Chinese leap month
		Category: "invalid",
		Scenario: "ماه چینی با نماد کبیسه",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       56,
		Content:  "2024-12-001", // Three-digit day (zero-padded)
		Category: "invalid",
		Scenario: "روز سه رقمی با صفر پیشرو",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       57,
		Content:  "2024-12-1", // Missing leading zero in day
		Category: "invalid",
		Scenario: "روز بدون صفر پیشرو",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       58,
		Content:  "1403-7-6", // Persian without leading zeros
		Category: "invalid",
		Scenario: "شمسی بدون صفر پیشرو در ماه/روز",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       59,
		Content:  "2024-00-00", // Zero month and day
		Category: "invalid",
		Scenario: "ماه و روز صفر در میلادی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       60,
		Content:  "2024-12-01 ", // Non-breaking space (U+00A0)
		Category: "invalid",
		Scenario: "فاصله غیرشکستنی در انتها",
		Provider: "DeepSeek",
		Expected: responseCode200("2024-12-01", nil),
	},
	{
		Id:       61,
		Content:  "MMXXIV-XII-01", // Roman numerals
		Category: "invalid",
		Scenario: "استفاده از اعداد رومی در تمام بخش‌ها",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       62,
		Content:  "2024--12--01", // Double hyphens
		Category: "invalid",
		Scenario: "جداکننده‌های تکراری خط تیره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       63,
		Content:  "٢٠٢۴-۱٢-٠۱", // Mixed Arabic numeral styles
		Category: "invalid",
		Scenario: "ترکیب شیوه‌های مختلف عددنویسی عربی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       64,
		Content:  "2024-12-01\t", // Tab character suffix
		Category: "invalid",
		Scenario: "کاراکتر تب در انتهای رشته",
		Provider: "DeepSeek",
		Expected: responseCode200("2024-12-01", nil),
	},
	{
		Id:       65,
		Content:  "2024~12~01", // Tilde separators
		Category: "invalid",
		Scenario: "جداکننده‌های تیلدا",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       66,
		Content:  "1403-13-32", // Exceed Persian calendar limits
		Category: "invalid",
		Scenario: "مقادیر فراتر از محدودیت‌های تقویم شمسی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       67,
		Content:  "2024-12-01.", // Leading dot (invalid position)
		Category: "invalid",
		Scenario: "نقطه در ابتدای رشته",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       68,
		Content:  "២០២៤-១២-០១", // Khmer numerals
		Category: "invalid",
		Scenario: "اعداد خمری (کامبوجی)",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       69,
		Content:  "2024-１２-０１", // Full-width numbers
		Category: "invalid",
		Scenario: "اعداد با عرض کامل یونیکد",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       70,
		Content:  "00140-003-00005", // Excessive zero-padding
		Category: "invalid",
		Scenario: "صفرهای پیشرو بیش از حد در تمام بخش‌ها",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       71,
		Content:  "2024-Jun-01", // Mixed numeric/alphanumeric
		Category: "invalid",
		Scenario: "ترکیب عددی-حروفی در بخش ماه",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       72,
		Content:  "2024-12-01 ", // Non-printable control character
		Category: "invalid",
		Scenario: "کاراکتر کنترل غیرقابل چاپ",
		Provider: "DeepSeek",
		Expected: responseCode200("2024-12-01", nil),
	},
	{
		Id:       73,
		Content:  "2024|12|01", // Pipe separators
		Category: "invalid",
		Scenario: "جداکننده‌های عمودی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       74,
		Content:  "-2024-12-01", // Negative year prefix
		Category: "invalid",
		Scenario: "سال منفی با پیشوند خط تیره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       75,
		Content:  "۲٬۰۲۴-۱۲-۰۱", // Persian with comma separator
		Category: "invalid",
		Scenario: "جداکننده هزارگان فارسی در سال",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       76,
		Content:  "2024-12-01 ", // Zero-width space
		Category: "invalid",
		Scenario: "فاصله صفر عرض یونیکد",
		Provider: "DeepSeek",
		Expected: responseCode200("2024-12-01", nil),
	},
	{
		Id:       77,
		Content:  "𝟮𝟬𝟮𝟰-𝟭𝟮-𝟬𝟭", // Mathematical bold digits
		Category: "invalid",
		Scenario: "اعداد ریاضی پررنگ یونیکد",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       78,
		Content:  "2024\u202E-12\u202C-01", // Hidden RTL marks
		Category: "invalid",
		Scenario: "نشانه‌های جهت‌دهی راست به چپ مخفی در متن",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       79,
		Content:  "۲۰۲۴۱۲۰۱", // Compact no-separator format
		Category: "invalid",
		Scenario: "فرمت فشرده بدون جداکننده",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       80,
		Content:  "2\u03360\u03362\u03364\u0336-1\u03362\u0336-0\u03361\u0336", // Overstrike characters
		Category: "invalid",
		Scenario: "کاراکترهای اوراسترایک یونیکد در هر رقم",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       81,
		Content:  "⓪②⓪②④-①②-⓪①", // Circled digits
		Category: "invalid",
		Scenario: "استفاده از اعداد دایره‌ای یونیکد",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       82,
		Content:  "２０２４１２０１", // Full-width compact format
		Category: "invalid",
		Scenario: "فرمت فشرده با اعداد تمام عرض",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       83,
		Content:  "2024-12-01Z", // Timezone suffix
		Category: "invalid",
		Scenario: "پسوند منطقه زمانی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       84,
		Content:  "᠒᠐᠒᠔-᠑᠒-᠐᠑", // Mongolian numerals
		Category: "invalid",
		Scenario: "اعداد مغولی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       85,
		Content:  "2024-\u001B12-01", // Escape sequence
		Category: "invalid",
		Scenario: "دنباله فرار در میان رشته",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       86,
		Content:  "𝟸𝟶𝟸𝟺-𝟷𝟸-𝟶𝟷", // Mathematical monospace digits
		Category: "invalid",
		Scenario: "اعداد تک‌فاصله ریاضی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       87,
		Content:  "2024-12-01T", // Incomplete datetime
		Category: "invalid",
		Scenario: "فرمت تاریخ-زمان ناتمام",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       88,
		Content:  "௨௦௨௪-௧௨-௦௧", // Tamil numbers
		Category: "invalid",
		Scenario: "اعداد تامیلی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       89,
		Content:  "2024-12-001", // Triple zero padding
		Category: "invalid",
		Scenario: "صفر پیشرو سه رقمی در روز",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       90,
		Content:  "2024-12-01 ", // Vertical tab suffix
		Category: "invalid",
		Scenario: "کاراکتر تب عمودی در انتها",
		Provider: "DeepSeek",
		Expected: responseCode200("2024-12-01", nil),
	},
	{
		Id:       91,
		Content:  "𓎆𓎆𓏻𓏽-𓎆𓏻-𓎆", // Egyptian hieroglyph numbers
		Category: "invalid",
		Scenario: "اعداد هیروگلیف مصری",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       92,
		Content:  "2٠24-1۲-0۱", // Mixed ASCII/Arabic-Indic
		Category: "invalid",
		Scenario: "ترکیب اعداد ASCII و عربی-هندی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       93,
		Content:  "2024-12-01 ", // Form feed character
		Category: "invalid",
		Scenario: "کاراکتر پیش‌برگ صفحه",
		Provider: "DeepSeek",
		Expected: responseCode200("2024-12-01", nil),
	},
	{
		Id:       94,
		Content:  "൨൦൨൪-൧൨-൦൧", // Malayalam numerals
		Category: "invalid",
		Scenario: "اعداد مالایالم",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       95,
		Content:  "2024\uFE4A12\uFE4A01", // Wavy hyphen
		Category: "invalid",
		Scenario: "جداکننده خط موجدار",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       96,
		Content:  "᪂0᪂᪄-᪂᪂-0᪂", // Tai Tham numerals
		Category: "invalid",
		Scenario: "اعداد تای تام",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       97,
		Content:  "2024-12-01\\", // Backslash suffix
		Category: "invalid",
		Scenario: "بک اسلش در انتهای رشته",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       98,
		Content:  "ⅿⅿⅿⅿ-ⅻ-ⅰ", // Roman numeral lowercase
		Category: "invalid",
		Scenario: "اعداد رومی کوچک",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       99,
		Content:  "2024-12-01%0A", // URL encoded newline
		Category: "invalid",
		Scenario: "کاراکتر جدید خط کدگذاری شده",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       100,
		Content:  "𝕿𝖍𝖎𝖘-𝖎𝖘-𝖋𝖆𝖐𝖊", // Mathematical bold fraktur
		Category: "invalid",
		Scenario: "فونت ریاضی فاکتوریل پررنگ",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
}
