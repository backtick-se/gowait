package task

type T interface {
	ID() ID
	Spec() *Spec
	State() Run
	Logs(file string) ([]string, error)
}
