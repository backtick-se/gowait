package task

import (
	"encoding/json"
	"time"
)

const EnvTaskID = "COWAIT_ID"
const EnvParentID = "COWAIT_PARENT"
const EnvTaskdef = "COWAIT_TASK"

type Spec struct {
	Name    string
	Image   string
	Command []string
	Input   json.RawMessage
	Timeout int
	Time    time.Time
}

func (t *Spec) ToEnv() (string, error) {
	if len(t.Input) == 0 {
		t.Input = json.RawMessage("{}")
	}
	encoded, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func SpecFromEnv(env string) (*Spec, error) {
	def := new(Spec)
	err := json.Unmarshal([]byte(env), def)
	if err != nil {
		return nil, err
	}
	return def, nil
}
