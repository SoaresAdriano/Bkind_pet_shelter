package domain

import (
	"time"

	"github.com/google/uuid"
)

// Pet is a set of characteristics that defines a pet
type Pet struct {
	ID           string
	Name         string
	Species      string
	Breed        string
	Gender       byte
	Age          uint8
	Weight       float32
	Size         byte
	Color        string
	Neutered     byte
	Origin       string
	Status       byte
	Vaccines     Vaccines
	RegisteredAt time.Time
}

// NewPet creates a new pet
func NewPet(name string, species string, breed string, gender byte, age uint8, weight float32, size byte, color string, neutered byte, origin string, status byte, vaccines Vaccines) Pet {
	now := time.Now()

	return Pet{
		ID:           uuid.New().String(),
		Name:         name,
		Species:      species,
		Breed:        breed,
		Gender:       gender,
		Age:          age,
		Weight:       weight,
		Size:         size,
		Color:        color,
		Neutered:     neutered,
		Origin:       origin,
		Status:       status,
		Vaccines:     vaccines,
		RegisteredAt: now,
	}
}
