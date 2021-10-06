package node

// func TestSubscriber(t *testing.T) {
// 	JSON := []byte(`{
// 		"id":"id",
// 		"name":"name",
// 		"action":"activate",
// 		"payload":"http://localhost:8080/subscribe"
// 	}`)
// 	command, _ := CommandFromJSON(JSON)

// 	node := NewSubscriberNode(command)
// 	child := &BaseNode{}
// 	node.AddChild(child)

// 	go node.Reducer(command)
// 	time.Sleep(1 * time.Second)
// 	if node.active != true {
// 		t.Error("node.activate should be true")
// 	}

// 	time.Sleep(30 * time.Second)
// 	node.Deactivate()
// 	if node.active != false {
// 		t.Error("node.activate should be false")
// 	}
// }
