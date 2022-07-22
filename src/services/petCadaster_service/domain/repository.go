package domain

import (
	"bkind_pet_shelter/src/services/petCadaster_service/domain/entity"
	"context"
)

type PetsRepository interface {
	Create(ctx context.Context, pet entity.Pet) error
	Find(ctx context.Context, id string) (entity.Pet, error)
	List(ctx context.Context, status byte) ([]entity.Pet, error)
	Count(ctx context.Context, status byte) (int64, error)
	Update(ctx context.Context, id string, in UpdatePetCadasterInput) error
	Delete(ctx context.Context, id string)
}
