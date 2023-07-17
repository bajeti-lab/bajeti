package main

import (
	// "errors"
	"fmt"
	// "net/http"
	// "os"
	// "time"

	//import my project packages
	// "github.com/bajeti-lab/bajeti/internal/handlers"
	// "github.com/bajeti-lab/bajeti/internal/models"
	"github.com/bajeti-lab/bajeti/internal/budget"
)

func main() {
	// take user input
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter Email: ")
	var email string
	fmt.Scanln(&email)

	fmt.Print("Enter Password: ")
	var password string
	fmt.Scanln(&password)

	var BudgetAmount float64
	fmt.Print("Enter Budget Amount: ")
	fmt.Scanln(&BudgetAmount)

	// create a new user using the createuser function
	user, err := budget.CreateUser(name, email, password)

	



	// print the user details
	fmt.Println("User Details: ", user, err)


	// // create a new router
	// router := http.NewServeMux()

	// // register the handler functions for different routes
	// router.HandleFunc("/", handlers.HomeHandler)
	// router.HandleFunc("/budgets", handlers.BudgetsHandler)
	// router.HandleFunc("/transactions", handlers.TransactionsHandler)

	// // create a new server
	// server := http.Server{
	// 	Addr:    fmt.Sprintf(":%d", 8080),
	// 	Handler: router,
	// }

	// // start the server
	// if err := server.ListenAndServe(); err != nil {
	// 	if !errors.Is(err, http.ErrServerClosed) {
	// 		fmt.Println("Error starting server: ", err)
	// 	}
	// }

	// // wait for 10 seconds
	// time.Sleep(10 * time.Second)

	// // send a request to the server
	// requestURL := fmt.Sprintf("http://localhost:%d", 8080)
	// request, err := http.Get(requestURL)

	// // check for errors
	// if err != nil {
	// 	fmt.Println("Error sending request: ", err)
	// 	os.Exit(1)
	// }

	// // print the response
	// fmt.Println("Response: ", request.Status)

}
