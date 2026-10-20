package fiberpkg

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

func MiddlewareStartExecuteDuration(ctx fiber.Ctx) error {
	ctx.Locals("start_execute_duration", time.Now())

	return ctx.Next()
}

// مدت زمان اجرا هر درخواست را برگشت خواهد داد
func GetExecuteDuration(ctx fiber.Ctx) int64 {
	return time.Since(ctx.Locals("start_execute_duration").(time.Time)).Milliseconds()
}
