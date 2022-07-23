package app

import "bkind_pet_shelter/src/services/petCadaster_service/domain"

type AppParams struct {
	PetsRepository domain.PetsRepository
}

type app struct {
	petsRepository domain.PetsRepository
}

func NewApp(params AppParams) (domain.Application, error) {
	if params.PetsRepository == nil {
		return nil, domain.ErrMissingPetsRepositoryDependency
	}
	return app{
		petsRepository: params.PetsRepository,
	}, nil
}

func MustNewApp(params AppParams) domain.Application {
	app, err := NewApp(params)
	if err != nil {
		panic(err)
	}
	return app
}
