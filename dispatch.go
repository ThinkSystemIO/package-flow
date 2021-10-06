package flow

import (
	"encoding/json"
	"errors"

	"github.com/thinksystemio/package-flow/command"
	"github.com/thinksystemio/package-flow/node"
	"github.com/thinksystemio/package-flow/pipeline"
	"github.com/thinksystemio/package-flow/tree"
)

func NewTree() *tree.Tree {
	return tree.NewTree()
}

func IdentifyCommand(data []byte) *command.BaseCommand {
	command := &command.BaseCommand{}
	err := json.Unmarshal(data, command)
	if err != nil {
		command.AppendError(err)
	}
	return command
}

func DispatchFromJSON(tree *tree.Tree, data []byte, options ...interface{}) command.Command {
	base := IdentifyCommand(data)
	if base.HasErrors() {
		return base
	}

	dispensed := command.Dispense(base.GetAction(), data, options...)
	if dispensed.HasErrors() {
		return dispensed
	}

	return Dispatch(tree, dispensed, options...)
}

func Dispatch(tree *tree.Tree, cmd command.Command, options ...interface{}) command.Command {
	switch cmd := cmd.(type) {

	//
	// Tree
	//

	case *command.CreateNode:
		if found, _ := tree.GetNodeByNameOrID(cmd.Name); found != nil {
			cmd.AppendError(errors.New("node already exists"))
			return cmd
		}

		n := node.NewNode(cmd)
		if n != nil {
			if err := tree.AddNode(n); err != nil {
				cmd.AppendError(err)
				return cmd
			}
		}

		cmd.Data = n.GetID()
	case *command.AddChild:
		parent, err := tree.GetNodeByNameOrID(cmd.Parent)
		if err != nil {
			cmd.AppendError(err)
			return cmd
		}

		child, err := tree.GetNodeByNameOrID(cmd.Child)
		if err != nil {
			cmd.AppendError(err)
			return cmd
		}

		pipe := pipeline.NewBasePipelineDefault(cmd.Pipeline)
		if pipe != nil {
			if err := tree.AddPipeline(pipe); err != nil {
				cmd.AppendError(err)
				return cmd
			}
		}

		parent.AddPipeline(child, pipe)
	case *command.CreatePipeline:
		if found, _ := tree.GetPipelineByNameOrID(cmd.Name); found != nil {
			cmd.AppendError(errors.New("pipeline already exists"))
			return cmd
		}

		p := pipeline.NewPipeline(cmd)
		if p != nil {
			if err := tree.AddPipeline(p); err != nil {
				cmd.AppendError(err)
				return cmd
			}
		}

	//
	// Filter Pipeline
	//

	case *command.UpdateFilterPipeline:
		p, err := tree.GetPipelineByNameOrID(cmd.Name)
		if err != nil {
			cmd.AppendError(err)
			return cmd
		}
		if filter, ok := p.(*pipeline.FilterPipeline); ok {
			filter.UpdatePipelineFilter(cmd)
		}

	//
	// Node
	//

	case *command.ActivateNode:
		n, err := tree.GetNodeByNameOrID(cmd.Node)
		if err != nil {
			cmd.AppendError(err)
			return cmd
		}
		n.Activate(cmd)
	case *command.DeactivateNode:
		n, err := tree.GetNodeByNameOrID(cmd.Node)
		if err != nil {
			cmd.AppendError(err)
			return cmd
		}
		n.Deactivate(cmd)

	//
	// Publisher
	//

	case *command.AddSubscriber:
		n, err := tree.GetNodeByNameOrID(cmd.Node)
		if err != nil {
			cmd.AppendError(err)
			return cmd
		}
		if publisher, ok := n.(*node.Publisher); ok {
			publisher.AddSubscriber(cmd)
		}

	//
	// Subscriber
	//

	case *command.UpdateURL:
		n, err := tree.GetNodeByNameOrID(cmd.Node)
		if err != nil {
			cmd.AppendError(err)
			return cmd
		}
		if subscriber, ok := n.(*node.Subscriber); ok {
			subscriber.UpdateURL(cmd)
		}

	case *command.ActivateWS:
		n, err := tree.GetNodeByNameOrID(cmd.Node)
		if err != nil {
			cmd.AppendError(err)
			return cmd
		}
		if subscriber, ok := n.(*node.Subscriber); ok {
			subscriber.ActivateWS(cmd)
		}
	case *command.DeactivateWS:
		n, err := tree.GetNodeByNameOrID(cmd.Node)
		if err != nil {
			cmd.AppendError(err)
			return cmd
		}
		if subscriber, ok := n.(*node.Subscriber); ok {
			subscriber.DeactivateWS(cmd)
		}

	//
	// Mongo
	//

	case *command.ConnectMongo:
		n, err := tree.GetNodeByNameOrID(cmd.Node)
		if err != nil {
			cmd.AppendError(err)
			return cmd
		}
		if mongo, ok := n.(*node.Mongo); ok {
			mongo.Connect(cmd)
		}
	case *command.AddMongo:
		n, err := tree.GetNodeByNameOrID(cmd.Node)
		if err != nil {
			cmd.AppendError(err)
			return cmd
		}
		if mongo, ok := n.(*node.Mongo); ok {
			mongo.Add(cmd)
		}
	case *command.UpdateMongo:
		n, err := tree.GetNodeByNameOrID(cmd.Node)
		if err != nil {
			cmd.AppendError(err)
			return cmd
		}
		if mongo, ok := n.(*node.Mongo); ok {
			mongo.Update(cmd)
		}
	case *command.UpdateByIDMongo:
		n, err := tree.GetNodeByNameOrID(cmd.Node)
		if err != nil {
			cmd.AppendError(err)
			return cmd
		}
		if mongo, ok := n.(*node.Mongo); ok {
			mongo.UpdateByID(cmd)
		}
	case *command.RemoveMongo:
		n, err := tree.GetNodeByNameOrID(cmd.Node)
		if err != nil {
			cmd.AppendError(err)
			return cmd
		}
		if mongo, ok := n.(*node.Mongo); ok {
			mongo.Remove(cmd)
		}
	case *command.RemoveByIDMongo:
		n, err := tree.GetNodeByNameOrID(cmd.Node)
		if err != nil {
			cmd.AppendError(err)
			return cmd
		}
		if mongo, ok := n.(*node.Mongo); ok {
			mongo.RemoveByID(cmd)
		}
	case *command.QueryAllMongo:
		n, err := tree.GetNodeByNameOrID(cmd.Node)
		if err != nil {
			cmd.AppendError(err)
			return cmd
		}
		if mongo, ok := n.(*node.Mongo); ok {
			mongo.QueryAll(cmd)
		}
	}

	return cmd
}
