package model

import "gorm.io/gorm"

type ShiftRequest struct {
	gorm.Model
	OffDay int
}
