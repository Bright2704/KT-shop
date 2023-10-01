package midlewaresUsecases

import "github.com/Bright2704/KT-shop-tutorial/modules/middlewares/middlewaresRepositories"




type IMiddlewaresUsecase interface {

}

type middlewaresUsecase struct {
	middlewaresRepository middlewaresRepositories.IMiddlewaresRepository
}

func MiddlewaresUsecase(middlewaresRepository middlewaresRepositories.IMiddlewaresRepository) IMiddlewaresUsecase{
	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}
}