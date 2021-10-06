package flow

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/thinksystemio/package-flow/command"
	"github.com/thinksystemio/package-flow/tree"
)

type TestingDoc = map[string]interface{}

func CreateNode(name string, nodeType string) []byte {
	data := TestingDoc{
		"action": command.CREATE_NODE,
		"name":   name,
		"type":   nodeType,
	}
	JSON, _ := json.Marshal(data)
	return JSON
}

func AddChild(parent string, child string, pipeline string) []byte {
	data := TestingDoc{
		"action":   command.ADD_CHILD,
		"parent":   parent,
		"child":    child,
		"pipeline": pipeline,
	}
	JSON, _ := json.Marshal(data)
	return JSON
}

func CreatePipeline(name string, pipelineType string) []byte {
	data := TestingDoc{
		"action": command.CREATE_PIPELINE,
		"name":   name,
		"type":   pipelineType,
	}
	JSON, _ := json.Marshal(data)
	return JSON
}

func UpdateFilterPipeline(name string, filter map[string]struct{}) []byte {
	data := TestingDoc{
		"action": command.UPDATE_FILTER_PIPELINE,
		"name":   name,
		"filter": filter,
	}
	JSON, _ := json.Marshal(data)
	return JSON
}

func ActivateNode(node string, pipelineType string) []byte {
	data := TestingDoc{
		"action": command.ACTIVATE_NODE,
		"node":   node,
	}
	JSON, _ := json.Marshal(data)
	return JSON
}

func DeactivateNode(node string, pipelineType string) []byte {
	data := TestingDoc{
		"action": command.DEACTIVATE_NODE,
		"node":   node,
	}
	JSON, _ := json.Marshal(data)
	return JSON
}

func UpdateURL(node string, url string) []byte {
	data := TestingDoc{
		"action": command.UPDATE_URL,
		"node":   node,
		"url":    url,
	}
	JSON, _ := json.Marshal(data)
	return JSON
}

func ConnectMongo(node string, url string) []byte {
	data := TestingDoc{
		"action": command.CONNECT_MONGO,
		"node":   node,
		"url":    url,
	}
	JSON, _ := json.Marshal(data)
	return JSON
}

func AddMongo(node string, document map[string]interface{}) []byte {
	data := TestingDoc{
		"action":   command.ADD_MONGO,
		"node":     node,
		"document": document,
	}
	JSON, _ := json.Marshal(data)
	return JSON
}

func UpdateMongo(node string, document TestingDoc, filter TestingDoc) []byte {
	data := TestingDoc{
		"action":   command.UPDATE_MONGO,
		"node":     node,
		"document": document,
		"filter":   filter,
	}
	JSON, _ := json.Marshal(data)
	return JSON
}

func UpdateByIDMongo(node string, id string, document TestingDoc) []byte {
	data := TestingDoc{
		"action":   command.UPDATE_BY_ID_MONGO,
		"node":     node,
		"id":       id,
		"document": document,
	}
	JSON, _ := json.Marshal(data)
	return JSON
}

func RemoveMongo(node string, filter TestingDoc) []byte {
	data := TestingDoc{
		"action": command.REMOVE_MONGO,
		"node":   node,
		"filter": filter,
	}
	JSON, _ := json.Marshal(data)
	return JSON
}

func RemoveByIDMongo(node string, id string) []byte {
	data := TestingDoc{
		"action": command.REMOVE_BY_ID_MONGO,
		"node":   node,
		"id":     id,
	}
	JSON, _ := json.Marshal(data)
	return JSON
}

func QueryAllMongo(node string) []byte {
	data := TestingDoc{
		"action": command.QUERY_ALL_MONGO,
		"node":   node,
	}
	JSON, _ := json.Marshal(data)
	return JSON
}

func TestTree(t *testing.T) {
	tree := tree.NewTree()

	fmt.Println(DispatchFromJSON(tree, CreateNode("node_mongo", "mongo")))
	fmt.Println(DispatchFromJSON(tree, CreateNode("node_base", "base")))
	fmt.Println(DispatchFromJSON(tree, CreateNode("node_publisher", "publisher")))

	fmt.Println(DispatchFromJSON(tree, CreatePipeline("pipe_base", "base")))
	fmt.Println(DispatchFromJSON(tree, CreatePipeline("pipe_filter", "filter")))
	fmt.Println(DispatchFromJSON(tree, UpdateFilterPipeline("pipe_filter", map[string]struct{}{"1": {}})))

	fmt.Println(DispatchFromJSON(tree, AddChild("node_mongo", "node_base", "pipe_base")))
	fmt.Println(DispatchFromJSON(tree, AddChild("node_mongo", "node_publisher", "pipe_filter")))

	fmt.Println(DispatchFromJSON(tree, ConnectMongo("node_mongo", "localhost")))
	fmt.Println(DispatchFromJSON(tree, AddMongo("node_mongo", TestingDoc{"1": "1", "2": "2"})))
	fmt.Println(DispatchFromJSON(tree, QueryAllMongo("node_mongo")))
	fmt.Println(DispatchFromJSON(tree, RemoveMongo("node_mongo", TestingDoc{"1": "1"})))

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			name := strings.Split(r.URL.Path, "/")[1]
			cmd := &command.AddSubscriber{Node: name, W: w, R: r}
			fmt.Println(Dispatch(tree, cmd))
		case "POST":
			fmt.Println(DispatchFromJSON(tree, AddMongo("node_mongo", TestingDoc{"1": "1", "2": "2"})))
			fmt.Println(DispatchFromJSON(tree, RemoveMongo("node_mongo", TestingDoc{"1": "1"})))
		}
	})

	http.ListenAndServe(":8082", mux)
}
