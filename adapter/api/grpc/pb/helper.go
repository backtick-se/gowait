package pb

import (
	"github.com/backtick-se/gowait/core"
	"github.com/backtick-se/gowait/core/executor"
	"github.com/backtick-se/gowait/core/task"

	"encoding/json"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func UnpackTaskHeader(h *Header) task.Header {
	return task.Header{
		ID:   h.Id,
		Time: h.Time.AsTime(),
	}
}

func UnpackExecutorHeader(h *Header) executor.Header {
	return executor.Header{
		ID:   h.Id,
		Time: h.Time.AsTime(),
	}
}

func PackTaskSpec(def *task.Spec) *TaskSpec {
	return &TaskSpec{
		Name:    def.Name,
		Image:   def.Image,
		Command: def.Command,
		Input:   string(def.Input),
		Timeout: int64(def.Timeout),
		Time:    timestamppb.New(def.Time),
	}
}

func UnpackTaskSpec(def *TaskSpec) *task.Spec {
	return &task.Spec{
		Name:    def.Name,
		Image:   def.Image,
		Command: def.Command,
		Input:   json.RawMessage(def.Input),
		Timeout: int(def.Timeout),
		Time:    def.Time.AsTime(),
	}
}

func PackTaskState(s *task.Run) *Task {
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

func UnpackTaskState(s *Task) task.Run {
	var err error
	if s.Error != "" {
		err = core.NewError(s.Error)
	}
	return task.Run{
		ID:        task.ID(s.TaskId),
		Parent:    task.ID(s.Parent),
		Status:    task.Status(s.Status),
		Spec:      UnpackTaskSpec(s.Spec),
		Scheduled: s.Scheduled.AsTime(),
		Started:   s.Started.AsTime(),
		Completed: s.Completed.AsTime(),
		Result:    task.Result(s.Result),
		Err:       err,
	}
}
