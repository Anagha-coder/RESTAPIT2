# Employee Management System Using REST API
Employee Management System is a RESTful API service that allows you to perform various operations on employee data. It provides endpoints for creating, updating, deleting, and searching employees. The system is designed to manage employee records efficiently.


# Prerequisites
1. Go programming language installed.
2. Required dependencies installed (e.g., Gorilla Mux, CSV libraries).
3. Git (for cloning the repository).

# Swagger/OpenAPI Documentation
The APIs are documented using Swagger/OpenAPI specification. 
https://swagger.io/specification/
The OpenAPI spec (swagger.json) is provided along with the Go code on GitHub. 
You can use tools like Swagger UI to visualize and interact with the APIs.

# Code structure
📂 EmpManegmentSystem
    📂 docs
      📄 docs.go
      📄 swagger.json
      📄 swagger.yaml
    📂 handlers
      📄 handlers.go
    📂 controller
      📄 routes.go
    📂 models
      📄 employeeData.go
    📂 utils
      📄 Authenticate.go
      📄 csv_utils.go
      📄 employee.csv
      📄 jwt.go
      📄 logger.go
  📄 main.go
  📄 go.mod
  📄 go.sum
  📄 application.log


# Getting Started
1. Clone the repository:
git clone git@github.com:Anagha-coder/RESTAPIT2.git

2. cd employeemanagementsystem
Install dependencies:
go get -d ./...

3. Run the application:
go run main.go


# List of Operations Implemented as REST APIs

1. View/Search Employee Details

  a. View All Employee Details
  Endpoint: /employees
  Method: GET
  Description: Retrieve a list of all employees from the database.
  Response: Array of employee objects.

  b. View Details of a Specific Employee by ID
  Endpoint: /employees/{id}
  Method: GET
  Description: Retrieve details of a specific employee by their ID.
  Request Parameter: id (integer) - Employee ID
  Response: Employee object if found, else 404 Not Found.

  c. Search an Employee by First Name, Last Name, Email, or Role
  Endpoint: /employees/search
  Method: GET
  Description: Search for employees based on first name, last name, email, or role.

    Query Parameters:
    field (string) - Field to search (e.g., "FirstName", "LastName", "Email", "Role")
    value (string) - Value to search
    Response: Array of employee objects matching the search criteria.


2. Create Employees

  a. Create Employee
  Endpoint: /employees
  Method: POST
  Description: Create a new employee.
  Request Body: Employee object to be created.
  Response: Status 201 Created if successful, else 400 Bad Request if invalid data provided.

3. Update Employee Details

  a. Update Employee by ID
  Endpoint: /employees/{id}
  Method: PUT
  Description: Update details of an existing employee by their ID.
  Request Parameter: id (integer) - Employee ID
  Request Body: Updated employee object.
  Response: Status 200 OK if successful, else 400 Bad Request if invalid data provided, 404 Not Found if employee not found.

4. Delete Employee Details

  a. Delete Employee by ID
  Endpoint: /employees/{id}
  Method: DELETE
  Description: Delete an employee by their ID.
  Request Parameter: id (integer) - Employee ID
  Response: Status 200 OK if successful, else 400 Bad Request if invalid ID provided, 404 Not Found if employee not found.








