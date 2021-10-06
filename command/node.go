package command

import "errors"

type CreateNode struct {
	BaseCommand
	Name string `json:"name"`
	Type string `json:"type"`
}

func (cmd *CreateNode) Valid() error {
	if cmd.Action == "" || cmd.Name == "" || cmd.Type == "" {
		return errors.New("command is not valid")
	}
	return nil
}

type ActivateNode struct {
	BaseCommand
	Node string `json:"node"`
}

func (cmd *ActivateNode) Valid() error {
	if cmd.Action == "" || cmd.Node == "" {
		return errors.New("command is not valid")
	}
	return nil
}

type DeactivateNode struct {
	BaseCommand
	Node string `json:"node"`
}

func (cmd *DeactivateNode) Valid() error {
	if cmd.Action == "" || cmd.Node == "" {
		return errors.New("command is not valid")
	}
	return nil
}

type AddChild struct {
	BaseCommand
	Parent   string `json:"parent"`
	Child    string `json:"child"`
	Pipeline string `json:"pipeline"`
}

func (cmd *AddChild) Valid() error {
	if cmd.Action == "" || cmd.Parent == "" || cmd.Child == "" || cmd.Pipeline == "" {
		return errors.New("command is not valid")
	}
	return nil
}
