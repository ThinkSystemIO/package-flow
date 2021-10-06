package command

import "errors"

type BaseCommand struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
	Errors []Error     `json:"errors"`
}

func (cmd *BaseCommand) GetAction() string {
	return cmd.Action
}

func (cmd *BaseCommand) GetData() interface{} {
	if cmd.Data == nil {
		return map[string]interface{}{}
	}
	return cmd.Data
}

func (cmd *BaseCommand) SetData(data interface{}) {
	cmd.Data = data
}

func (cmd *BaseCommand) GetErrors() []Error {
	return cmd.Errors
}

func (cmd *BaseCommand) AppendError(err error) {
	if err != nil {
		cmd.Errors = append(cmd.Errors, Error{Message: err.Error()})
	}
}

func (cmd *BaseCommand) HasErrors() bool {
	return len(cmd.Errors) != 0
}

func (cmd *BaseCommand) Valid() error {
	if cmd.Action == "" {
		return errors.New("command is not valid")
	}
	return nil
}

type Error struct {
	Message string `json:"errors"`
}
