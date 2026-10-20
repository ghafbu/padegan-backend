package methods

import (
	"reflect"

	"github.com/microcosm-cc/bluemonday"
)

func CustomTagSanitizeBluePKG(input interface{}) error {
	if input == nil {
		return nil
	}

	v := reflect.ValueOf(input)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	// از آنجایی که در بالاترین سطح، هنوز تگی نداریم،
	// "سیاست جاری" را nil می‌گذاریم تا فقط فیلدهایی که خودشان تگ دارند پاک‌سازی شوند.
	sanitizeAction(v, nil)

	return nil
}

func sanitizeAction(v reflect.Value, currentPolicy *bluemonday.Policy) {
	// اگر معتبر و قابل Set نیست، برگرد
	if !v.IsValid() || !v.CanSet() {
		return
	}

	switch v.Kind() {
	case reflect.String:
		// اگر سیاست جاری داریم، مقدار را پاک‌سازی کن
		if currentPolicy != nil {
			original := v.String()
			v.SetString(currentPolicy.Sanitize(original))
		}

	case reflect.Struct:
		// پیمایش فیلدهای استراکچر
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			fieldValue := v.Field(i)
			fieldType := t.Field(i)

			// اگر فیلد غیرقابل‌دسترس باشد، ادامه نده
			if !fieldValue.CanSet() {
				continue
			}

			// ببین آیا تگی داریم؟
			tagPolicyName, hasTag := fieldType.Tag.Lookup("sanitize_blue")
			if hasTag {
				// اگر فیلد تگ دارد، یک سیاست جدید بساز (مثلاً high, medium, low)
				newPolicy := getPolicyByName(tagPolicyName)
				// با سیاست جدید، وارد فیلد شو
				sanitizeAction(fieldValue, newPolicy)
			} else {
				// اگر تگ ندارد، با همان سیاست جاری ادامه بده
				sanitizeAction(fieldValue, currentPolicy)
			}
		}

	case reflect.Ptr:
		// اشاره‌گر را دنبال کن
		if v.IsNil() {
			return
		}
		elemVal := v.Elem()
		sanitizeAction(elemVal, currentPolicy)

	case reflect.Slice, reflect.Array:
		// اگر Slice نال نباشد:
		if !v.IsNil() {
			length := v.Len()
			for i := 0; i < length; i++ {
				item := v.Index(i)
				sanitizeAction(item, currentPolicy)
			}
		}

	case reflect.Map:
		if v.IsNil() {
			return
		}
		// برای پاک‌سازی map: باید کلیدها و مقدارها را بررسی کنیم.
		// در این مدل، "تگ" روی فیلد map در استراکچر می‌آید
		// (نه روی تک‌تک کلیدها و مقدارهای map). پس از سیاست جاری استفاده می‌کنیم.
		// کلیدها را (در صورت string بودن) هم می‌توان پاکسازی کرد؛
		// ولی این کار ممکن است ساختار map را دگرگون کند (کلید عوض می‌شود).
		// بسته به نیاز، می‌توان فقط مقدارها را پاکسازی کرد.
		//
		// اینجا برای سادگی، "کلیدها را دست‌نخورده" می‌گذاریم
		// و فقط مقدار را - اگر ساختار یا string است - پاک می‌کنیم.

		// مپ جدید از همان نوع بسازیم
		//keyType := v.Type().Key()
		elemType := v.Type().Elem()
		newMap := reflect.MakeMap(v.Type())

		for _, mapKey := range v.MapKeys() {
			originalVal := v.MapIndex(mapKey)

			// مقدار جدیدی می‌سازیم تا روی آن sanitizeValue صدا بزنیم
			newVal := reflect.New(elemType).Elem()
			newVal.Set(originalVal)
			sanitizeAction(newVal, currentPolicy)

			// همان mapKey را (بدون تغییر) استفاده می‌کنیم
			newMap.SetMapIndex(mapKey, newVal)
		}
		v.Set(newMap)

	default:
		// int, float, bool و... کاری نداریم
	}
}

func getPolicyByName(name string) *bluemonday.Policy {
	switch name {
	case "high":
		// حذف کامل تگ‌های HTML (سیاست سختگیرانه)
		return bluemonday.StrictPolicy()
	case "medium":
		// کمی ملایم‌تر؛ UGCPolicy
		return bluemonday.UGCPolicy()
	case "low":
		// اجازه‌ی برخی تگ‌ها
		p := bluemonday.UGCPolicy()
		// به عنوان مثال، تگ b را هم اجازه می‌دهیم:
		// TODO: تعدادی تگ برای حالت LOW قرار بدید
		p.AllowElements("b", "i", "u")
		return p
	default:
		// اگر مقدار ناشناس بود، یا تگ خالی بود:
		return nil // یعنی عملاً هیچ پاک‌سازی انجام نشود
	}
}
