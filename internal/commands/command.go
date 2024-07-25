package commands

import (
	"context"
)

type Command func(context.Context, []string) error
