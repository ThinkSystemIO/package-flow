package command

import "errors"

type UpdateURL struct {
	BaseCommand
	Node string `json:"node"`
	URL  string `json:"url"`
}

func (cmd *UpdateURL) UpdateURL() error {
	if cmd.Action == "" || cmd.Node == "" || cmd.URL == "" {
		return errors.New("command is not valid")
	}
	return nil
}

type ActivateWS struct {
	BaseCommand
	Node string `json:"node"`
}

func (cmd *ActivateWS) ActivateWS() error {
	if cmd.Action == "" || cmd.Node == "" {
		return errors.New("command is not valid")
	}
	return nil
}

type DeactivateWS struct {
	BaseCommand
	Node string `json:"node"`
}

func (cmd *DeactivateWS) DeactivateWS() error {
	if cmd.Action == "" || cmd.Node == "" {
		return errors.New("command is not valid")
	}
	return nil
}
