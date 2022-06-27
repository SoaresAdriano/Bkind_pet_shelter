package app

import (
	"bkind_pet_shelter/src/backend/petCadaster_service/domain"
)

type ApplicationParams struct {
	PetsRepo domain.PetsRepository
	// Validator domain.Validator
}

type application struct {
	petsRepo domain.PetsRepository
	// validator domain.Validator
}

func NewApplication(params ApplicationParams) (domain.Application, error) {
	if params.PetsRepo == nil {
		return nil, domain.ErrMissingPetsRepositoryDependency
	}
	// if params.Validator == nil {
	// 	return nil, domain.ErrMissingValidatorDependency
	// }
	return application{
		petsRepo: params.PetsRepo,
	}, nil
}
