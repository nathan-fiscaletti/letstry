package internal

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/application"
)

func LetsTry() {
	// Initialize the Context
	application.NewApplication(
		context.Background(),
	).Start()
}
