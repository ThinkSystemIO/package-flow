package node

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/thinksystemio/package-flow/command"
	"github.com/thinksystemio/package-flow/pipeline"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseNode struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Active    bool   `json:"activate"`
	Type      string `json:"type"`
	Children  map[Node]pipeline.Pipeline
	SyncMutex sync.Mutex
}

func NewBaseNode(command *command.CreateNode) *BaseNode {
	return &BaseNode{
		ID:        primitive.NewObjectID().Hex(),
		Name:      command.Name,
		Active:    true,
		Children:  map[Node]pipeline.Pipeline{},
		SyncMutex: sync.Mutex{},
	}
}

func (node *BaseNode) GetID() string {
	return node.ID
}

func (node *BaseNode) GetName() string {
	return node.Name
}

func (node *BaseNode) GetActive() bool {
	return node.Active
}

func (node *BaseNode) SetActive(active bool) {
	node.Active = active
}

func (node *BaseNode) GetType() string {
	return node.Type
}

func (node *BaseNode) GetChildren() map[Node]pipeline.Pipeline {
	return node.Children
}

func (node *BaseNode) AddChild(child Node) {
	node.SyncMutex.Lock()
	node.Children[child] = nil
	node.SyncMutex.Unlock()
}

func (node *BaseNode) RemoveChild(child Node) {
	node.SyncMutex.Lock()
	delete(node.Children, child)
	node.SyncMutex.Unlock()
	node.RemovePipeline(child)
}

func (node *BaseNode) HasChild(child Node) bool {
	_, exists := node.GetChildren()[child]
	return exists
}

// Addpipeline adds a pipeline from the parent to the child node.
// This can be done concurrently.
func (node *BaseNode) AddPipeline(child Node, pipeline pipeline.Pipeline) {
	node.SyncMutex.Lock()
	node.Children[child] = pipeline
	node.SyncMutex.Unlock()
}

// RemovePipeline removes a pipeline from parent to the child node.
// This can be done concurrently.
func (node *BaseNode) RemovePipeline(child Node) {
	node.SyncMutex.Lock()
	delete(node.Children, child)
	node.SyncMutex.Unlock()
}

func (node *BaseNode) Receive(cmd command.Command) {
	log.Printf("base receive - %v", cmd.GetData())
	if !node.GetActive() {
		return
	}

	node.Send(cmd)
}

func (node *BaseNode) Send(cmd command.Command) {
	log.Printf("base send - %v", cmd.GetData())
	if !node.GetActive() {
		return
	}

	if cmd.HasErrors() {
		return
	}

	for child := range node.Children {
		log.Printf("base send before - %v", cmd.GetData())
		pipeline := node.Children[child]
		if pipeline != nil {
			pipeline.Apply(cmd)
		}
		log.Printf("base send after - %v", cmd.GetData())
		child.Receive(cmd)
	}
}

func (node *BaseNode) Activate(cmd command.Command) {
	node.SetActive(true)
}

func (node *BaseNode) Deactivate(cmd command.Command) {
	node.SetActive(false)
}

func (node *BaseNode) ToJSON() ([]byte, error) {
	children := map[string]string{}
	for key, child := range node.Children {
		pipelineID := node.Children[key].GetID()
		children[child.GetID()] = pipelineID
	}

	s := &struct {
		ID       string            `json:"id"`
		Name     string            `json:"name"`
		Active   bool              `json:"active"`
		Type     string            `json:"type"`
		Children map[string]string `json:"children"`
	}{
		node.ID,
		node.Name,
		node.Active,
		node.Type,
		children,
	}

	return json.Marshal(s)
}

//
// Base Utils
//

func (node *BaseNode) ToJSONStruct() map[string]interface{} {
	m := map[string]interface{}{}
	return m
}
