package internal

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/application"
	"github.com/nathan-fiscaletti/letstry/internal/session_manager"
)

func LetsTry() {
	// Initialize the Context
	appContext := context.Background()
	appContext = session_manager.ContextWithSessionManager(appContext)

	// Initialize the Application
	app := application.NewApplication(appContext)

	// Start the Application
	app.Start()
}
