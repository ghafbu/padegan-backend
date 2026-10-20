package middleware

import (
	"github.com/ghafbu/padegan-backend/lib/consts"
	"github.com/ghafbu/padegan-backend/lib/types"
	"github.com/ghafbu/padegan-backend/pkg/fiberpkg"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
	"github.com/gofiber/fiber/v3"
)

func FiberUserCheckAccess() func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		var (
			thisMessage []messagepkg.Response
			thisLang    = headerpkg.GetIndexAsString(c.GetReqHeaders()["S-Language"])
			thisId      = fiber.Params[uint64](c, "user_id")
		)

		// Token check is not nil
		// In middleware_auth: ghafpkg.UserIdContextKey always uint64, therefore not err check
		if c.Locals(consts.ContextKeyUserId) != nil {
			if c.Locals(consts.ContextKeyUserId).(uint64) != thisId {
				thisMessage = append(thisMessage, messagepkg.NoAccess(&thisLang, messagepkg.Keyword.Field["no_access"]))
				return fiberpkg.ReturnAuto(c,
					responsepkg.Code401(
						thisMessage,
						nil,
					))
			}
		}

		// check my access | EX: ./lib/user access.go
		// check db
		var UserModules = types.UserAccess{
			Role: "admin",
			Modules: types.UserModules{
				Organization: types.UserPermission{
					OnlyMySelf: false,
					Create:     false,
					Delete:     false,
					Edit:       false,
				},
			},
		}

		c.Locals("user-access", UserModules)

		return c.Next()
	}
}
