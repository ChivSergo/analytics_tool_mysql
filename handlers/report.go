package handlers

import (
	"analytics_tool/db"
	"analytics_tool/models"
	"fmt"
	"log"
	"time"
)

// ReportTop5LongTasks выводит пять задач, на которые потрачено больше всего времени
func ReportTop5LongTasks() {
	var timesheets []models.Timesheet
	if err := db.DB.Order("end_time - start_time DESC").Limit(5).Find(&timesheets).Error; err != nil {
		log.Fatalf("Failed to get top 5 long tasks: %v", err)
	}

	for _, timesheet := range timesheets {
		var task models.Task
		if err := db.DB.Where("id = ?", timesheet.TaskID).First(&task).Error; err != nil {
			log.Fatalf("Failed to get task: %v", err)
		}
		duration := timesheet.EndTime.Sub(timesheet.StartTime)
		fmt.Printf("Task name: %s, Duration: %s\n", task.Title, duration)
	}
}

// ReportTop5CostTasks выводит пять задач, на которые потрачено больше всего денег
func ReportTop5CostTasks() {
	var timesheets []models.Timesheet
	var position models.Position
	var employeeIDs []uint

	// Получаем должность с самой высокой ставкой
	if err := db.DB.Order("hourly_rate desc").First(&position).Error; err != nil {
		log.Fatalf("Failed to get top position: %v", err)
	}

	// Получаем идентификаторы сотрудников с этой должностью
	if err := db.DB.Model(&models.Employee{}).Select("id").Where("position_id = ?", position.ID).Find(&employeeIDs).Error; err != nil {
		log.Fatalf("Failed to get employee IDs: %v", err)
	}

	// Получаем таймшиты для этих сотрудников и сортируем по стоимости
	if err := db.DB.Where("employee_id IN (?)", employeeIDs).Order("end_time - start_time DESC").Limit(5).Find(&timesheets).Error; err != nil {
		log.Fatalf("Failed to get top 5 cost tasks: %v", err)
	}

	for _, timesheet := range timesheets {
		var task models.Task
		if err := db.DB.Where("id = ?", timesheet.TaskID).First(&task).Error; err != nil {
			log.Fatalf("Failed to get task: %v", err)
		}
		duration := timesheet.EndTime.Sub(timesheet.StartTime).Hours()
		cost := duration * position.HourlyRate
		fmt.Printf("Task Name: %s, Cost: %f\n", task.Title, cost)
	}
}

// ReportTop5Employees выводит пять сотрудников, отработавших наибольшее количество времени за всё время
func ReportTop5Employees() {
	var employees []models.Employee
	var timesheets []models.Timesheet

	// Получаем всех сотрудников
	if err := db.DB.Find(&employees).Error; err != nil {
		log.Fatalf("Failed to get employees: %v", err)
	}

	// Получаем все таймшиты
	if err := db.DB.Find(&timesheets).Error; err != nil {
		log.Fatalf("Failed to get timesheets: %v", err)
	}

	// Создаем мапу для хранения общего времени для каждого сотрудника
	employeeTotalTime := make(map[uint]time.Duration)

	// Вычисляем общее время для каждого сотрудника
	for _, timesheet := range timesheets {
		employeeTotalTime[timesheet.EmployeeID] += timesheet.EndTime.Sub(timesheet.StartTime)
	}

	// Сортируем сотрудников по общему времени
	type EmployeeTotalTime struct {
		Employee  models.Employee
		TotalTime time.Duration
	}

	var employeeTotalTimes []EmployeeTotalTime

	for _, employee := range employees {
		employeeTotalTimes = append(employeeTotalTimes, EmployeeTotalTime{
			Employee:  employee,
			TotalTime: employeeTotalTime[employee.ID],
		})
	}

	// Сортировка сотрудников по общему времени
	for i := 0; i < len(employeeTotalTimes)-1; i++ {
		for j := i + 1; j < len(employeeTotalTimes); j++ {
			if employeeTotalTimes[i].TotalTime < employeeTotalTimes[j].TotalTime {
				employeeTotalTimes[i], employeeTotalTimes[j] = employeeTotalTimes[j], employeeTotalTimes[i]
			}
		}
	}

	// Выводим пять сотрудников, отработавших наибольшее количество времени
	for i := 0; i < 5 && i < len(employeeTotalTimes); i++ {
		fmt.Printf("Employee: %s, Total Time: %s\n", employeeTotalTimes[i].Employee.Name, employeeTotalTimes[i].TotalTime)
	}
}
