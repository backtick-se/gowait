package core

import (
	"encoding/json"
)

const EnvTaskdef = "COWAIT_TASK"

type TaskDef struct {
	ID        TaskID
	Name      string
	Namespace string
	Image     string
	Command   []string
	Input     json.RawMessage
	Timeout   int
}

func (t *TaskDef) ToEnv() (string, error) {
	encoded, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func TaskDefFromEnv(env string) (*TaskDef, error) {
	def := new(TaskDef)
	err := json.Unmarshal([]byte(env), def)
	if err != nil {
		return nil, err
	}
	return def, nil
}
