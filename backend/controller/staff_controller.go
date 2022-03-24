package controller

import (
	"encoding/json"
	"fmt"
	"shift-creater/backend/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StaffController struct {
	db *gorm.DB
}

func NewStaffController() *StaffController {
	db, err := gorm.Open(sqlite.Open("shift.db"), &gorm.Config{})
	if err != nil {
		panic(nil)
	}

	return &StaffController{
		db: db,
	}
}

func (s StaffController) Create(staff map[string]interface{}) model.Staff {
	staffStruct := model.Staff{}
	// Convert map to json string
	jsonStr, err := json.Marshal(staff)
	if err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(jsonStr, &staffStruct); err != nil {
		fmt.Println(err)
	}
	s.db.Create(&staffStruct)
	return staffStruct
}

func (s StaffController) First(staff model.Staff) model.Staff {
	s.db.First(&staff)

	return staff
}

func (s StaffController) Update(id int, updateStaff model.Staff) {
	s.db.Save(updateStaff)
}
