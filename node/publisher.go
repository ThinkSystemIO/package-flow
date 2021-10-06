package node

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/thinksystemio/package-flow/command"
	"github.com/thinksystemio/package-flow/pipeline"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nhooyr.io/websocket"
)

type Publisher struct {
	BaseNode
	Subscribers map[*websocket.Conn]struct{}
	Mu          sync.Mutex
}

//
// Publisher Base
//

func NewPublisherNode(cmd *command.CreateNode) *Publisher {
	node := &Publisher{Subscribers: map[*websocket.Conn]struct{}{}}
	node.ID = primitive.NewObjectID().Hex()
	node.Name = cmd.Name
	node.Active = true
	node.Type = "publisher"
	node.Children = map[Node]pipeline.Pipeline{}
	return node
}

func (node *Publisher) Receive(cmd command.Command) {
	log.Printf("publisher receive - %v", cmd.GetData())
	if !node.GetActive() {
		return
	}

	JSON, err := json.Marshal(cmd.GetData())
	if err != nil {
		return
	}

	for subscriber := range node.Subscribers {
		log.Printf("publisher send - %v", cmd.GetData())
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		subscriber.Write(ctx, websocket.MessageText, JSON)
		log.Printf("publisher sent - %v", cmd.GetData())
	}

	node.Send(cmd)
}

//
// Publisher Command API
//

func (node *Publisher) AddSubscriber(cmd *command.AddSubscriber) error {
	node.Mu.Lock()
	defer node.Mu.Unlock()

	subscriber, err := websocket.Accept(cmd.W, cmd.R, nil)
	if err != nil {
		return err
	}
	defer node.RemoveSubscriber(subscriber)
	ctx := subscriber.CloseRead(cmd.R.Context())

	node.Subscribers[subscriber] = struct{}{}

	<-ctx.Done()
	return nil
}

//
// Publisher Utils
//

func (node *Publisher) RemoveSubscriber(subscriber *websocket.Conn) {
	node.Mu.Lock()
	defer node.Mu.Unlock()

	if subscriber == nil {
		return
	}

	defer subscriber.Close(websocket.StatusNormalClosure, "closed")
	delete(node.Subscribers, subscriber)
}

func (node *Publisher) ToJSONStruct() map[string]interface{} {
	m := map[string]interface{}{}
	return m
}
