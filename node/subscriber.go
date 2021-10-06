package node

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/thinksystemio/package-flow/command"
	"github.com/thinksystemio/package-flow/pipeline"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nhooyr.io/websocket"
)

type Subscriber struct {
	BaseNode
	URL      string          `json:"url"`
	WSActive bool            `json:"wsactive"`
	Signal   chan bool       `json:"-"`
	Client   *websocket.Conn `json:"-"`
}

//
// Subscriber Base
//

func NewSubscriberNode(cmd *command.CreateNode) *Subscriber {
	node := &Subscriber{}
	node.ID = primitive.NewObjectID().Hex()
	node.Name = cmd.Name
	node.Active = true
	node.Type = "subscriber"
	node.Children = map[Node]pipeline.Pipeline{}
	return node
}

//
// Subscriber Command API
//

func (node *Subscriber) ActivateWS(cmd *command.ActivateWS) {
	node.Close()

	node.WSActive = true
	node.Signal = make(chan bool)

	go func() {
		for node.WSActive {
			go node.Listen(cmd)
			<-node.Signal
			time.Sleep(5 * time.Second)
		}
	}()
}

func (node *Subscriber) DeactivateWS(cmd *command.DeactivateWS) {
	node.WSActive = false
	node.Close()
}

func (node *Subscriber) UpdateURL(cmd *command.UpdateURL) {
	node.URL = cmd.URL
}

//
// Subscriber Utils
//

func (node *Subscriber) Close() {
	if node.Client != nil {
		node.Client.Close(websocket.StatusNormalClosure, "closing current connection")
	}
}

func (node *Subscriber) Listen(cmd command.Command) {
	defer func() { node.Signal <- true }()

	ctx := context.Background()
	client, _, err := websocket.Dial(ctx, node.URL, nil)
	if err != nil {
		cmd.AppendError(err)
		return
	}

	node.Client = client
	defer node.Close()

	for node.WSActive {
		_, msg, err := client.Read(ctx)
		if err != nil {
			cmd.AppendError(err)
			return
		}

		data := map[string]interface{}{}
		if err := json.Unmarshal(msg, &data); err != nil {
			cmd.AppendError(err)
			return
		}
		log.Println("original data")
		log.Println(data)

		copied := &command.UpdateURL{URL: node.URL}
		copied.Action = cmd.GetAction()

		// copy errors
		copied.Errors = make([]command.Error, len(cmd.GetErrors()))
		copy(copied.Errors, cmd.GetErrors())

		// copy data
		copied.Data = DeepCopyMap(data)

		log.Println("value data")
		log.Println(copied.Data)

		node.Send(copied)
	}
}

func (node *Subscriber) ToJSONStruct() map[string]interface{} {
	m := map[string]interface{}{}
	m["url"] = node.URL
	m["wsactive"] = node.WSActive
	return m
}
