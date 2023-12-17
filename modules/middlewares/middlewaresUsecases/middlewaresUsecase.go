package middlewaresusecases

import "github.com/chonlasit2000/kawaii-shop/modules/middlewares/middlewaresRepositories"

type IMiddlewareUseCase interface {
}

type middlewareUseCase struct {
	middlewareRepository middlewaresRepositories.IMiddlewareRepository
}

func MiddlewareUseCase(middlewareRepository middlewaresRepositories.IMiddlewareRepository) IMiddlewareUseCase {
	return &middlewareUseCase{
		middlewareRepository: middlewareRepository,
	}
}
