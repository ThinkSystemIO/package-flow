package pipeline

import (
	"encoding/json"

	flowcommand "github.com/thinksystemio/package-flow/command"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BasePipeline struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewBasePipelineDefault(name string) Pipeline {
	return &BasePipeline{
		ID:   primitive.NewObjectID().Hex(),
		Name: name,
	}
}

func NewBasePipeline(cmd *flowcommand.CreatePipeline) Pipeline {
	return &BasePipeline{
		ID:   primitive.NewObjectID().Hex(),
		Name: cmd.Name,
	}
}

func (pipeline *BasePipeline) GetID() string {
	return pipeline.ID
}

func (pipeline *BasePipeline) GetName() string {
	return pipeline.Name
}

func (pipeline *BasePipeline) GetType() string {
	return pipeline.Type
}

func (pipeline *BasePipeline) Apply(cmd flowcommand.Command) {}

func (pipeline *BasePipeline) ToJSON() ([]byte, error) {
	return json.Marshal(pipeline)
}
