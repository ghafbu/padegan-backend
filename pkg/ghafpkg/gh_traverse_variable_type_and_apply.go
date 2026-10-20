package ghafpkg

import "reflect"

// اینجا قراره توابع اماده تولید بکنیم که در پیمایش توابع نا معلوم بهمون کمک بکنه
// مثلا در request ارسالی از سمت کاربر شاید *string,string,map,*map و.... برای متغییر تعیین بشه
// ما میخوایم یک الگو بسازیم که هر نوع داده ای ارسال شد، در مقابل تابع درخواستی روش اعمال یا apply بشه

func GHTraverseVariableTypeAndApplyStringFuncM1(paramField reflect.Value, paramFunc func(string) bool) bool {
	// اگر nil بود، معتبر حساب می‌کنیم
	if !paramField.IsValid() || (paramField.Kind() == reflect.Ptr && paramField.IsNil()) {
		return true
	}

	switch paramField.Kind() {
	case reflect.Ptr:
		return GHTraverseVariableTypeAndApplyStringFuncM1(paramField.Elem(), paramFunc)
	case reflect.String:
		return paramFunc(paramField.String())
	case reflect.Map:
		for _, key := range paramField.MapKeys() {
			val := paramField.MapIndex(key)
			if !GHTraverseVariableTypeAndApplyStringFuncM1(val, paramFunc) {
				return false
			}
		}
		return true

	case reflect.Slice, reflect.Array:
		for i := 0; i < paramField.Len(); i++ {
			if !GHTraverseVariableTypeAndApplyStringFuncM1(paramField.Index(i), paramFunc) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

// یک متغییر به این تابع میدی و یک قانون هم برایش میگزاری به شکل رشته و سپس متغییر و قانون برای تابع ای که در پارامتر سوم معرفی کردی ارسال میشه
// مثلا فرض کن میخوای طول رشته را بررسی بکنی که ایا طول 1-3 را رعایت میکنن یا خیر
// در پارامتر اول متغییر رو میدیم ، دقت بکنید که ما از نوع داده ای متغییر هیچ خبری نداریم در این سناریو
// در پارامتر دوم قانون را به شکل رشته تعیین میکنیم، مثلا 1-3 یعنی حداقل 1 و حداکثر 3 کارکتر
// در پارامتر سوم تابع GHIsLength رو میدیم ، یعنی قراره اون 2 پارامتر برای GHIsLength ارسال بشه و نتیجه اعلام بشه
// این یک مثال بود و شما میتوانید در همه توابع دیگه که از این الگو استفاده میکنن برای پایش متغییر نا معلوم استفاده بکنید

func GHTraverseVariableTypeAndApplyStringFuncM2(paramField reflect.Value, ParamRole string, paramFunc func(string, string) bool) bool {
	// اگر nil بود، معتبر حساب می‌کنیم
	if !paramField.IsValid() || (paramField.Kind() == reflect.Ptr && paramField.IsNil()) {
		return true
	}

	switch paramField.Kind() {
	case reflect.Ptr:
		return GHTraverseVariableTypeAndApplyStringFuncM2(paramField.Elem(), ParamRole, paramFunc)
	case reflect.String:
		return paramFunc(paramField.String(), ParamRole)
	case reflect.Map:
		for _, key := range paramField.MapKeys() {
			val := paramField.MapIndex(key)
			if !GHTraverseVariableTypeAndApplyStringFuncM2(val, ParamRole, paramFunc) {
				return false
			}
		}
		return true

	case reflect.Slice, reflect.Array:
		for i := 0; i < paramField.Len(); i++ {
			if !GHTraverseVariableTypeAndApplyStringFuncM2(paramField.Index(i), ParamRole, paramFunc) {
				return false
			}
		}
		return true
	default:
		return false
	}
}
