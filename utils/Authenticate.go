package utils

import (
	"empsystem/model"
	
)

func Authenticate(email, password string) (bool, *models.Employee) {
	employees, err := ReadCSV()
	if err != nil {
		// Handle error reading CSV file
		return false, nil
	}

	for _, emp := range employees {
		if emp.Email == email && emp.Password == password {
			// Return authenticated employee if email and password match
			return true, &emp
		}
	}

	// Return false and nil employee if no matching credentials are found
	return false, nil
}




