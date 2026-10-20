package fiberpkg

import (
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
	"github.com/gofiber/fiber/v3"
)

type CloneReturnPublic struct {
	Response responsepkg.EntityV1
}

func ReturnPublic() *CloneReturnPublic {
	var clone = CloneReturnPublic{}
	return &clone
}

func (r CloneReturnPublic) SetStatus(status bool) CloneReturnPublic {
	r.Response.Status = status
	return r
}

func (r CloneReturnPublic) SetStatusCode(statusCode int) CloneReturnPublic {
	r.Response.StatusCode = statusCode
	return r
}

func (r CloneReturnPublic) SetMessage(message interface{}) CloneReturnPublic {
	r.Response.Message = message
	return r
}

func (r CloneReturnPublic) SetData(data interface{}) CloneReturnPublic {
	r.Response.Data = data
	return r
}

func (r CloneReturnPublic) build(ctx fiber.Ctx) error {
	r.Response.Path = ctx.OriginalURL()
	r.Response.ExecuteDuration = GetExecuteDuration(ctx)

	return ctx.Status(r.Response.StatusCode).JSON(r.Response)
}
