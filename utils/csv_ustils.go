package utils

import (
	"encoding/csv"
	"os"
	"strconv"
	"sync"

	models "empsystem/model"
)

const csvFilePath = "D:/Gotask2/utils/employees.csv"

var csvMutex sync.Mutex

func InitCSV() error {
	// Check if the CSV file exists, if not, create it with headers
	_, err := os.Stat(csvFilePath)
	if os.IsNotExist(err) {
		file, err := os.Create(csvFilePath)
		defer file.Close()
		if err != nil {
			return err
		}
		writer := csv.NewWriter(file)
		defer writer.Flush()

		// add password field here
		headers := []string{"ID", "FirstName", "LastName", "Email", "Password", "Role"}
		err = writer.Write(headers)
		if err != nil {
			return err
		}
	}
	return nil
}

// Read from CSV file

func ReadCSV() ([]models.Employee, error) {
	csvMutex.Lock()
	defer csvMutex.Unlock()

	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var employees []models.Employee
	for i, line := range lines {
		if i == 0 { // Skip header row
			continue
		}
		id, _ := strconv.Atoi(line[0])
		employee := models.Employee{
			ID:        id,
			FirstName: line[1],
			LastName:  line[2],
			Email:     line[3],
			Password:  line[4],
			Role:      line[5],
		}
		employees = append(employees, employee)
	}

	return employees, nil
}


// Write to the CSV file

func WriteToCSV(employees []models.Employee) error {
	csvMutex.Lock()
	defer csvMutex.Unlock()

	file, err := os.OpenFile(csvFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header row
	headers := []string{"ID", "FirstName", "LastName", "Email", "Password","Role"}
	err = writer.Write(headers)
	if err != nil {
		return err
	}

	// Write employee data
	for _, employee := range employees {
		err := writer.Write([]string{
			strconv.Itoa(employee.ID),
			employee.FirstName,
			employee.LastName,
			employee.Email,
			employee.Password,
			employee.Role,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

