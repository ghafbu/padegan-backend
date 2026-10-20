package fiberpkg

import (
	"fmt"

	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
	"github.com/gofiber/fiber/v3"
)

func ReturnAuto(ctx fiber.Ctx, paramResponse responsepkg.EntityV1UseCase) error {
	// Logger response
	// var thisPrimaryResponseInit = responsepkg.NewV1()
	// var thisPrimaryResponseJson = thisPrimaryResponseInit.SetStatus(paramResponse.Status).
	//	SetStatusCode(paramResponse.StatusCode).
	//	SetMessage(paramResponse.Message).
	//	SetData(paramResponse.Data).
	//	SetPath(ctx.OriginalURL()).
	//	SetExecutionDuration(GetExecuteDuration(ctx)).
	//	Build()
	// log(thisPrimaryResponseJson)
	if paramResponse.Data != nil { // is not nil or empty paramResponse.Data
		ctx.Locals(
			ContextKeyData,
			fmt.Sprintf("%v", paramResponse.Data),
		) // 'data', usage in logger middleware
	}

	if paramResponse.Message != nil {
		ctx.Locals(
			ContextKeyMessage,
			fmt.Sprintf("%v", paramResponse.Message),
		) // 'message', usage in logger middleware
	}

	// Back-end response
	// fmt.Println("paramResponse", paramResponse)
	var thisBackEndResponse = responsepkg.NewV2(paramResponse)
	// fmt.Println("hisBackEndResponse", thisBackEndResponse)

	return ctx.Status(paramResponse.StatusCode).JSON(thisBackEndResponse)

	// Way 1
	// return ReturnPublic().
	//	SetStatus(_res.Status).
	//	SetStatusCode(_res.StatusCode).
	//	SetMessage(_res.Message).
	//	SetData(_res.Data).
	//	build(ctx)

	// thisResponse=thisResponse.
	// Way 2
	// var thisResponse = responsepkg.NewV1().
	//	SetStatus(_res.Status).
	//	SetStatusCode(_res.StatusCode).
	//	SetMessage(_res.Message).
	//	SetData(_res.Data).
	//	SetPath(ctx.OriginalURL()).
	//	SetExecutionDuration(GetExecuteDuration(ctx)).
	//	Build()

	// Way 3
	// var c func(ctx fiber.Ctx, _mess any, _data ...any) error
	//
	// if _res.StatusCode == 200 {
	//	c = ReturnOk
	// } else if _res.StatusCode == 201 {
	//	c = ReturnCreated
	// } else if _res.StatusCode == 400 {
	//	c = ReturnBadRequest
	// } else if _res.StatusCode == 500 {
	//	c = ReturnInternalServerError
	// }
	//
	// return c(ctx, _res.Message, _res.Data)

	// Way 4
	// if _res.StatusCode == 200 {
	//	return ReturnOk(ctx, _res.Message, _res.Data)
	// } else if _res.StatusCode == 201 {
	//	return ReturnCreated(ctx, _res.Message, _res.Data)
	// } else if _res.StatusCode == 400 {
	//	return ReturnBadRequest(ctx, _res.Message, _res.Data)
	// } else if _res.StatusCode == 500 {
	//	return ReturnInternalServerError(ctx, _res.Message, _res.Data)
	// }
	//
	// return ReturnInternalServerError(ctx, ghafpkg.AppGlobalErrorReturnAutoNotFind, _res)
}
