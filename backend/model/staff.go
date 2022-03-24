package model

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	Name              string
	AbleCategory      Category
	WeekCommutingDays int
}
