package handlers

import (
	"analytics_tool/db"
	"analytics_tool/models"
	"fmt"
	"log"
)

// RemoveEmployee удаляет данные по сотруднику из таймшита по его имени
func RemoveEmployee(employeeName string) {
	var employee models.Employee
	if err := db.DB.Where("name = ?", employeeName).First(&employee).Error; err != nil {
		log.Printf("Employee not found: %v", err)
		return
	}

	if err := db.DB.Where("employee_id = ?", employee.ID).Delete(&models.Timesheet{}).Error; err != nil {
		log.Printf("Failed to remove timesheets: %v", err)
		return
	}

	fmt.Printf("Timesheets for employee %s removed successfully\n", employeeName)
}
