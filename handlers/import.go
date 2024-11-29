package handlers

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	"analytics_tool/db"
	"analytics_tool/models"
)

// ImportPositions импортирует должности из CSV файла
func ImportPositions(filePath string) {
	index := 0
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range records {
		hourlyRate, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		position := models.Position{Name: row[0], HourlyRate: hourlyRate}
		// Проверка на существование записи и добавление, если не существует
		if err := db.DB.FirstOrCreate(&position, models.Position{Name: row[0]}).Error; err != nil {
			log.Fatalf("Failed to create position: %v", err)
		}
		index++
	}
	log.Printf("%d Raw inserted\n", index)
}

// ImportEmployees импортирует сотрудников из CSV файла
func ImportEmployees(filePath string) {
	index := 0
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range records {
		var position models.Position
		if err := db.DB.Where("name = ?", row[1]).First(&position).Error; err != nil {
			log.Printf("Position not found: %v", err)
			continue
		}
		employee := models.Employee{Name: row[0], PositionID: position.ID}
		// Проверка на существование записи и добавление, если не существует
		if err := db.DB.FirstOrCreate(&employee, models.Employee{Name: row[0]}).Error; err != nil {
			log.Fatalf("Failed to create employee: %v", err)
		}
		index++
	}
	log.Printf("%d Raw inserted\n", index)
}

// ImportTimesheet импортирует таймшиты из CSV файла
func ImportTimesheet(filePath string) {
	var index int
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range records {
		// var task models.Task
		index = 0
		task := models.Task{Title: row[0]}
		// Проверка на существование записи и добавление, если не существует
		if err := db.DB.FirstOrCreate(&task, models.Task{Title: row[0]}).Error; err != nil {
			log.Fatalf("Failed to create timesheet: %v", err)
		}

		index++
	}
	log.Printf("%d Raw inserted in tasks\n", index)

	for _, row := range records {
		task := models.Task{Title: row[0]}
		var employee models.Employee
		index = 0
		if err := db.DB.Where("name = ?", row[1]).First(&employee).Error; err != nil {
			log.Printf("Employee not found: %v", err)
			continue
		}

		if err := db.DB.Where("title = ?", row[0]).First(&task).Error; err != nil {
			log.Printf("Task not found: %v", err)
			continue
		}

		startTime, err := time.Parse("2006-01-02 15:04:05", row[2])
		if err != nil {
			log.Fatalf("Failed to parse start_time: %v", err)
		}

		endTime, err := time.Parse("2006-01-02 15:04:05", row[3])
		if err != nil {
			log.Fatalf("Failed to parse end_time: %v", err)
		}

		timesheet := models.Timesheet{EmployeeID: employee.ID, TaskID: task.ID, StartTime: startTime, EndTime: endTime}
		// Проверка на существование записи и добавление, если не существует
		if err := db.DB.FirstOrCreate(&timesheet, models.Timesheet{EmployeeID: employee.ID, TaskID: task.ID, StartTime: startTime, EndTime: endTime}).Error; err != nil {
			log.Fatalf("Failed to create timesheet: %v", err)
		}
		index++
	}
	log.Printf("%d Raw inserted in Timesheets\n", index)
}
