package node

import (
	"context"
	"time"

	"github.com/thinksystemio/package-flow/command"
	"github.com/thinksystemio/package-flow/pipeline"
	gomongo "github.com/thinksystemio/package-gomongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	BaseNode
	client *mongo.Client
}

//
// Mongo Base
//

func NewMongoNode(cmd *command.CreateNode) *Mongo {
	node := &Mongo{}
	node.ID = primitive.NewObjectID().Hex()
	node.Name = cmd.Name
	node.Active = true
	node.Type = "mongo"
	node.Children = map[Node]pipeline.Pipeline{}
	return node
}

func (node *Mongo) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := node.client.Disconnect(ctx)
	return err
}

//
// Mongo Command API
//

func (node *Mongo) Connect(cmd *command.ConnectMongo) {
	client, err := gomongo.CreateMongoClient(cmd.URL)
	if err != nil {
		cmd.AppendError(err)
		return
	}

	node.client = client
}

func (node *Mongo) Add(cmd *command.AddMongo) {
	collection := node.client.Database("data").Collection(node.Name)
	data, err := gomongo.Add(collection, cmd.Document)
	cmd.SetData(data)
	cmd.AppendError(err)
	node.Send(cmd)
}

func (node *Mongo) Update(cmd *command.UpdateMongo) {
	collection := node.client.Database("data").Collection(node.Name)
	data, err := gomongo.Update(collection, cmd.Filter, cmd.Document)
	cmd.SetData(data)
	cmd.AppendError(err)
	node.Send(cmd)
}

func (node *Mongo) UpdateByID(cmd *command.UpdateByIDMongo) {
	collection := node.client.Database("data").Collection(node.Name)
	data, err := gomongo.UpdateByID(collection, cmd.ID, cmd.Document)
	cmd.SetData(data)
	cmd.AppendError(err)
	node.Send(cmd)
}

func (node *Mongo) Remove(cmd *command.RemoveMongo) {
	collection := node.client.Database("data").Collection(node.Name)
	data, err := gomongo.Remove(collection, cmd.Filter)
	cmd.SetData(data)
	cmd.AppendError(err)
	node.Send(cmd)
}

func (node *Mongo) RemoveByID(cmd *command.RemoveByIDMongo) {
	collection := node.client.Database("data").Collection(node.Name)
	data, err := gomongo.RemoveByID(collection, cmd.ID)
	cmd.SetData(data)
	cmd.AppendError(err)
	node.Send(cmd)
}

func (node *Mongo) QueryAll(cmd *command.QueryAllMongo) {
	collection := node.client.Database("data").Collection(node.Name)
	data, err := gomongo.GetAll(collection)
	if err != nil {
		cmd.AppendError(err)
		return
	}
	cmd.SetData(data)
	node.Send(cmd)
}

//
// Mongo Utils
//

func (node *Mongo) ToJSONStruct() map[string]interface{} {
	m := map[string]interface{}{}
	return m
}
