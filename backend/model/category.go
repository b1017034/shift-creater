package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	FruitsAndVegetables string
	WesternDaily        string
	Sum                 string
	ProcessedMeat       string
	Frozen              string
}
