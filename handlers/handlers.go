
package handlers

import (
	models "empsystem/model"
	"empsystem/utils"
	"encoding/json"
	"net/http"
	"strconv"
    
	"github.com/gorilla/mux"
	
)



func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	authenticated, authenticatedEmployee := utils.Authenticate(loginRequest.Email, loginRequest.Password)
	if authenticated {
		// Authentication successful, respond with authenticated employee data
		// respondWithJSON(w, http.StatusOK, authenticatedEmployee)


		token, err := utils.GenerateJWT(authenticatedEmployee.Email, authenticatedEmployee.Role)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to generate JWT token")
			return
		}
		respondWithJSON(w, http.StatusOK, map[string]string{"token": token})

	} else {
		// Authentication failed, respond with unauthorized status
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
	}
}



// GetAllEmployees gets a list of all employees.

// @Summary Get all employees
// @Description Get a list of all employees
// @Tags employees
// @Accept json
// @Produce json
// @Success 200 {array} model.Employee
// @Router /employees [get]




func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
    
	
	employees, err := utils.ReadCSV()
	if err != nil {
		// utils.InfoLogger.Println("Error reading CSV file:", err)
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	respondWithJSON(w, http.StatusOK, employees)
}

/*GetEmployeeByID gets an employee by ID.
// @Summary Get employee by ID
// @Description Get an employee by ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} model.Employee
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees/{id} [get]

*/


func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid employee ID")
		return
	}

	
	employees, err := utils.ReadCSV()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	var foundEmployee *models.Employee
	for _, employee := range employees {
		if employee.ID == id {
			foundEmployee = &employee
			break
		}
	}

	if foundEmployee == nil {
		respondWithError(w, http.StatusNotFound, "Employee not found")
		return
	}

	respondWithJSON(w, http.StatusOK, foundEmployee)
}



/* SearchEmployees searches employees by field and value.
// @Summary Search employees
// @Description Search employees by field and value
// @Tags employees
// @Accept json
// @Produce json
// @Param field query string true "Field to search (e.g., FirstName)"
// @Param value query string true "Value to search"
// @Success 200 {array} model.Employee
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees/search [get]

*/




func SearchEmployees(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	field := queryValues.Get("field")
	value := queryValues.Get("value")

	employees, err := utils.ReadCSV()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	var filteredEmployees []models.Employee
	for _, employee := range employees {
		// Match the field and value based on your logic
		switch field {
		case "FirstName":
			if employee.FirstName == value {
				filteredEmployees = append(filteredEmployees, employee)
			}
		case "LastName":
			if employee.LastName == value {
				filteredEmployees = append(filteredEmployees, employee)
			}
			// Add more cases for other fields if needed email & role
		}
	}

	respondWithJSON(w, http.StatusOK, filteredEmployees)
}


/* CreateEmployeeHandler creates a new employee.
// @Summary Create a new employee
// @Description Create a new employee
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body model.Employee true "Employee object to be created"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees [post]

*/


func CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {

	
	var employee models.Employee
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Check if the user is an admin
	if employee.Role != "admin" {
		respondWithError(w, http.StatusForbidden, "Forbidden: Only admins can create employees")
		return
	}

	// Read existing employees from CSV
	existingEmployees, err := utils.ReadCSV()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to read employee data")
		return
	}

	// Generate a unique ID for the new employee
	newEmployeeID := generateUniqueEmployeeID(existingEmployees)

	// Set the new employee ID
	employee.ID = newEmployeeID

	// Add the new employee to the existing employees slice
	existingEmployees = append(existingEmployees, employee)

	// Write the updated employees to CSV
	err = utils.WriteToCSV(existingEmployees)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update employee data")
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Employee created successfully"})
}

// Function to generate a unique employee ID (for example, increment the highest existing ID)
func generateUniqueEmployeeID(employees []models.Employee) int {
	highestID := 0
	for _, employee := range employees {
		if employee.ID > highestID {
			highestID = employee.ID
		}
	}
	// Increment the highest existing ID to generate a new unique ID
	return highestID + 1
}




/* UpdateEmployeeHandler updates an existing employee by ID.
// @Summary Update an existing employee
// @Description Update an existing employee by ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body model.Employee true "Updated employee object"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees/{id} [put]

*/

func UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid employee ID")
		return
	}

	// var employee models.Employee
	var updatedEmployee models.Employee
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedEmployee); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Read existing employees from CSV
	employees, err := utils.ReadCSV()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to read employee data")
		return
	}

	// Find and update the specific employee in the 'employees' slice
	employeeFound := false
	for i, emp := range employees {
		if emp.ID == id {
			// employees[i] = employee
			// Preserve the original ID
			updatedEmployee.ID = id
			employees[i] = updatedEmployee
			employeeFound = true
			break
		}
	}

	if !employeeFound {
		respondWithError(w, http.StatusNotFound, "Employee not found")
		return
	}

	// Write the updated employees to CSV
	err = utils.WriteToCSV(employees)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update employee data")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Employee updated successfully"})
}


/* DeleteEmployeeHandler deletes an employee by ID.
// @Summary Delete an employee
// @Description Delete an employee by ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees/{id} [delete]


*/

func DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid employee ID")
		return
	}

	// Read existing employees from CSV
	employees, err := utils.ReadCSV()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to read employee data")
		return
	}

	// Find and remove the specific employee from the 'employees' slice
	employeeIndex := -1
	for i, emp := range employees {
		if emp.ID == id {
			employeeIndex = i
			break
		}
	}

	if employeeIndex == -1 {
		respondWithError(w, http.StatusNotFound, "Employee not found")
		return
	}

	// Remove the employee from the slice
	employees = append(employees[:employeeIndex], employees[employeeIndex+1:]...)

	// Write the updated employees to CSV
	err = utils.WriteToCSV(employees)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update employee data")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Employee deleted successfully"})
}



func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}