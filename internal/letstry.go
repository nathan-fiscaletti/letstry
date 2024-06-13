package internal

import "github.com/nathan-fiscaletti/letstry/internal/application"

func LetsTry() {
	// This function is called when the package is imported
	app := application.NewApplication()
	app.Start()
}
