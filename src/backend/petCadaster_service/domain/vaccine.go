package domain

import "time"

type Vaccines []Vaccine

type Vaccine struct {
	Name      string
	ApplyedAt time.Time
	NextApply time.Time
}
