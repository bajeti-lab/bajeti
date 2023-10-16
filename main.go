package main

import (
	"github.com/bajeti-lab/bajeti/pkg/utils"
)

func main() {
	// create the database
	db, err := utils.CreateDatabase()

	if err != nil {
		panic("failed to connect database")
	}

	// create GUI for the user to interact with using fyne

}
