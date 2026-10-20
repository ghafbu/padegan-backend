package defaulttestpkgdtomobile

var variablesInvalidDeepSeek = []myStruct{
	{
		Id:       1,
		Content:  "+989112253390",
		Category: "invalid",
		Scenario: "شماره معتبر با فرمت E.164",
		Provider: "DeepSeek",
		Expected: responseCode200("", nil),
	},
	{
		Id:       2,
		Content:  "+۹۸۹۱۱۲۲۳۳۴۴",
		Category: "invalid",
		Scenario: "کاراکترهای فارسی در شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       3,
		Content:  "+१२०४३२१०९८७", // هندی
		Category: "invalid",
		Scenario: "کاراکترهای هندی در شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       4,
		Content:  "+98912ABC5678",
		Category: "invalid",
		Scenario: "حروف لاتین در شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       5,
		Content:  "00989122334455",
		Category: "invalid",
		Scenario: "شروع با 00 به جای +",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       6,
		Content:  "+981123344556677",
		Category: "invalid",
		Scenario: "طول بیش از 15 کاراکتر",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       7,
		Content:  "+",
		Category: "invalid",
		Scenario: "فقط کاراکتر +",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       8,
		Content:  "  +989121234567  ",
		Category: "invalid",
		Scenario: "فاصله اضافی در ابتدا و انتها",
		Provider: "DeepSeek",
		Expected: responseCode200("", nil),
	},
	{
		Id:       9,
		Content:  "+۹۸۹۱۲۳۴۵۶۷۸۹۰", // فارسی با طول 15
		Category: "invalid",
		Scenario: "کاراکترهای فارسی با طول معتبر",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       10,
		Content:  "+@#$%^&*()",
		Category: "invalid",
		Scenario: "کاراکترهای خاص",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       11,
		Content:  "",
		Category: "invalid",
		Scenario: "مقدار خالی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       12,
		Content:  "+98912 345 6789",
		Category: "invalid",
		Scenario: "فاصله در میان شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       13,
		Content:  "+982188776655", // تهران
		Category: "invalid",
		Scenario: "شماره ثابت به جای موبایل",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       14,
		Content:  "+16135550123", // آمریکا
		Category: "invalid",
		Scenario: "شماره بین المللی غیر ایرانی - نوع شماره FIXED_LINE_OR_MOBILE است",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       15,
		Content:  "+۹۸۹۱۲۳۴۵۶۷۸۹", // فارسی با طول 14
		Category: "invalid",
		Scenario: "طول غیر استاندارد با کاراکتر فارسی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       16,
		Content:  "+98903۱۴۵۶۷۸۹", // ترکیب فارسی و انگلیسی
		Category: "invalid",
		Scenario: "ترکیب کاراکترهای فارسی و انگلیسی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       17,
		Content:  "+۹۸۹۱۲۳۴۵۶۷۸۹۰۱۲۳۴۵", // طول زیاد
		Category: "invalid",
		Scenario: "طول بسیار زیاد با کاراکتر فارسی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       18,
		Content:  "+123", // طول بسیار کم
		Category: "invalid",
		Scenario: "طول بسیار کم",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       19,
		Content:  "+۹۸۹۱۲۳۴۵۶۷۸۹", // فارسی با طول 13
		Category: "invalid",
		Scenario: "طول ناکافی با کاراکتر فارسی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       20,
		Content:  "+98912-345-6789",
		Category: "invalid",
		Scenario: "استفاده از خط تیره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       21,
		Content:  "+447700900①②③", // کاراکترهای دایرهای انگلیسی
		Category: "invalid",
		Scenario: "استفاده از نمادهای دایرهای به جای اعداد",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       22,
		Content:  "+๙๘๙๑๒๓๔๕๖๗๘", // تایلندی
		Category: "invalid",
		Scenario: "استفاده از اعداد تایلندی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       23,
		Content:  "+۹۸۹۱۲۳۴۵۶۷۸۹۰۱۲۳۴۵۶۷", // 20 کاراکتر
		Category: "invalid",
		Scenario: "طول بسیار طولانی با کاراکتر فارسی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       24,
		Content:  "+۱٢۳۴۵٦۷۸۹", // ترکیب اعداد فارسی/عربی/انگلیسی
		Category: "invalid",
		Scenario: "ترکیب چندین سیستم عددی مختلف",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       25,
		Content:  "+𝟴𝟵𝟭𝟮𝟯𝟰𝟱𝟲𝟳𝟴𝟵", // اعداد ریاضی bold
		Category: "invalid",
		Scenario: "استفاده از کاراکترهای یونیکد خاص",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       26,
		Content:  "+۹۸_۹۱۲_۳۴۵۶۷۸", // underline جداکننده
		Category: "invalid",
		Scenario: "استفاده از آندراسکو به عنوان جداکننده",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       27,
		Content:  "+۹۸.۹۱۲.۳۴۵.۶۷۸", // نقطه جداکننده
		Category: "invalid",
		Scenario: "استفاده از نقطه به عنوان جداکننده",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       28,
		Content:  "+۰۰۹۸۹۱۲۳۴۵۶۷۸۹", // چند صفر قبل از کد کشور
		Category: "invalid",
		Scenario: "صفرهای اضافی قبل از کد کشور",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       29,
		Content:  "＋９８۹１２۳۴５۶۷۸۹", // کاراکتر + یونیکد کامل
		Category: "invalid",
		Scenario: "استفاده از کاراکتر + یونیکد گسترده",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       30,
		Content:  "+۹۸۹۱۲۳۴۵۶۷۸۹۰a", // حرف کوچک در انتها
		Category: "invalid",
		Scenario: "ترکیب حروف و اعداد",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       31,
		Content:  "+۸۶۱۳۵۵۵۰۱۲۳۴۵", // چین با پیش‌شماره غیرمعمول
		Category: "invalid",
		Scenario: "کد کشور معتبر با ساختار غیرمعمول",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       32,
		Content:  "+800123456789", // شماره رایگان بین‌المللی
		Category: "invalid",
		Scenario: "کد سرویس جهانی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       33,
		Content:  "+888123456789", // کد کشور اختصاص نیافته
		Category: "invalid",
		Scenario: "کد کشور نامعتبر",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       34,
		Content:  "+0000000000000", // تمام صفر
		Category: "invalid",
		Scenario: "مقدار تمام صفر",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       35,
		Content:  "+۹۸۹۱۲۳۴۵۶۷۸۹ ", // فاصله در انتها بدون trim
		Category: "invalid",
		Scenario: "فاصله انتهایی بدون عملیات trim",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       36,
		Content:  "+\u06F9\u06F8\u06F9\u06F1\u06F2\u06F3\u06F4\u06F5\u06F6\u06F7\u06F8\u06F9", // اعداد فارسی یونیکد
		Category: "invalid",
		Scenario: "نمایش یونیکد اعداد فارسی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       37,
		Content:  "+𝟭𝟮𝟯𝟰𝟱𝟲𝟳𝟴𝟵𝟬", // اعداد ریاضی bold انگلیسی
		Category: "invalid",
		Scenario: "استایل ریاضی برای اعداد",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       38,
		Content:  "+⁹⁸⁹¹²³⁴⁵⁶⁷⁸⁹", // اعداد بالانویس
		Category: "invalid",
		Scenario: "استفاده از اعداد بالانویس",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       39,
		Content:  "+₉₈₉₁₂₃₄₅₆₇₈₉", // اعداد زیرنویس
		Category: "invalid",
		Scenario: "استفاده از اعداد زیرنویس",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       40,
		Content:  "+९८९१२३४५६७८९", // ترکیب دوانگاریسی و هندی
		Category: "invalid",
		Scenario: "ترکیب اعداد هندی و انگلیسی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       41,
		Content:  "+⓪⑨⑧⑨①②③④⑤⑥⑦⑧", // اعداد دایره‌ای توخالی
		Category: "invalid",
		Scenario: "استفاده از اعداد دایره‌ای توخالی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       42,
		Content:  "+𝟧𝟢𝟣𝟤𝟥𝟦𝟧𝟨𝟩𝟪𝟫", // اعداد ریاضی sans-serif
		Category: "invalid",
		Scenario: "استایل عددی sans-serif",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       43,
		Content:  "+９８۹１٢٣٤５６７８", // ترکیب اعداد تمام‌عرض
		Category: "invalid",
		Scenario: "ترکیب اعداد تمام‌عرض و نیم‌عرض",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       44,
		Content:  "+🖩۹۸۹۱۲۳۴۵۶۷۸", // ترکیب نماد ماشین حساب
		Category: "invalid",
		Scenario: "ترکیب نمادهای گرافیکی با اعداد",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       45,
		Content:  "+⁶⁵⁴⁷⁸⁹¹²³⁴⁵", // اعداد بالانویس معکوس
		Category: "invalid",
		Scenario: "نمایش اعداد با بالانویس معکوس",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       46,
		Content:  "+９８۹１２۳۴５۶۷۸۹۰۱۲۳۴۵۶۷", // 25 کاراکتر
		Category: "invalid",
		Scenario: "طول فوق العاده زیاد (25 کاراکتر)",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       47,
		Content:  "+ǀǀǀǀǀǀǀǀǀǀǀ", // کاراکتر خط عمودی (|)
		Category: "invalid",
		Scenario: "استفاده از کاراکترهای خط عمودی تکراری",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       48,
		Content:  "+۹۸۹۱۲۳۴۵۶۷۸۹۰۱۲۳۴۵", // 18 کاراکتر فارسی
		Category: "invalid",
		Scenario: "طول نامتعارف با کاراکترهای فارسی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       49,
		Content:  "+⒈⒉⒊⒋⒌⒍⒎⒏⒐", // اعداد نقطه‌دار
		Category: "invalid",
		Scenario: "استفاده از اعداد نقطه‌دار",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       50,
		Content:  "+②⓪②③④⑤⑥⑦⑧⑨", // ترکیب اعداد دایره‌ای
		Category: "invalid",
		Scenario: "ترکیب اعداد دایره‌ای مختلف",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       51,
		Content:  "+۹۸۹۱۲۳۴۵۶۷۸۹", // ترکیب فارسی و انگلیسی (۹ انگلیسی)
		Category: "invalid",
		Scenario: "ترکیب ناهمگون اعداد فارسی و انگلیسی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       52,
		Content:  "+🇮🇷989123456789", // افزودن پرچم
		Category: "invalid",
		Scenario: "استفاده از ایموجی پرچم",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       53,
		Content:  "+९८९१२३४५६७८९", // اعداد دیواناگری (هندی)
		Category: "invalid",
		Scenario: "استفاده از اعداد دیواناگری خالص",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       54,
		Content:  "+９８۹１２۳۴５６７۸۹۰۱۲", // 16 کاراکتر تمام‌عرض
		Category: "invalid",
		Scenario: "طول نامتعارف با اعداد تمام‌عرض",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       55,
		Content:  "+₉₈₉₁₂₃₄₅₆₇₈₉₀", // اعداد زیرنویس طولانی
		Category: "invalid",
		Scenario: "اعداد زیرنویس با طول معتبر",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       56,
		Content:  "+９８۹１２۳۴５۶７۸۹۰۱۲۳۴۵", // 20 کاراکتر تمام‌عرض
		Category: "invalid",
		Scenario: "حداکثر طول غیرمجاز تمام‌عرض",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       57,
		Content:  "+⓿❶❷❸❹❺❻❼❽❾", // اعداد دایره‌ای توخالی با پس‌زمینه
		Category: "invalid",
		Scenario: "اعداد دایره‌ای توخالی با نمادهای خاص",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       58,
		Content:  "+𝟗𝟠𝟡𝟙𝟚𝟛𝟜𝟝𝟞𝟟𝟠𝟡", // ترکیب استایل‌های مختلف ریاضی
		Category: "invalid",
		Scenario: "ترکیب فونت‌های ریاضی ناهمگون",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       59,
		Content:  "+۹۸۹۱۲۳۴۵۶۷۸۹۰۱۲", // 15 کاراکتر فارسی
		Category: "invalid",
		Scenario: "حداکثر طول مجاز با اعداد فارسی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       60,
		Content:  "+ⅠⅡⅢⅣⅤⅥⅦⅧⅨⅩ", // اعداد رومی
		Category: "invalid",
		Scenario: "استفاده از اعداد رومی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
}
