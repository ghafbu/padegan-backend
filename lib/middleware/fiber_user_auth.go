package middleware

import (
	"strconv"
	"strings"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"

	_interface "github.com/ghafbu/padegan-backend/lib/interface"
	"github.com/ghafbu/padegan-backend/lib/methods"

	"github.com/ghafbu/padegan-backend/lib/consts"
	"github.com/ghafbu/padegan-backend/pkg/fiberpkg"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
	"github.com/gofiber/fiber/v3"
)

func FiberUserAuth(
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
		ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, "Middleware FiberUserAuth start...")

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
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, "auth-1 , HeaderSplitLen not 2")

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
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, "auth-2 , paramJWTAccess invalid || paramDefaultKeyJWTOnRedis invalid Error:", err.Error())

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
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, "auth-3 , token is valid BUT claims ParseRawUint64 ERROR:", err.Error())

			thisMessage = append(thisMessage, messagepkg.UnAuthorized(&thisLang, messagepkg.Keyword.Field["un-authorize"]))
			return fiberpkg.ReturnAuto(ctx,
				responsepkg.Code401(
					thisMessage,
					nil,
				))
		}

		ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, "auth-4-2 ", strconv.FormatUint(claims.Data, 10))
		ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, "auth-4-3 ", claims.Data)
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
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, "auth-5 , error is Nil BUT not exist")

			thisMessage = append(thisMessage, messagepkg.UnAuthorized(&thisLang, messagepkg.Keyword.Field["un_authorize"]))
			return fiberpkg.ReturnAuto(ctx,
				responsepkg.Code401(
					thisMessage,
					nil,
				),
			)
		}

		// if err != nil {
		// 	if errors.Is(err, redis.Nil) { // Not exist on cache
		// 		log.Println("auth-4")
		//
		// 		thisMessage = append(thisMessage, messagepkg.UnAuthorized(&thisLang, messagepkg.Keyword.Field["un_authorize"]))
		// 		return fiberpkg.ReturnAuto(ctx,
		// 			responsepkg.Code401(
		// 				thisMessage,
		// 				nil,
		// 			))
		// 	}
		// 	log.Println("auth-5")
		//
		// 	thisMessage = append(thisMessage, messagepkg.StatusCode500(&thisLang, err.Error()))
		// 	return fiberpkg.ReturnAuto(ctx,
		// 		responsepkg.Code500(
		// 			thisMessage,
		// 			err,
		// 		),
		// 	)
		// }

		if authToken != cacheToken { // Request token not equal cache token
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, "auth-6 , Request token not equal cache token")

			thisMessage = append(thisMessage, messagepkg.UnAuthorized(&thisLang, messagepkg.Keyword.Field["un_authorize"]))
			return fiberpkg.ReturnAuto(ctx,
				responsepkg.Code401(
					thisMessage,
					nil,
				),
			)
		}

		ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, "auth-7 END--- , create [consts.ContextKeyUserId][UserId] && [consts.ContextKeyToken][Token]")
		ctx.Locals(consts.ContextKeyUserId, claims.Data)
		ctx.Locals(consts.ContextKeyToken, authToken)

		return ctx.Next()
	}
}
