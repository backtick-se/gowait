package pb

import (
	"cowait/core"
	"cowait/core/msg"
	"encoding/json"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func UnpackHeader(h *Header) msg.Header {
	return msg.Header{
		ID:   h.Id,
		Time: h.Time.AsTime(),
	}
}

func PackTaskSpec(def *core.TaskSpec) *TaskSpec {
	return &TaskSpec{
		Name:    def.Name,
		Image:   def.Image,
		Command: def.Command,
		Input:   string(def.Input),
		Timeout: int64(def.Timeout),
		Time:    timestamppb.New(def.Time),
	}
}

func UnpackTaskSpec(def *TaskSpec) *core.TaskSpec {
	return &core.TaskSpec{
		Name:    def.Name,
		Image:   def.Image,
		Command: def.Command,
		Input:   json.RawMessage(def.Input),
		Timeout: int(def.Timeout),
		Time:    def.Time.AsTime(),
	}
}
