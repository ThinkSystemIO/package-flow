package tree

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/thinksystemio/package-flow/node"
	"github.com/thinksystemio/package-flow/pipeline"
)

// Tree is a flat structure that contains a map of nodes. The
// individual nodes are responsible for keeping track of their
// children.
type Tree struct {
	Nodes     map[string]node.Node
	Pipelines map[string]pipeline.Pipeline
	Mu        sync.Mutex
}

// NewTree creates a new instance of a tree.
func NewTree() *Tree {
	return &Tree{
		Nodes:     map[string]node.Node{},
		Pipelines: map[string]pipeline.Pipeline{},
	}
}

// IsEmpty checks if the tree contains zero nodes.
func (tree *Tree) IsEmpty() bool {
	return len(tree.Nodes) == 0
}

// IsNotEmpty checks if the tree contains at least one node.
func (tree *Tree) IsNotEmpty() bool {
	return len(tree.Nodes) != 0
}

// GetNodeByNameOrID returns a node if found. When searching by ID, the
// algorithm has O(1) time complexity. When searching by name, the algorithm
// has O(n) time complexity.
func (tree *Tree) GetNodeByNameOrID(nameOrID string) (node.Node, error) {
	if node, hasKey := tree.Nodes[nameOrID]; hasKey {
		return node, nil
	}

	for _, node := range tree.Nodes {
		if node.GetName() == nameOrID {
			return node, nil
		}
	}

	err := fmt.Errorf("node with name or ID of %s does not exist", nameOrID)
	return nil, err
}

// AddNode adds a node to the tree. This can be done
// concurrently.
func (tree *Tree) AddNode(node node.Node) error {
	if node.GetID() == "" || node.GetName() == "" {
		return errors.New("node must have ID and Name")
	}

	if n, _ := tree.GetNodeByNameOrID(node.GetName()); n != nil {
		return errors.New("node already exists")
	}

	tree.Mu.Lock()
	tree.Nodes[node.GetID()] = node
	tree.Mu.Unlock()

	return nil
}

// RemoveNode removes a node from the tree. This can
// be done concurrently.
func (tree *Tree) RemoveNode(node node.Node) {
	tree.Mu.Lock()
	delete(tree.Nodes, node.GetID())
	tree.Mu.Unlock()
}

// GetPipelineByNameOrID returns a pipeline if found. When searching by ID, the
// algorithm has O(1) time complexity. When searching by name, the algorithm
// has O(n) time complexity.
func (tree *Tree) GetPipelineByNameOrID(nameOrID string) (pipeline.Pipeline, error) {
	if pipeline, hasKey := tree.Pipelines[nameOrID]; hasKey {
		return pipeline, nil
	}

	for _, pipeline := range tree.Pipelines {
		if pipeline.GetName() == nameOrID {
			return pipeline, nil
		}
	}

	err := fmt.Errorf("pipeline with name or ID of %s does not exist", nameOrID)
	return nil, err
}

// AddPipeline adds a pipeline to the tree. This can be done
// concurrently.
func (tree *Tree) AddPipeline(pipeline pipeline.Pipeline) error {
	if pipeline.GetID() == "" || pipeline.GetName() == "" {
		return errors.New("pipeline must have ID and Name")
	}

	if pipe, _ := tree.GetPipelineByNameOrID(pipeline.GetID()); pipe != nil {
		return errors.New("pipeline already exists")
	}

	tree.Mu.Lock()
	tree.Pipelines[pipeline.GetID()] = pipeline
	tree.Mu.Unlock()

	return nil
}

// RemovePipeline removes a pipeline from the tree. This can
// be done concurrently.
func (tree *Tree) RemovePipeline(pipeline pipeline.Pipeline) {
	tree.Mu.Lock()
	delete(tree.Pipelines, pipeline.GetID())
	tree.Mu.Unlock()
}

// ToJSON converts the tree to sendable bytes. This
// is necessary as each node in the tree must also
// having its own ToJSON function that converts each
// child node of a particular node into an ID. This ID
// can then be used by the tree to lookup a particular node
// when sent to the client as JSON.
func (tree *Tree) ToJSON() ([]byte, error) {
	nodes := map[string]interface{}{}
	for key, n := range tree.Nodes {
		children := map[string]string{}
		for key, child := range n.GetChildren() {
			children[key.GetID()] = child.GetID()
		}

		s := &struct {
			ID       string                 `json:"id"`
			Name     string                 `json:"name"`
			Type     string                 `json:"type"`
			Active   bool                   `json:"active"`
			Children map[string]string      `json:"children"`
			Props    map[string]interface{} `json:"props"`
		}{
			n.GetID(),
			n.GetName(),
			n.GetType(),
			n.GetActive(),
			children,
			n.ToJSONStruct(),
		}

		nodes[key] = s
	}

	result := map[string]interface{}{
		"nodes":     nodes,
		"pipelines": tree.Pipelines,
	}

	return json.Marshal(result)
}
