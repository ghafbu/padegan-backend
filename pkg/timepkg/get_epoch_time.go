package timepkg

import "time"

func GetEpochTime() int64 {
	// دریافت تایم‌استمپ به صورت نانوثانیه
	now := time.Now()
	timestampMillis := now.UnixNano() / int64(time.Millisecond)

	return timestampMillis
}
