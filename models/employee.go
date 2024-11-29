package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name       string
	Position   Position
	PositionID uint
}

// Save сохраняет запись сотрудника в базу данных
func (e *Employee) Save(db *gorm.DB) error {
	return db.Create(e).Error
}

// GetList возвращает список всех сотрудников
func GetEmployees(db *gorm.DB) ([]Employee, error) {
	var employees []Employee
	if err := db.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}
