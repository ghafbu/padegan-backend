package fiberpkg

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

func InitSession() fiber.Handler {
	return session.New()
}

func SessionFromContext(ctx fiber.Ctx) *session.Middleware {
	return session.FromContext(ctx)
}

func SessionGetAndDelete[T any](sess *session.Middleware, key string) (T, bool) {
	var zero T

	v := sess.Get(key)
	if v == nil {
		return zero, false
	}

	val, ok := v.(T)
	if !ok {
		return zero, false
	}

	sess.Delete(key)

	return val, true
}

func SessionGet[T any](sess *session.Middleware, key string) (T, bool) {
	var zero T

	v := sess.Get(key)
	if v == nil {
		return zero, false
	}

	val, ok := v.(T)
	if !ok {
		return zero, false
	}

	return val, true
}

func SessionDelete(sess *session.Middleware, key string) {
	sess.Delete(key)
}
