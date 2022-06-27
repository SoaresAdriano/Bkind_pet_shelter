package domain

import "time"

// Vaccines is a list of vaccine
type Vaccines []Vaccine

// Vaccine encapsulates the mainly characteristics of pet vaccination
type Vaccine struct {
	Name      string
	AppliedAt time.Time
	NextApply time.Time
}
