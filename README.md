# Employee Management System Using REST API
Employee Management System is a RESTful API service that allows you to perform various operations on employee data. It provides endpoints for creating, updating, deleting, and searching employees. The system is designed to manage employee records efficiently.

# Directory Structure

Employee-Management-System
├── Controller
│   └── routes.go
├── Docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── Handlers
│   └── handlers.go
├── Model
│   └── employeeData.go
├── Utils
│   ├── authenticate.go
│   ├── csv_utils.go
│   ├── employee.csv
│   └── logger.go
├── main.go
├── .gitignore
├── CODE_OF_CONDUCT.md
├── CONTRIBUTING.md
└── LICENSE


# Setup
1. Clone the repository:
bash
git clone <repository-url>

Navigate to the project directory:
bash
cd <project-folder>

Install dependencies:
bash
go mod tidy

Run the application:
bash
go run main.go













