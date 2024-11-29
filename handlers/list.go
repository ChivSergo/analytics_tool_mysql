package handlers

import (
	"analytics_tool/db"
	"analytics_tool/models"
	"fmt"
	"log"
)

// ListEmployees выводит список сотрудников
func ListEmployees() {
	index := 0
	var employees []models.Employee
	if err := db.DB.Find(&employees).Error; err != nil {
		log.Fatalf("Failed to list employees: %v", err)
	}
	for _, employee := range employees {
		fmt.Println(employee.Name)
		index++
	}
	fmt.Printf("%d Records loaded successfully\n", index)
}
