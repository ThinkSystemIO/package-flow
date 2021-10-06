package command

import (
	"errors"
	"net/http"
)

type AddSubscriber struct {
	BaseCommand
	Node string `json:"node"`
	W    http.ResponseWriter
	R    *http.Request
}

func (cmd *AddSubscriber) Valid() error {
	if cmd.Action == "" || cmd.Node == "" || cmd.W == nil || cmd.R == nil {
		return errors.New("command is not valid")
	}
	return nil
}
