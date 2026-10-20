package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/ghafbu/padegan-backend/config"

	"github.com/ghafbu/padegan-backend/lib/consts"
	_interface "github.com/ghafbu/padegan-backend/lib/interface"
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

func RateLimiter(cfg config.RateLimiterConfig, storage _interface.IRateLimiterStorage) limiter.Config {
	return limiter.Config{
		Storage:           storage,
		LimiterMiddleware: limiter.SlidingWindow{},
		KeyGenerator: func(c fiber.Ctx) string {
			return newG().GenKey(c)
		},
		Max:        cfg.Max,
		Expiration: cfg.Expiration,
	}
}

type G struct {
}

func newG() *G {
	return &G{}
}

func (rl *G) GenKey(c fiber.Ctx) string {
	if userID := c.Locals(consts.ContextKeyUserId); userID != nil {
		return fmt.Sprintf("%s%s%v",
			consts.RateLimiterPrefixKey,
			consts.RateLimiterKeyUserLoggedIn,
			userID)
	}

	return rl.strategy(c)
}

func (rl *G) strategy(c fiber.Ctx) string {
	var factors []string

	ip := rl.getRealIP(c)
	factors = append(factors, ip)

	userAgent := c.Get("User-Agent")
	factors = append(factors, userAgent)

	pathRoute := c.Path()
	factors = append(factors, pathRoute)

	methodRoute := c.Method()
	factors = append(factors, methodRoute)

	// acceptLanguage := c.Get("Accept-Language")
	// factors = append(factors, acceptLanguage)

	guestID := c.Cookies(consts.RateLimiterCookieName)
	if guestID == "" {
		guestID = rl.generateFingerprint(factors)
		rl.setGuestCookie(c)
	}

	factors = append(factors, guestID)

	fingerprint := rl.generateFingerprint(factors)

	return fmt.Sprintf(
		consts.RateLimiterPattern,
		consts.RateLimiterPrefixKey,
		consts.RateLimiterKeyGuest,
		fingerprint,
	)
}

func (rl *G) generateFingerprint(factors []string) string {
	return strings.Join(factors, consts.RateLimiterSeparator)
}

func (rl *G) getRealIP(c fiber.Ctx) string {
	if ip := c.Get("X-Forwarded-For"); ip != "" {
		ips := strings.Split(ip, ",")

		return strings.TrimSpace(ips[0])
	}

	if ip := c.Get("X-Real-IP"); ip != "" {
		return ip
	}

	return c.IP()
}

func (rl *G) setGuestCookie(c fiber.Ctx) {
	cookie := &fiber.Cookie{
		Name:     consts.RateLimiterCookieName,
		Value:    ghafpkg.GenUUID(),
		Expires:  time.Now().Add(consts.RateLimiterCookieExpireTime),
		HTTPOnly: consts.RateLimiterCookieHTTPOnly,
		Secure:   consts.RateLimiterCookieSecure,
		SameSite: consts.RateLimiterCookieSameSite,
		// Path:     "/",
	}

	c.Cookie(cookie)
}
