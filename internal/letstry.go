package internal

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/application"
)

func LetsTry() {
	// Initialize the Context
	appContext := context.Background()

	// Initialize the Application
	app := application.NewApplication(appContext)

	// Start the Application
	app.Start()
}
