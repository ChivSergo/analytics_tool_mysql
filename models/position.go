package models

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	Name       string
	HourlyRate float64
}

// Save сохраняет запись должности в базу данных
func (p *Position) Save(db *gorm.DB) error {
	return db.Create(p).Error
}

// GetList возвращает список всех должностей
func GetPositions(db *gorm.DB) ([]Position, error) {
	var positions []Position
	if err := db.Find(&positions).Error; err != nil {
		return nil, err
	}
	return positions, nil
}
