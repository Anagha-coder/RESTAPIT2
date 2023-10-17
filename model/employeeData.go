package models

type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Deleted   bool   `json:"deleted"`
}


type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}