package internal

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/application"
)

func LetsTry() {
	// Initialize the Context
	ctx := context.Background()

	// Start the application
	application.NewApplication(ctx).Start()
}
