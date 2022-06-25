package domain

import "context"

type PetsRepository interface {
	Create(ctx context.Context, pet Pet) error
	Find(ctx context.Context, id string) (Pet, error)
	List(ctx context.Context, status byte) ([]Pet, error)
	Count(ctx context.Context, status byte) (int64, error)
	Update(ctx context.Context, id string, in UpdatePetCadasterInput) error
	Delete(ctx context.Context, id string)
}
