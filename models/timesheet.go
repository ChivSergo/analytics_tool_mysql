package models

import (
	"time"

	"gorm.io/gorm"
)

type Timesheet struct {
	gorm.Model
	EmployeeID uint
	TaskID     uint
	StartTime  time.Time
	EndTime    time.Time
}

// Save сохраняет запись таймшита в базу данных
func (t *Timesheet) Save(db *gorm.DB) error {
	return db.Create(t).Error
}

// GetList возвращает список всех таймшитов
func GetTimesheets(db *gorm.DB) ([]Timesheet, error) {
	var timesheets []Timesheet
	if err := db.Find(&timesheets).Error; err != nil {
		return nil, err
	}
	return timesheets, nil
}
