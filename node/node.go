package node

import (
	"errors"
	"strings"

	"github.com/thinksystemio/package-flow/command"
	"github.com/thinksystemio/package-flow/pipeline"
)

// Node is an interface for all nodes in the flow tree.
type Node interface {
	GetID() string
	GetName() string
	GetActive() bool
	SetActive(bool)
	GetType() string

	Activate(command.Command)
	Deactivate(command.Command)

	Send(command.Command)
	Receive(command.Command)

	GetChildren() map[Node]pipeline.Pipeline
	AddChild(Node)
	RemoveChild(Node)
	HasChild(Node) bool

	AddPipeline(Node, pipeline.Pipeline)
	RemovePipeline(Node)

	ToJSON() ([]byte, error)
	ToJSONStruct() map[string]interface{}
}

func NewNode(command *command.CreateNode) Node {
	nodeType := strings.ToLower(command.Type)

	switch nodeType {
	case "base":
		return NewBaseNode(command)
	case "publisher":
		return NewPublisherNode(command)
	case "subscriber":
		return NewSubscriberNode(command)
	case "mongo":
		return NewMongoNode(command)
	default:
		err := errors.New("Node.New - invalid node type")
		command.AppendError(err)
	}

	return nil
}
