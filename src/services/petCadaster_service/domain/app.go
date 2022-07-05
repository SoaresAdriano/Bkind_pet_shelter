package domain

import (
	"context"
)

// Application defines the contracts with the service application layer
type Application interface {
	FindPet(ctx context.Context, in FindPetInput) (Pet, []error)
	CadasterPet(ctx context.Context, in CadasterPetInput) (string, []error)
	ListPets(ctx context.Context, in ListPetsInput) ([]Pet, []error)
	CountPets(ctx context.Context, in CountPetsInput) (uint64, []error)
	UpdateCadaster(ctx context.Context, in UpdatePetCadasterInput) []error
	DeleteCadaster(ctx context.Context, in DeletePetCadasterInput) []error
	RenamePet(ctx context.Context, in RenamePetInput) []error
}

// FindPetCadasterInput  encapsulate the necessary field to perform a  FindPet command
type FindPetInput struct {
	ID string
}

// CadasterPetInput encapsulates the necessary fields to perform a pet cadaster
type CadasterPetInput struct {
	Name      string
	Species   string
	Breed     string
	Sex       byte
	Age       uint8
	Weight    float32
	Size      byte
	Color     string
	Neutered  byte
	Origin    string
	Situation byte
}

// ListPetsInput encapsulates the necessary fields to perform a pet list according with species command
type ListPetsInput struct {
	Species string
}

// CountPetsInput encapsulates the necessary fields to perform a pet count command
type CountPetsInput struct {
	Species string
}

// UpdateCadasterInput encapsulates the necessary fields to perform a pet update cadaster command
type UpdatePetCadasterInput struct {
	Pet *Pet
}

// DeleteCadasterInput encapsulates the necessary fields to perform a pet delete command
type DeletePetCadasterInput struct {
	ID string
}

// RenamePetInput encapsulates the necessary fields to perform a pet rename command
type RenamePetInput struct {
	ID   string
	Name string
}
