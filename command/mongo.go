package command

import "errors"

type ConnectMongo struct {
	BaseCommand
	Node string `json:"node"`
	URL  string `json:"url"`
}

func (cmd *ConnectMongo) Valid() error {
	if cmd.Action == "" || cmd.Node == "" || cmd.URL == "" {
		return errors.New("command is not valid")
	}

	return nil
}

type AddMongo struct {
	BaseCommand
	Node     string                 `json:"node"`
	Document map[string]interface{} `json:"document"`
}

func (cmd *AddMongo) Valid() error {
	if cmd.Action == "" || cmd.Node == "" || cmd.Document == nil {
		return errors.New("command is not valid")
	}

	return nil
}

type UpdateMongo struct {
	BaseCommand
	Node     string                 `json:"node"`
	Document map[string]interface{} `json:"document"`
	Filter   map[string]interface{} `json:"filter"`
}

func (cmd *UpdateMongo) Valid() error {
	if cmd.Action == "" || cmd.Node == "" || cmd.Document == nil || cmd.Filter == nil {
		return errors.New("command is not valid")
	}

	return nil
}

type UpdateByIDMongo struct {
	BaseCommand
	Node     string                 `json:"node"`
	ID       string                 `json:"id"`
	Document map[string]interface{} `json:"document"`
}

func (cmd *UpdateByIDMongo) Valid() error {
	if cmd.Action == "" || cmd.Node == "" || cmd.ID == "" || cmd.Document == nil {
		return errors.New("command is not valid")
	}

	return nil
}

type RemoveMongo struct {
	BaseCommand
	Node   string                 `json:"node"`
	Filter map[string]interface{} `json:"filter"`
}

func (cmd *RemoveMongo) Valid() error {
	if cmd.Action == "" || cmd.Node == "" || cmd.Filter == nil {
		return errors.New("command is not valid")
	}

	return nil
}

type RemoveByIDMongo struct {
	BaseCommand
	Node string `json:"node"`
	ID   string `json:"id"`
}

func (cmd *RemoveByIDMongo) Valid() error {
	if cmd.Action == "" || cmd.Node == "" || cmd.ID == "" {
		return errors.New("command is not valid")
	}

	return nil
}

type QueryAllMongo struct {
	BaseCommand
	Node string `json:"node"`
}

func (cmd *QueryAllMongo) Valid() error {
	if cmd.Action == "" || cmd.Node == "" {
		return errors.New("command is not valid")
	}

	return nil
}
