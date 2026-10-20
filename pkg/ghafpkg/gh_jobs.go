package ghafpkg

type JobInfo struct {
	Farsi   string
	English string
}

func ConvertToJobInfo(jobID *uint) JobInfo {
	if jobID != nil {
		return findJob(*jobID)
	}

	return getUnknownJob()
}

var jobs = map[uint]JobInfo{
	0:   {Farsi: "نامشخص", English: "Unknown"},
	1:   {Farsi: "مدیرعامل", English: "CEO"},
	2:   {Farsi: "مدیر داخلی", English: "Operations Manager"},
	3:   {Farsi: "مدرس", English: "Teacher"},
	4:   {Farsi: "دستیار آموزشی", English: "Teaching Assistant"},
	5:   {Farsi: "دانشجو", English: "Student"},
	6:   {Farsi: "پژوهشگر", English: "Researcher"},
	7:   {Farsi: "برنامه‌نویس", English: "Software Developer"},
	8:   {Farsi: "طراح رابط کاربری", English: "UI Designer"},
	9:   {Farsi: "تحلیل‌گر داده", English: "Data Analyst"},
	10:  {Farsi: "مدیر محصول", English: "Product Manager"},
	11:  {Farsi: "مدیر پروژه", English: "Project Manager"},
	12:  {Farsi: "حسابدار", English: "Accountant"},
	13:  {Farsi: "کارشناس منابع انسانی", English: "HR Specialist"},
	14:  {Farsi: "کارشناس بازاریابی", English: "Marketing Specialist"},
	15:  {Farsi: "پشتیبان فنی", English: "Technical Support"},
	16:  {Farsi: "مدیر فناوری اطلاعات", English: "IT Manager"},
	17:  {Farsi: "طراح گرافیک", English: "Graphic Designer"},
	18:  {Farsi: "توسعه‌دهنده بک‌اند", English: "Backend Developer"},
	19:  {Farsi: "توسعه‌دهنده فرانت‌اند", English: "Frontend Developer"},
	20:  {Farsi: "مدیر فروش", English: "Sales Manager"},
	21:  {Farsi: "کارشناس فروش", English: "Sales Representative"},
	22:  {Farsi: "نویسنده محتوا", English: "Content Writer"},
	23:  {Farsi: "مدیر شبکه", English: "Network Administrator"},
	24:  {Farsi: "کارشناس امنیت اطلاعات", English: "Cybersecurity Specialist"},
	25:  {Farsi: "تحلیل‌گر سیستم", English: "System Analyst"},
	26:  {Farsi: "مهندس نرم‌افزار", English: "Software Engineer"},
	27:  {Farsi: "مهندس سخت‌افزار", English: "Hardware Engineer"},
	28:  {Farsi: "مدیر آموزش", English: "Training Manager"},
	29:  {Farsi: "مشاور آموزشی", English: "Education Consultant"},
	30:  {Farsi: "مربی ورزشی", English: "Fitness Coach"},
	31:  {Farsi: "روان‌شناس", English: "Psychologist"},
	32:  {Farsi: "پزشک", English: "Doctor"},
	33:  {Farsi: "پرستار", English: "Nurse"},
	34:  {Farsi: "داروساز", English: "Pharmacist"},
	35:  {Farsi: "دندان‌پزشک", English: "Dentist"},
	36:  {Farsi: "مهندس عمران", English: "Civil Engineer"},
	37:  {Farsi: "مهندس برق", English: "Electrical Engineer"},
	38:  {Farsi: "مهندس مکانیک", English: "Mechanical Engineer"},
	39:  {Farsi: "معمار", English: "Architect"},
	40:  {Farsi: "راننده", English: "Driver"},
	41:  {Farsi: "آشپز", English: "Chef"},
	42:  {Farsi: "پیک", English: "Courier"},
	43:  {Farsi: "کارگر", English: "Worker"},
	44:  {Farsi: "کارمند اداری", English: "Administrative Staff"},
	45:  {Farsi: "دستیار شخصی", English: "Personal Assistant"},
	46:  {Farsi: "مترجم", English: "Translator"},
	47:  {Farsi: "حقوقدان", English: "Lawyer"},
	48:  {Farsi: "قاضی", English: "Judge"},
	49:  {Farsi: "روزنامه‌نگار", English: "Journalist"},
	50:  {Farsi: "عکاس", English: "Photographer"},
	51:  {Farsi: "فیلم‌بردار", English: "Cinematographer"},
	52:  {Farsi: "کارگردان", English: "Director"},
	53:  {Farsi: "بازیگر", English: "Actor"},
	54:  {Farsi: "خواننده", English: "Singer"},
	55:  {Farsi: "نوازنده", English: "Musician"},
	56:  {Farsi: "هنرمند", English: "Artist"},
	57:  {Farsi: "تاجر", English: "Businessman"},
	58:  {Farsi: "کارآفرین", English: "Entrepreneur"},
	59:  {Farsi: "دانشمند", English: "Scientist"},
	60:  {Farsi: "استاد دانشگاه", English: "Professor"},
	61:  {Farsi: "دانش‌آموز", English: "School Student"},
	62:  {Farsi: "بازنشسته", English: "Retired"},
	63:  {Farsi: "خانه‌دار", English: "Homemaker"},
	64:  {Farsi: "کارگر ساختمانی", English: "Construction Worker"},
	65:  {Farsi: "کشاورز", English: "Farmer"},
	66:  {Farsi: "دامدار", English: "Rancher"},
	67:  {Farsi: "ماهی‌گیر", English: "Fisherman"},
	68:  {Farsi: "راننده تاکسی", English: "Taxi Driver"},
	69:  {Farsi: "راننده کامیون", English: "Truck Driver"},
	70:  {Farsi: "کارشناس لجستیک", English: "Logistics Specialist"},
	71:  {Farsi: "پیک موتوری", English: "Motor Courier"},
	72:  {Farsi: "کارمند بانک", English: "Bank Clerk"},
	73:  {Farsi: "کارشناس بیمه", English: "Insurance Agent"},
	74:  {Farsi: "کارمند پست", English: "Postal Worker"},
	75:  {Farsi: "آتش‌نشان", English: "Firefighter"},
	76:  {Farsi: "پلیس", English: "Police Officer"},
	77:  {Farsi: "سرباز", English: "Soldier"},
	78:  {Farsi: "افسر", English: "Officer"},
	79:  {Farsi: "نظامی بازنشسته", English: "Retired Military"},
	80:  {Farsi: "مدیر محتوا", English: "Content Manager"},
	81:  {Farsi: "متخصص سئو", English: "SEO Specialist"},
	82:  {Farsi: "کارشناس شبکه‌های اجتماعی", English: "Social Media Manager"},
	83:  {Farsi: "تحلیل‌گر مالی", English: "Financial Analyst"},
	84:  {Farsi: "مشاور مالی", English: "Financial Advisor"},
	85:  {Farsi: "معلم", English: "School Teacher"},
	86:  {Farsi: "دبیر", English: "High School Teacher"},
	87:  {Farsi: "مدیر مدرسه", English: "School Principal"},
	88:  {Farsi: "کارشناس کنترل کیفیت", English: "QA Specialist"},
	89:  {Farsi: "کارشناس پشتیبانی", English: "Support Specialist"},
	90:  {Farsi: "برنامه‌نویس موبایل", English: "Mobile Developer"},
	91:  {Farsi: "توسعه‌دهنده گیم", English: "Game Developer"},
	92:  {Farsi: "مهندس هوش مصنوعی", English: "AI Engineer"},
	93:  {Farsi: "تحلیل‌گر داده", English: "Data Scientist"},
	94:  {Farsi: "متخصص یادگیری ماشین", English: "ML Specialist"},
	95:  {Farsi: "مهندس DevOps", English: "DevOps Engineer"},
	96:  {Farsi: "کارشناس Cloud", English: "Cloud Engineer"},
	97:  {Farsi: "مدیر سیستم", English: "System Administrator"},
	98:  {Farsi: "کارشناس پایگاه داده", English: "Database Administrator"},
	99:  {Farsi: "طراح تجربه کاربری", English: "UX Designer"},
	100: {Farsi: "تحلیل‌گر کسب‌وکار", English: "Business Analyst"},
}

func getUnknownJob() JobInfo {
	return jobs[0]
}

func findJob(id uint) JobInfo {
	val, exists := jobs[id]
	if !exists {
		return jobs[0]
	}

	return val
}
