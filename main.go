//go:generate swag init
package main

import (
	"empsystem/controller"
	"empsystem/utils"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "empsystem/docs"
	_ "empsystem/handlers"
	
	httpSwagger "github.com/swaggo/http-swagger"

	
)


//@title Employee Management System
//@version 1.0
//@description This is a sample server which perform REST operations.
//@host localhost:8080/employees



func main() {
    utils.InitLogger() 

	err := utils.InitCSV()
	if err != nil {
		// log.Fatal("Error initializing CSV file: ", err)
		utils.ErrorLogger.Fatalln("Error initializing CSV file:", err)
	}
	

	r := mux.NewRouter()

	controller.InitializeRoutes(r) // Initialize routes here

	

	http.Handle("/", r)
	
	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))


    r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(httpSwagger.URL("/docs/swagger.json")))

   
	utils.InfoLogger.Println("Head over to http://localhost:8080/employees")
    utils.InfoLogger.Println("Server is up and running...")

	// mt.Println("headover to 8080. Server is up and running...")
	// lofg.Fatal(http.ListenAndServe(":8080", nil))

}
