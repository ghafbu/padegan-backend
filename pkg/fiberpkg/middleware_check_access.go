package fiberpkg

import (
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
	"github.com/gofiber/fiber/v3"
)

func MiddlewareCheckAccess() func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		var (
			thisMessage []messagepkg.Response
			thisLang    = headerpkg.GetIndexAsString(c.GetReqHeaders()["S-Language"])
			thisId      = fiber.Params[uint64](c, "id")
		)

		// Token check is not nil
		// In middleware_auth: ghafpkg.UserIdContextKey always uint64, therefore not err check
		if c.Locals(ContextKeyUserId) == nil || c.Locals(ContextKeyUserId).(uint64) != thisId {
			thisMessage = append(thisMessage, messagepkg.NoAccess(&thisLang, messagepkg.Keyword.Field["no_access"]))

			return ReturnAuto(c,
				responsepkg.Code403(
					thisMessage,
					nil,
				))
		}

		return c.Next()
	}
}
