package http

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v3/middleware/static"

	"github.com/ghafbu/padegan-backend/lib/bootstrap"
	"github.com/ghafbu/padegan-backend/pkg/validatepkg/govalidatorcustom"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v3"
)

type Server struct {
	app   *bootstrap.Application
	fiber *fiber.App
}

func New(paramApp *bootstrap.Application) Server {
	return Server{
		app: paramApp,
	}
}

func (s Server) Init() error {
	// **************** Initialize Fiber Framework ***************

	// Init Template Engine
	thisEngine := html.New(s.app.Config.ViewInfo.Path,
		s.app.Config.ViewInfo.Extension) // For development
	// var thisEngine := html.NewFileSystem(http.FS(viewsfs), ".html") // For production

	s.fiber = fiber.New(
		fiber.Config{
			Views: thisEngine,
			StructValidator: govalidatorcustom.NewBindStructTags(s.app.Config.Other.PasswordEntropy).
				Setup(),
		},
	)

	// Usage Assets File
	s.fiber.Use("/", static.New(s.app.Config.StaticFile.Path))

	// ************* End Of Initialize Fiber Framework **************

	// ************* Configure Services ****************
	s.UsersConfigure()

	return s.fiber.Listen(fmt.Sprintf(":%s", s.app.Config.HTTP.PortHttp))
}

func (s Server) ShutdownWithContext(paramCtx context.Context) error {
	return s.fiber.ShutdownWithContext(paramCtx)
}
