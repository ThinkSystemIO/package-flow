package command

import "errors"

type CreatePipeline struct {
	BaseCommand
	Name string `json:"name"`
	Type string `json:"type"`
}

func (cmd *CreatePipeline) Valid() error {
	if cmd.Action == "" || cmd.Name == "" || cmd.Type == "" {
		return errors.New("command is not valid")
	}
	return nil
}
