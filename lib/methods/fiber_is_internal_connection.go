package methods

import (
	"github.com/gofiber/fiber/v3"
)

func FiberIsInternalConnection(paramCtx fiber.Ctx) bool {
	var thisStatus, ok = paramCtx.Locals("x-Internal-Connection-Status").(bool)
	if ok && thisStatus {
		return true
	}

	return false
}

// func FiberCheckHeaderServiceInternalConnectionIsTrue(paramCtx fiber.Ctx) bool {
//	// بازیابی مقدار از Locals
//	value := paramCtx.Locals("x-Internal-Connection-Status")
//
//	// بررسی مقدار و نوع
//	if status, ok := value.(bool); ok {
//		fmt.Println("x-Internal-Connection-Status-------:", status)
//		return status
//	}
//
//	return false
// }
