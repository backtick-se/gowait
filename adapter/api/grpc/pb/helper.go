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

func PackTaskState(s *core.TaskState) *Task {
	err := ""
	if s.Err != nil {
		err = s.Err.Error()
	}
	return &Task{
		TaskId:    string(s.ID),
		Parent:    string(s.Parent),
		Status:    string(s.Status),
		Spec:      PackTaskSpec(s.Spec),
		Scheduled: timestamppb.New(s.Scheduled),
		Started:   timestamppb.New(s.Started),
		Completed: timestamppb.New(s.Completed),
		Result:    string(s.Result),
		Error:     err,
	}
}

func UnpackTaskState(s *Task) core.TaskState {
	var err error
	if s.Error != "" {
		err = core.NewError(s.Error)
	}
	return core.TaskState{
		ID:        core.TaskID(s.TaskId),
		Parent:    core.TaskID(s.Parent),
		Status:    core.TaskStatus(s.Status),
		Spec:      UnpackTaskSpec(s.Spec),
		Scheduled: s.Scheduled.AsTime(),
		Started:   s.Started.AsTime(),
		Completed: s.Completed.AsTime(),
		Result:    json.RawMessage(s.Result),
		Err:       err,
	}
}
