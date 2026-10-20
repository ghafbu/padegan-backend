package ghafpkg

import (
	"fmt"
	"math/rand"
	"time"
)

func GHRandomDate(paramStartDate time.Time, paramEndDate time.Time) (time.Time, error) {
	if paramStartDate.After(paramEndDate) {
		return time.Time{}, fmt.Errorf("the start date must be before the end date")
	}

	// محاسبه مدت زمان بین دو تاریخ
	duration := paramEndDate.Sub(paramStartDate)

	// تولید یک مدت زمان رندوم بین 0 و duration
	randomDuration := time.Duration(rand.Int63n(duration.Nanoseconds()))

	// اضافه کردن مدت زمان رندوم به تاریخ شروع برای بدست آوردن تاریخ رندوم
	randomTime := paramStartDate.Add(randomDuration)

	// قالب بندی تاریخ رندوم بر اساس فرمت ارائه شده
	return randomTime, nil
}
