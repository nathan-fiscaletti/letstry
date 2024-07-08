package internal

import (
	"context"

	"github.com/nathan-fiscaletti/letstry/internal/application"
)

func LetsTry() {
	application.NewApplication(
		context.Background(),
	).Start()
}
