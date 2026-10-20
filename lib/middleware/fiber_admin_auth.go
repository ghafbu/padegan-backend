package middleware

import (
	"log"
	"strconv"
	"strings"

	_interface "github.com/ghafbu/padegan-backend/lib/interface"

	"github.com/ghafbu/padegan-backend/lib/methods"

	"github.com/ghafbu/padegan-backend/lib/consts"
	"github.com/ghafbu/padegan-backend/pkg/fiberpkg"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
	"github.com/gofiber/fiber/v3"
)

func FiberAdminAuth(
	paramJWTAccess string,
	paramDefaultKeyJWTOnRedis string,
	paramToken _interface.IToken,
	paramRedis _interface.IRedis,
	// paramOwnerName string,
	// paramExcludePath string,
) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		// shouldExclude := ctx.Path() == paramExcludePath
		// if shouldExclude {
		//	//return ctx.Next()
		// }

		// Internal Connection Check
		var thisIsInternalConnection = methods.FiberIsInternalConnection(ctx)
		if thisIsInternalConnection {
			return ctx.Next()
		}

		var (
			thisMessage []messagepkg.Response
			thisLang    = headerpkg.GetIndexAsString(ctx.GetReqHeaders()["S-Language"])
			authHeader  = ctx.GetReqHeaders()["Authorization"]
		)

		authHeaderSplit := strings.Split(strings.Join(authHeader, ""), " ")

		if len(authHeaderSplit) != _authorizationHeaderSplitLen {
			log.Println("auth-1")

			thisMessage = append(thisMessage, messagepkg.UnAuthorized(&thisLang, messagepkg.Keyword.Field["un_authorize"]))
			return fiberpkg.ReturnAuto(ctx,
				responsepkg.Code401(
					thisMessage,
					nil,
				),
			)

		}
		authToken := authHeaderSplit[1]

		authorized, err := paramToken.IsAuthorizedRaw(paramJWTAccess, authToken)
		if !authorized {
			log.Println("auth-2")

			thisMessage = append(thisMessage, messagepkg.UnAuthorized(&thisLang, messagepkg.Keyword.Field["un_authorize"]))
			return fiberpkg.ReturnAuto(ctx,
				responsepkg.Code401(
					thisMessage,
					err.Error(), // TODO - Farsi error implement
					// nil,
				),
			)
		}

		/*
		*
		*
		* Token is Valid , continue Logic
		*
		*
		 */
		// claims
		claims, err := paramToken.ParseRawUint64(paramJWTAccess, authToken)
		if err != nil {
			log.Println("auth-3")
			thisMessage = append(thisMessage, messagepkg.UnAuthorized(&thisLang, messagepkg.Keyword.Field["un-authorize"]))
			return fiberpkg.ReturnAuto(ctx,
				responsepkg.Code401(
					thisMessage,
					nil,
				))
		}

		// Extract Token To Claims Struct || Use tokenPkg/claims.go Of The Claims
		cacheToken, exist, err := paramRedis.ExistAndGetRaw(
			ctx.Context(),
			paramDefaultKeyJWTOnRedis,
			strconv.FormatUint(claims.Data, 10),
		)
		if !exist && err != nil {
			thisMessage = append(thisMessage, messagepkg.StatusCode500(&thisLang, err.Error()))
			return fiberpkg.ReturnAuto(ctx,
				responsepkg.Code500(
					thisMessage,
					err,
				),
			)
		}
		if !exist && err == nil {
			thisMessage = append(thisMessage, messagepkg.UnAuthorized(&thisLang, messagepkg.Keyword.Field["un_authorize"]))
			return fiberpkg.ReturnAuto(ctx,
				responsepkg.Code401(
					thisMessage,
					nil,
				),
			)
		}

		if authToken != cacheToken { // Request token not equal cache token
			log.Println("auth-6")

			thisMessage = append(thisMessage, messagepkg.UnAuthorized(&thisLang, messagepkg.Keyword.Field["un_authorize"]))
			return fiberpkg.ReturnAuto(ctx,
				responsepkg.Code401(
					thisMessage,
					nil,
				),
			)
		}

		log.Println("auth=7")

		ctx.Locals(consts.ContextKeyUserId, claims.Data)
		ctx.Locals(consts.ContextKeyToken, authToken)

		return ctx.Next()
	}
}
