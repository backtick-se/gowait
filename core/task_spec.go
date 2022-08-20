package core

import (
	"encoding/json"
)

const EnvTaskID = "COWAIT_ID"
const EnvParentID = "COWAIT_PARENT"
const EnvTaskdef = "COWAIT_TASK"

type TaskSpec struct {
	Cluster string
	Name    string
	Image   string
	Command []string
	Input   json.RawMessage
	Timeout int
}

func (t *TaskSpec) ToEnv() (string, error) {
	if len(t.Input) == 0 {
		t.Input = json.RawMessage("{}")
	}
	encoded, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func TaskDefFromEnv(env string) (*TaskSpec, error) {
	def := new(TaskSpec)
	err := json.Unmarshal([]byte(env), def)
	if err != nil {
		return nil, err
	}
	return def, nil
}
