package pipeline

import (
	"errors"
	"strings"

	"github.com/thinksystemio/package-flow/command"
)

type Pipeline interface {
	GetID() string
	GetName() string
	GetType() string

	Apply(command.Command)

	ToJSON() ([]byte, error)
}

func NewPipeline(command *command.CreatePipeline) Pipeline {
	pipelineType := strings.ToLower(command.Type)

	switch pipelineType {
	case "base":
		p := NewBasePipeline(command)
		return p
	case "filter":
		p := NewFilterPipeline(command)
		return p
	default:
		err := errors.New("Pipeline.New - invalid pipeline type")
		command.AppendError(err)
		return nil
	}
}
