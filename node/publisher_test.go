package node

// func TestPublisher(t *testing.T) {
// 	JSON := []byte(`{
// 		"id":"id",
// 		"name":"name",
// 		"action":"activate",
// 		"payload":"test message"
// 	}`)
// 	command, _ := CommandFromJSON(JSON)

// 	mux := http.NewServeMux()
// 	node := NewPublisherNode(command)

// 	mux.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request) {
// 		node.AddSubscriber(w, r)
// 	})

// 	mux.HandleFunc("/receive", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(len(node.subscribers))
// 		config := []byte("{\"payload\":\"test\"}")
// 		context, _ := CommandFromJSON(config)
// 		node.Receive(context)
// 	})

// 	http.ListenAndServe(":8081", mux)
// }
