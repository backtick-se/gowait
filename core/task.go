package core

type Task interface {
	ID() TaskID
	Spec() *TaskSpec
	State() TaskState
	Logs(file string) []string
}
