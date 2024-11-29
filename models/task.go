package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title string
}

// Save сохраняет запись задачи в базу данных
func (t *Task) Save(db *gorm.DB) error {
	return db.Create(t).Error
}

// GetList возвращает список всех задач
func GetTasks(db *gorm.DB) ([]Task, error) {
	var tasks []Task
	if err := db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
