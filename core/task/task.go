package task

type T interface {
	ID() ID
	Spec() *Spec
	State() State
	Logs(file string) ([]string, error)
}
