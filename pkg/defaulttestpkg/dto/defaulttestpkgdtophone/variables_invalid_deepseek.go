package defaulttestpkgdtophone

var variablesInvalidDeepSeek = []myStruct{
	{
		Id:       1,
		Content:  "+982144445678",
		Category: "invalid",
		Scenario: "شماره تلفن ثابت معتبر تهران با پیش‌شماره 21",
		Provider: "DeepSeek",
		Expected: responseCode200("", nil),
	},
	{
		Id:       2,
		Content:  "+۹۸۲۱۴۴۴۴۵۶۷۸", // اعداد فارسی
		Category: "invalid",
		Scenario: "استفاده از اعداد فارسی در کل شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       3,
		Content:  "+989121234567", // شماره موبایل
		Category: "invalid",
		Scenario: "شماره موبایل ایران با پیش‌شماره 912",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       4,
		Content:  "+98-21-4445678", // خط تیره
		Category: "invalid",
		Scenario: "استفاده از کاراکترهای خاص در شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       5,
		Content:  "+982944445678", // کد منطقه نامعتبر
		Category: "invalid",
		Scenario: "پیش‌شماره نامعتبر 29 برای ایران",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       6,
		Content:  "+12125552345", // شماره آمریکا
		Category: "invalid",
		Scenario: "شماره ثابت معتبر از کشور دیگر (آمریکا)",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       7,
		Content:  "+98214", // طول کوتاه
		Category: "invalid",
		Scenario: "شماره با طول نادرست (بسیار کوتاه)",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       8,
		Content:  "+۹۸21۴۴۴۵6۷۸", // ترکیب اعداد هندی/فارسی
		Category: "invalid",
		Scenario: "ترکیب اعداد فارسی/هندی با لاتین",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       9,
		Content:  "+98a214445678", // حروف لاتین
		Category: "invalid",
		Scenario: "وجود حروف لاتین در شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       10,
		Content:  "+983114445678", // اصفهان
		Category: "invalid",
		Scenario: "پیش‌شماره معتبر اصفهان (311)",
		Provider: "DeepSeek",
		Expected: responseCode200("", nil),
	},
	{
		Id:       11,
		Content:  "982144445678", // بدون +
		Category: "invalid",
		Scenario: "شماره بدون علامت + در ابتدا",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       12,
		Content:  "+9802144445678", // صفر اضافه
		Category: "invalid",
		Scenario: "وجود صفر بعد از کد کشور",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       13,
		Content:  "+987144445678", // شیراز
		Category: "invalid",
		Scenario: "پیش‌شماره معتبر شیراز (71)",
		Provider: "DeepSeek",
		Expected: responseCode200("", nil),
	},
	{
		Id:       14,
		Content:  "+98 21 444 5678", // فاصله
		Category: "invalid",
		Scenario: "استفاده از فاصله در شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       15,
		Content:  "+982594445678", // کد منطقه ناموجود
		Category: "invalid",
		Scenario: "کد منطقه Qom در ایران سالم",
		Provider: "DeepSeek",
		Expected: responseCode200("", nil),
	},
	{
		Id:       16,
		Content:  "+۹۸2۱۴4۴۵۶۷۸", // ترکیب فارسی و لاتین
		Category: "invalid",
		Scenario: "ترکیب اعداد فارسی و لاتین در بخش‌های مختلف",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       17,
		Content:  "+985144445678", // مشهد
		Category: "invalid",
		Scenario: "پیش‌شماره معتبر مشهد (51)",
		Provider: "DeepSeek",
		Expected: responseCode200("", nil),
	},
	{
		Id:       18,
		Content:  "+98(21)4445678", // پرانتز
		Category: "invalid",
		Scenario: "استفاده از پرانتز در شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       19,
		Content:  "+9821444456789011", // طولانی
		Category: "invalid",
		Scenario: "شماره با طول بسیار زیاد",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       20,
		Content:  "+", // فقط +
		Category: "invalid",
		Scenario: "فقط علامت + بدون اعداد",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       21,
		Content:  "02144445678", // فرمت بدون + و کد کشور
		Category: "invalid",
		Scenario: "استفاده از شماره بدون کد کشور و +",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       22,
		Content:  "+98214444567", // طول کوتاه (10 رقم بعد از +98)
		Category: "invalid",
		Scenario: "طول نادرست شماره ثابت (کمتر از حد مجاز)",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       23,
		Content:  "+۹۸۲۱٤٤٤۵۶۷۸", // ترکیب اعداد فارسی/عربی/هندی
		Category: "invalid",
		Scenario: "ترکیب چندین نوع رقم غیرلاتین در شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       24,
		Content:  "+98 21 4444 5678", // فاصله در فرمت
		Category: "invalid",
		Scenario: "استفاده از فاصله در میان ارقام",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       25,
		Content:  "+9821000000000", // طول بلند با اعداد تکراری
		Category: "invalid",
		Scenario: "شماره با ارقام تکراری و طول غیرمعمول",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       26,
		Content:  "+9821۴۴۴۵۶۷۸", // ترکیب اعداد فارسی در میانه شماره
		Category: "invalid",
		Scenario: "اعداد فارسی در بخش محلی شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       27,
		Content:  "+9821-444-5678", // خط تیره در شماره
		Category: "invalid",
		Scenario: "استفاده از جداکننده در شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       28,
		Content:  "+98000000000", // پیش شماره نامعتبر با صفرهای متوالی
		Category: "invalid",
		Scenario: "پیش‌شماره تمام صفر",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       29,
		Content:  "+۹۸۲۱۱۲۳۴۵۶۷", // اعداد فارسی با پیش‌شماره تهران
		Category: "invalid",
		Scenario: "کل شماره با اعداد فارسی نوشته شده",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       30,
		Content:  "+982191234567", // پیش‌شماره موبایل (912 در پیش‌شماره ثابت)
		Category: "invalid",
		Scenario: "شماره ثابت درست از استان تهران",
		Provider: "DeepSeek",
		Expected: responseCode200("", nil),
	},
	{
		Id:       31,
		Content:  "+9821444456789", // طول اضافی
		Category: "invalid",
		Scenario: "شماره با ارقام اضافی در انتها",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       32,
		Content:  "+۹८۲۱44۴۵67۸", // ترکیب اعداد در زبان‌های مختلف
		Category: "invalid",
		Scenario: "ترکیب چندین سیستم عددی در یک شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       33,
		Content:  "+9821PhoneNumber", // حروف در شماره
		Category: "invalid",
		Scenario: "استفاده از کلمات به جای ارقام",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       34,
		Content:  "+98214444", // شماره ناقص
		Category: "invalid",
		Scenario: "شماره ناتمام بدون ارقام کافی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       35,
		Content:  "+9821 444 45678", // فاصله نامتعارف
		Category: "invalid",
		Scenario: "فاصله گذاری غیراستاندارد در شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       36,
		Content:  "+9821/444/5678", // اسلش در شماره
		Category: "invalid",
		Scenario: "استفاده از کاراکتر اسلش در شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       37,
		Content:  "+۹۸۲۱44۴۵۶۷۸", // اعداد فارسی و لاتین مخلوط
		Category: "invalid",
		Scenario: "ترکیب اعداد فارسی و لاتین در کل شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       38,
		Content:  "+98214444567A", // حرف در انتهای شماره
		Category: "invalid",
		Scenario: "وجود حرف لاتین در انتهای شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       39,
		Content:  "00982144445678", // فرمت بین‌المللی با صفر اضافی
		Category: "invalid",
		Scenario: "استفاده از صفرهای اضافی قبل از کد کشور",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       40,
		Content:  "+982144445678 ", // فاصله در انتهای شماره
		Category: "invalid",
		Scenario: "فاصله خالی در انتهای شماره",
		Provider: "DeepSeek",
		Expected: responseCode200("+982144445678", nil),
	},
	{
		Id:       41,
		Content:  "+9821555555555", // طول اضافی با اعداد تکراری
		Category: "invalid",
		Scenario: "طول غیرمجاز با الگوی عددی تکراری",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       42,
		Content:  "+9821٠٤٤٤٥٦٧٨", // اعداد عربی-هندی (شرقی)
		Category: "invalid",
		Scenario: "استفاده از ارقام عربی-هندی در پیش‌شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       43,
		Content:  "+9821\u06F4\u06F4\u06F4\u06F4\u06F5\u06F6\u06F7\u06F8", // اعداد فارسی با یونیکد
		Category: "invalid",
		Scenario: "ارقام فارسی با کدهای یونیکد صریح",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       44,
		Content:  "+98 21 4444 5678", // فاصله پس از کد کشور
		Category: "invalid",
		Scenario: "فاصله بعد از کد کشور و بین پیش‌شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       45,
		Content:  "+9821-4444-5678", // الگوی شماره‌گیری محلی
		Category: "invalid",
		Scenario: "فرمت شماره با جداکننده مانند شماره‌های محلی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       46,
		Content:  "+98214444", // پیش‌شماره بدون شماره محلی
		Category: "invalid",
		Scenario: "فقط پیش‌شماره منطقه بدون ارقام بعدی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       47,
		Content:  "+9821٠۴۴۴۵۶۷۸", // ترکیب اعداد فارسی و عربی
		Category: "invalid",
		Scenario: "ترکیب ارقام فارسی (۴) و عربی (٠) در یک شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       48,
		Content:  "+9821ABCDEFGH", // حروف به جای ارقام
		Category: "invalid",
		Scenario: "جایگزینی کامل ارقام با حروف لاتین",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       49,
		Content:  "+9821~!@#$%^", // کاراکترهای ویژه
		Category: "invalid",
		Scenario: "استفاده از کاراکترهای ویژه به جای ارقام",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       50,
		Content:  "+9821۰۹۱۲۳۴۵۶", // پیش‌شماره موبایل در قالب ثابت
		Category: "invalid",
		Scenario: "شماره موبایل استتار شده در فرمت ثابت",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       51,
		Content:  "+9821⁴⁴⁴⁵⁶⁷⁸", // اعداد بالاوند (Superscript)
		Category: "invalid",
		Scenario: "استفاده از نمادهای ریاضی به جای ارقام",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       52,
		Content:  "+۹۸۲۱۴۴۴۵۶۷۸", // اعداد فارسی با فونت متفاوت
		Category: "invalid",
		Scenario: "ارقام فارسی با شکل ظاهری متفاوت (فونت ناآشنا)",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       53,
		Content:  "+98214444۵۶۷۸", // ترکیب لاتین و فارسی در ارقام
		Category: "invalid",
		Scenario: "ارقام لاتین در ابتدا و فارسی در انتها",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       54,
		Content:  "+9821.444.5678", // جداکننده نقطه
		Category: "invalid",
		Scenario: "استفاده از نقطه به عنوان جداکننده",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       55,
		Content:  "+98214444-5678", // خط تیره در موقعیت نادرست
		Category: "invalid",
		Scenario: "خط تیره در بخش شماره محلی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       56,
		Content:  "+9821۴۴۴۵۶۷۸", // اعداد فارسی بعد از پیش‌شماره لاتین
		Category: "invalid",
		Scenario: "پیش‌شماره لاتین با ارقام محلی فارسی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       57,
		Content:  "+9821٤٤٤٥٦٧٨", // ترکیب اعداد عربی و هندی
		Category: "invalid",
		Scenario: "ترکیب ارقام عربی (٤) و هندی (٥) در شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       58,
		Content:  "+۹۸02144445678", // صفر اضافه بعد از کد کشور
		Category: "invalid",
		Scenario: "صفر غیرمجاز بعد از کد کشور ایران",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       59,
		Content:  "+9821۴۴۴۵۶۷۸۹۰۱۲", // طول زیاد با ارقام فارسی
		Category: "invalid",
		Scenario: "طول بیش از حد با ارقام غیرلاتین",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       60,
		Content:  "+9821NULL", // کلمه کلیدی برنامه‌نویسی
		Category: "invalid",
		Scenario: "استفاده از کلمه کلیدی NULL در شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       61,
		Content:  "+9821' OR 1=1;--", // حمله تزریق SQL
		Category: "invalid",
		Scenario: "رشته حاوی دستورات SQL تزریقی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       62,
		Content:  "+98\u202E2144445678", // کاراکتر RTL (راست به چپ)
		Category: "invalid",
		Scenario: "وجود کاراکتر کنترل جهت متن در شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       63,
		Content:  "+98๒๑44445678", // ارقام تایلندی (๒=2, ๑=1)
		Category: "invalid",
		Scenario: "استفاده از ارقام تایلندی در پیش‌شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       64,
		Content:  "+9821🎉444🎉5678", // ایموجی در شماره
		Category: "invalid",
		Scenario: "وجود ایموجی در میان ارقام",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       65,
		Content:  "+9821999999999999999", // مقدار عددی بسیار بزرگ
		Category: "invalid",
		Scenario: "شماره با مقدار عددی خارج از محدوده int64",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       66,
		Content:  "+9821\x00\x00\x00", // کاراکترهای NULL
		Category: "invalid",
		Scenario: "کاراکترهای NULL جاسازی شده در رشته",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       67,
		Content:  "+9821O4445678", // حرف O به جای صفر
		Category: "invalid",
		Scenario: "جایگزینی صفر با حرف لاتین O",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       68,
		Content:  "+98214444Z5678", // حرف Z در میانه شماره
		Category: "invalid",
		Scenario: "حرف لاتین در میانه ارقام",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       69,
		Content:  "+9821𝟮𝟭𝟰𝟰𝟰𝟱𝟲𝟳𝟴", // اعداد Mathematical Bold
		Category: "invalid",
		Scenario: "استفاده از ارقام یونیکد Mathematical Bold",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       70,
		Content:  "+9821⁰⁴⁴⁴⁵⁶⁷⁸", // اعداد نمادین (Superscript)
		Category: "invalid",
		Scenario: "ارقام بالاوند با نمادهای ریاضی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       71,
		Content:  "+9821", // کاراکترهای Private Use Area
		Category: "invalid",
		Scenario: "کاراکترهای ویژه یونیکد در ناحیه خصوصی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       72,
		Content:  "+9821۴۴۴۵۶۷۸", // اعداد فارسی با فاصله نامرئی
		Category: "invalid",
		Scenario: "ترکیب اعداد فارسی با فاصله‌های Zero-width",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       73,
		Content:  "+9821👁️🗨️44445678", // ایموجی ترکیبی
		Category: "invalid",
		Scenario: "ایموجی‌های ترکیبی در شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       74,
		Content:  "+9821２１４４４５６７８", // ارقام پهن (Full-width)
		Category: "invalid",
		Scenario: "استفاده از ارقام پهن لاتین",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       75,
		Content:  "+9821𝟒𝟒𝟒𝟓𝟔𝟕𝟖", // اعداد Bold یونیکد
		Category: "invalid",
		Scenario: "ارقام Bold یونیکد در بخش محلی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       76,
		Content:  "+9821𑁦𑁧𑁨𑁩𑁪𑁫", // ارقام برمه‌ای
		Category: "invalid",
		Scenario: "استفاده از سیستم عددی برمه‌ای",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       77,
		Content:  "+9821١٢٣٤٥٦٧٨", // ارقام عربی-هندی در بخش محلی
		Category: "invalid",
		Scenario: "ارقام عربی در بخش شماره محلی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       78,
		Content:  "+9821⓪②⓸④⑤⑥⑦⑧", // اعداد محصور شده
		Category: "invalid",
		Scenario: "استفاده از اعداد محصور شده (Enclosed Alphanumerics)",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       79,
		Content:  "+9821𝟭𝟮𝟯𝟺𝟻𝟼𝟽𝟾", // اعداد Sans-Serif Bold
		Category: "invalid",
		Scenario: "ارقام با استایل فونت متفاوت",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       80,
		Content:  "+9821᠒᠑᠔᠔᠕᠖᠗᠘", // ارقام مغولی
		Category: "invalid",
		Scenario: "استفاده از سیستم عددی مغولی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       81,
		Content:  "+9821\uFF12\uFF11\uFF14\uFF14\uFF15\uFF16\uFF17\uFF18", // ارقام پهن لاتین (Full-width)
		Category: "invalid",
		Scenario: "استفاده از ارقام لاتین پهن (Full-width) در کل شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       82,
		Content:  "+9821↨444↨5678", // کاراکترهای کنترل قدیمی (VT)
		Category: "invalid",
		Scenario: "کاراکترهای کنترل Vertical Tab در میانه شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       83,
		Content:  "+9821𝟷𝟸𝟹𝟺𝟻𝟼𝟽𝟾", // اعداد Monospace یونیکد
		Category: "invalid",
		Scenario: "ارقام با استایل Monospace یونیکد",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       84,
		Content:  "+9821൪൫൬൭൮൯൰൱", // ارقام مالایالم
		Category: "invalid",
		Scenario: "استفاده از سیستم عددی مالایالم",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       85,
		Content:  "+9821\u202A2144445678\u202C", // کاراکترهای جهت‌دهی متن
		Category: "invalid",
		Scenario: "جاسازی کاراکترهای جهت‌دهی متن در شماره",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       86,
		Content:  "+9821𐌎𐌏𐌐𐌑𐌒𐌓𐌔𐌕", // ارقام اتروسکی
		Category: "invalid",
		Scenario: "استفاده از ارقام سیستم اتروسکی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       87,
		Content:  "+9821⃢4⃢4⃢4⃢5⃢6⃢7⃢8", // اعداد با نشانه ترکیبی
		Category: "invalid",
		Scenario: "ارقام با نشانگرهای ترکیبی یونیکد",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       88,
		Content:  "+9821🄀🄁🄂🄃🄄🄅🄆🄇", // اعداد محصور شده با حاشیه
		Category: "invalid",
		Scenario: "ارقام محصور شده با حاشیه (Enclosed Alphanumeric Supplement)",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       89,
		Content:  "+9821𝟜𝟜𝟜𝟝𝟞𝟟𝟠", // اعداد دوخطی (Double-struck)
		Category: "invalid",
		Scenario: "ارقام دوخطی ریاضی یونیکد",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       90,
		Content:  "+9821④④④⑤⑥⑦⑧", // اعداد محصور شده سفید
		Category: "invalid",
		Scenario: "ارقام محصور شده سفید (White Circle Numbers)",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       91,
		Content:  "+9821🅂🄴🄲🅄🅁🄸🅃🅈", // حروف جایگزین با استایل خاص
		Category: "invalid",
		Scenario: "استفاده از حروف استایل‌دار به جای ارقام",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       92,
		Content:  "+9821⁰¹²³⁴⁵⁶⁷⁸", // اعداد نمادین ریاضی
		Category: "invalid",
		Scenario: "ترکیب نمادهای ریاضی به عنوان ارقام",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       93,
		Content:  "+9821⓿⓫⓬⓭⓮⓯⓰⓱", // اعداد محصور شده غیراستاندارد
		Category: "invalid",
		Scenario: "اعداد محصور شده با نمادهای غیرمعمول",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       94,
		Content:  "+9821᠐᠑᠒᠓᠔᠕᠖᠗", // ارقام مغولی قدیمی
		Category: "invalid",
		Scenario: "استفاده از ارقام مغولی سنتی",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       95,
		Content:  "+9821𝟦𝟦𝟦𝟧𝟨𝟩𝟪", // اعداد Sans-Serif
		Category: "invalid",
		Scenario: "ارقام با فونت Sans-Serif یونیکد",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       96,
		Content:  "+9821₀₁₂₃₄₅₆₇₈", // اعداد زیرنویسی
		Category: "invalid",
		Scenario: "استفاده از ارقام زیرنویسی (Subscript)",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       97,
		Content:  "+9821⑪⑫⑬⑭⑮⑯⑰⑱", // اعداد محصور شده دو رقمی
		Category: "invalid",
		Scenario: "اعداد محصور شده دو رقمی در قالب تک کاراکتری",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       98,
		Content:  "+9821🄫🄬🄭🄮🄯🄰🄱🄲", // نمادهای ترکیبی یونیکد
		Category: "invalid",
		Scenario: "استفاده از نمادهای ترکیبی خاص یونیکد",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       99,
		Content:  "+9821𐏑𐏒𐏓𐏔𐏕𐏖𐏗𐏘", // ارخط فارسی باستان
		Category: "invalid",
		Scenario: "استفاده از سیستم عددی فارسی باستان",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
	{
		Id:       100,
		Content:  "+9821𝕏𝕀𝕀𝕀𝕍𝕍𝕀𝕀𝕀", // اعداد با استایل Blackboard Bold
		Category: "invalid",
		Scenario: "استفاده از نمادهای ریاضی Blackboard Bold",
		Provider: "DeepSeek",
		Expected: responseCode400("", nil),
	},
}
