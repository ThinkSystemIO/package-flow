package pipeline

import (
	"github.com/thinksystemio/package-flow/command"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FilterPipeline struct {
	BasePipeline
	Filter map[string]struct{}
}

//
// FilterPipeline Base
//

func NewFilterPipeline(cmd *command.CreatePipeline) Pipeline {
	pipeline := &FilterPipeline{
		Filter: map[string]struct{}{},
	}

	pipeline.ID = primitive.NewObjectID().Hex()
	pipeline.Name = cmd.Name
	pipeline.Type = "filter"
	return pipeline
}

//
// FilterPipeline Command API
//

func (pipeline *FilterPipeline) UpdatePipelineFilter(cmd *command.UpdateFilterPipeline) {
	pipeline.Filter = cmd.Filter
}

//
// FilterPipeline Utils
//

func (pipeline *FilterPipeline) Apply(cmd command.Command) {
	if data, ok := cmd.GetData().(map[string]interface{}); ok {
		transformed := make(map[string]interface{}, len(data))
		for k := range pipeline.Filter {
			if value, ok := data[k]; ok {
				transformed[k] = value
			}
		}
		cmd.SetData(transformed)
		return
	}

	if data, ok := cmd.GetData().([]map[string]interface{}); ok {
		transformed := make([]map[string]interface{}, len(data))
		for _, item := range data {
			copied := make(map[string]interface{}, len(item))
			for k := range pipeline.Filter {
				if value, ok := item[k]; ok {
					copied[k] = value
				}
			}
			transformed = append(transformed, copied)
		}
		cmd.SetData(transformed)
		return
	}
}
