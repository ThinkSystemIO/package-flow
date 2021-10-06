package command

import "errors"

type UpdateFilterPipeline struct {
	BaseCommand
	Name   string              `json:"name"`
	Filter map[string]struct{} `json:"filter"`
}

func (cmd *UpdateFilterPipeline) Valid() error {
	if cmd.Action == "" || cmd.Name == "" || cmd.Filter == nil {
		return errors.New("command is not valid")
	}
	return nil
}
