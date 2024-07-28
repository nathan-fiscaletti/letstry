package commands

import (
	"context"
	"errors"

	"github.com/nathan-fiscaletti/letstry/internal/manager"
)

var (
	ErrMissingTemplateName = errors.New("missing template name")
)

func DeleteTemplateCommand() Command {
	return Command{
		Name:             CommandDeleteTemplate,
		ShortDescription: "Delete a template",
		Description:      "Delete a template by name.",
		Arguments: []Argument{
			{
				Name:        "name",
				Description: "The name of the template to delete.",
				Required:    true,
			},
		},
		Executor: func(ctx context.Context, args []string) error {
			if len(args) < 1 {
				return ErrMissingTemplateName
			}

			templateName := args[0]

			mgr, err := manager.GetManager(ctx)
			if err != nil {
				return err
			}

			template, err := mgr.GetTemplate(ctx, templateName)
			if err != nil {
				return err
			}

			err = mgr.DeleteTemplate(ctx, template)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
