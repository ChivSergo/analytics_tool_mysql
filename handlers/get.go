package handlers

import (
	"analytics_tool/db"
	"analytics_tool/models"
	"fmt"
	"log"
)

// GetEmployeeTimesheet выводит таймшиты сотрудника по его имени
func GetEmployeeTimesheet(employeeName string) {
	var employee models.Employee

	index := 0
	if err := db.DB.Where("name = ?", employeeName).First(&employee).Error; err != nil {
		log.Printf("Employee not found: %v", err)
		return
	}

	var timesheets []models.Timesheet
	if err := db.DB.Where("employee_id = ?", employee.ID).Find(&timesheets).Error; err != nil {
		log.Printf("Failed to get timesheets: %v", err)
		return
	}

	// for timesheet := range timesheets {
	// 	var task models.Task
	// 	if err := db.DB.Where("id in ?", timesheet.TaskID).First(&task).Error; err != nil {
	// 		log.Fatalf("Failed to get task: %v", err)
	// 	}

	// }

	for _, timesheet := range timesheets {
		var task models.Task
		if err := db.DB.Where("id = ?", timesheet.TaskID).First(&task).Error; err != nil {
			log.Fatalf("Failed to get task: %v", err)
		}
		fmt.Printf("Task Name: %s, Start Time: %s, End Time: %s\n", task.Title, timesheet.StartTime, timesheet.EndTime)
		index++
	}

	fmt.Printf("%d Records loaded successfully\n", index)

}
