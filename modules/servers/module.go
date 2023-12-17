package servers

import (
	middlewareshandler "github.com/chonlasit2000/kawaii-shop/modules/middlewares/middlewaresHandler"
	"github.com/chonlasit2000/kawaii-shop/modules/middlewares/middlewaresRepositories"
	middlewaresusecases "github.com/chonlasit2000/kawaii-shop/modules/middlewares/middlewaresUsecases"
	"github.com/chonlasit2000/kawaii-shop/modules/monitor/monitorHandlers"
	"github.com/gofiber/fiber/v2"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	router fiber.Router
	server *server
	mid    middlewareshandler.IMiddlewareHandler
}

func InitModule(r fiber.Router, s *server, mid middlewareshandler.IMiddlewareHandler) IModuleFactory {
	return &moduleFactory{
		router: r,
		server: s,
		mid:    mid,
	}
}

func InitMiddlewares(s *server) middlewareshandler.IMiddlewareHandler {
	repository := middlewaresRepositories.MiddlewareRepository(s.db)
	usecase := middlewaresusecases.MiddlewareUseCase(repository)
	return middlewareshandler.MiddlewareHandler(s.cfg, usecase)
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.server.cfg)

	m.router.Get("/", handler.HealthCheck)
}
