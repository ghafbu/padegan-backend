package middleware

//import (
//	"github.com/ghafbu/padegan-backend/lib/consts"
//	_interface "github.com/ghafbu/padegan-backend/lib/interface"
//	"github.com/gofiber/fiber/v3"
//	"github.com/gofiber/fiber/v3/middleware/logger"
//	"github.com/gofiber/fiber/v3/middleware/requestid"
//)
//
//func FiberLogger(
//	paramLogger _interface.ILogger,
//	paramServiceName string,
//	paramExcludePath string,
//) func(ctx fiber.Ctx) error {
//	return logger.New(logger.Config{
//		LoggerFunc: func(ctx fiber.Ctx, data *logger.Data, cfg logger.Config) error {
//			shouldExclude := ctx.Path() == paramExcludePath
//			if shouldExclude {
//				return ctx.Next()
//			}
//
//			thisData := ""
//			if ctx.Locals(consts.ContextKeyData) != nil {
//				thisData = ctx.Locals(consts.ContextKeyData).(string)
//			}
//
//			thisMessage := ""
//			if ctx.Locals(consts.ContextKeyMessage) != nil {
//				thisMessage = ctx.Locals(consts.ContextKeyMessage).(string)
//			}
//
//			paramLogger.Info(
//				[]string{paramServiceName, "MiddlewareLogger"},
//				map[string]interface{}{
//					"service_name":    paramServiceName,
//					"uri":             ctx.Route().Path,
//					"method":          ctx.Method(),
//					"query_string":    ctx.Queries(),
//					"request_body":    string(ctx.Request().Body()),
//					"status":          ctx.Response().StatusCode(),
//					"response_body":   string(ctx.Response().Body()),
//					"content-length":  len(ctx.Response().Body()),
//					"remote_ip":       ctx.IP(),
//					"http.user_agent": string(ctx.Request().Header.UserAgent()),
//					"host":            ctx.Host(),
//					"request_id":      requestid.FromContext(ctx),
//					"message":         thisMessage,
//					"data":            thisData,
//					"protocol":        ctx.Protocol(),
//				},
//				"--",
//			)
//
//			return nil
//		},
//		DisableColors: false,
//	})
//}
