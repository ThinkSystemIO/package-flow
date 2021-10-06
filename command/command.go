package command

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

const (
	CREATE_NODE     = "create_node"
	ADD_CHILD       = "add_child"
	ACTIVATE_NODE   = "activate_node"
	DEACTIVATE_NODE = "deactivate_node"

	ADD_SUBSCRIBER = "add_subscriber"
	UPDATE_URL     = "update_url"
	ACTIVATE_WS    = "activate_ws"
	DECTIVATE_WS   = "deactivate_ws"

	CONNECT_MONGO      = "connect_mongo"
	ADD_MONGO          = "add_mongo"
	UPDATE_MONGO       = "update_mongo"
	UPDATE_BY_ID_MONGO = "update_by_id_mongo"
	REMOVE_MONGO       = "remove_mongo"
	REMOVE_BY_ID_MONGO = "remove_by_id_mongo"
	QUERY_ALL_MONGO    = "query_all_mongo"

	CREATE_PIPELINE        = "create_pipeline"
	UPDATE_FILTER_PIPELINE = "update_filter_pipeline"
)

type Command interface {
	GetData() interface{}
	SetData(interface{})
	GetAction() string
	GetErrors() []Error
	AppendError(error)
	HasErrors() bool
	Valid() error
}

func Dispense(action string, data []byte, options ...interface{}) Command {
	formatted := strings.ToLower(action)
	switch formatted {

	// Tree
	case CREATE_NODE:
		return Decode(&CreateNode{}, data)
	case ADD_CHILD:
		return Decode(&AddChild{}, data)
	case CREATE_PIPELINE:
		return Decode(&CreatePipeline{}, data)

	// Filter Pipeline
	case UPDATE_FILTER_PIPELINE:
		return Decode(&UpdateFilterPipeline{}, data)

	// Node
	case ACTIVATE_NODE:
		return Decode(&ActivateNode{}, data)
	case DEACTIVATE_NODE:
		return Decode(&DeactivateNode{}, data)

	// Publisher Node
	case ADD_SUBSCRIBER:
		return DecodeWithOptions(&AddSubscriber{}, data, options...)

	// Subscriber Node
	case UPDATE_URL:
		return Decode(&UpdateURL{}, data)
	case ACTIVATE_WS:
		return Decode(&ActivateWS{}, data)
	case DECTIVATE_WS:
		return Decode(&DeactivateWS{}, data)

	// MongoDB Node
	case CONNECT_MONGO:
		return Decode(&ConnectMongo{}, data)
	case ADD_MONGO:
		return Decode(&AddMongo{}, data)
	case UPDATE_MONGO:
		return Decode(&UpdateMongo{}, data)
	case UPDATE_BY_ID_MONGO:
		return Decode(&UpdateByIDMongo{}, data)
	case REMOVE_MONGO:
		return Decode(&RemoveMongo{}, data)
	case REMOVE_BY_ID_MONGO:
		return Decode(&RemoveByIDMongo{}, data)
	case QUERY_ALL_MONGO:
		return Decode(&QueryAllMongo{}, data)

	default:
		command := &BaseCommand{}
		command.AppendError(errors.New("invalid action"))
		return command
	}
}

func Decode(cmd Command, data []byte) Command {
	err := json.Unmarshal(data, cmd)
	cmd.AppendError(err)

	if err := cmd.Valid(); err != nil {
		cmd.AppendError(err)
	}

	return cmd
}

func DecodeWithOptions(cmd Command, data []byte, options ...interface{}) Command {
	err := json.Unmarshal(data, cmd)
	cmd.AppendError(err)

	if cmd, ok := cmd.(*AddSubscriber); ok {
		for _, arg := range options {
			if value, ok := arg.(http.ResponseWriter); ok {
				cmd.W = value
			}

			if value, ok := arg.(*http.Request); ok {
				cmd.R = value
			}
		}
	}

	if err := cmd.Valid(); err != nil {
		cmd.AppendError(err)
	}

	return cmd
}
