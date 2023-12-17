package middlewareshandler

import (
	"github.com/chonlasit2000/kawaii-shop/config"
	"github.com/chonlasit2000/kawaii-shop/modules/entities"
	middlewaresusecases "github.com/chonlasit2000/kawaii-shop/modules/middlewares/middlewaresUsecases"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type middlewareHandlerErrCode string

const (
	RouterCheckErr middlewareHandlerErrCode = "middleware-001"
)

type IMiddlewareHandler interface {
	Cors() fiber.Handler
	RouterCheck() fiber.Handler
	Logger() fiber.Handler
}

type middlewareHandler struct {
	cfg               config.IConfig
	middlewareUsecase middlewaresusecases.IMiddlewareUseCase
}

func MiddlewareHandler(cfg config.IConfig, middlewareRepository middlewaresusecases.IMiddlewareUseCase) IMiddlewareHandler {
	return &middlewareHandler{
		cfg:               cfg,
		middlewareUsecase: middlewareRepository,
	}
}

func (h *middlewareHandler) Cors() fiber.Handler {
	return cors.New(cors.Config{
		Next:             cors.ConfigDefault.Next,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	},
	)
}

func (h *middlewareHandler) RouterCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return entities.NewResponse(c).Error(
			fiber.ErrNotFound.Code,
			string(RouterCheckErr),
			"router not found",
		).Res()
	}
}

func (h *middlewareHandler) Logger() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "${time} ${ip} ${status} - ${method} ${path}\n",
		TimeFormat: "02/01/2006",
		TimeZone:   "Bangkok/Asia",
	})
}
