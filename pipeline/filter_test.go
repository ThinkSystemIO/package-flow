package pipeline

// func createTree() (Node, Node) {
// 	parent := CreateBaseNode("parent", "parent")
// 	child := CreateBaseNode("child", "child")

// 	return parent, child
// }

// func createData() map[string]interface{} {
// 	return map[string]interface{}{
// 		"string":     "string",
// 		"stringNot":  "string",
// 		"number":     1,
// 		"numberNot":  1,
// 		"boolean":    true,
// 		"booleanNot": true,
// 	}
// }

// func createFilter() map[string]interface{} {
// 	return map[string]interface{}{
// 		"string":  struct{}{},
// 		"number":  struct{}{},
// 		"boolean": struct{}{},
// 	}
// }

// func TestCreate(t *testing.T) {
// 	parent, child := createTree()
// 	filter := createFilter()
// 	pipeline := NewFilterPipeline(filter)

// 	parent.AddChild(child)
// 	parent.AddPipeline(child, pipeline)

// 	data := createData()
// 	command.Payload = data
// 	parent.Send(command)
// }
