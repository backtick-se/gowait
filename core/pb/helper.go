package pb

import (
	"cowait/core"
	"cowait/core/msg"
	"encoding/json"
)

func UnpackHeader(h *Header) msg.Header {
	return msg.Header{
		ID:   core.TaskID(h.Id),
		Time: h.Time.AsTime(),
	}
}

func PackTaskSpec(def *core.TaskSpec) *TaskSpec {
	return &TaskSpec{
		Cluster: def.Cluster,
		Name:    def.Name,
		Image:   def.Image,
		Command: def.Command,
		Input:   string(def.Input),
		Timeout: int64(def.Timeout),
	}
}

func UnpackTaskSpec(def *TaskSpec) *core.TaskSpec {
	return &core.TaskSpec{
		Cluster: def.Cluster,
		Name:    def.Name,
		Image:   def.Image,
		Command: def.Command,
		Input:   json.RawMessage(def.Input),
		Timeout: int(def.Timeout),
	}
}
