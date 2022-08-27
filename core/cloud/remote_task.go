package cloud

import (
	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/task"

	"fmt"
)

type instance struct {
	cluster cluster.T
	state   *task.Run
}

var _ task.T = &instance{}

func (i *instance) ID() task.ID { return i.state.ID }

func (i *instance) Spec() *task.Spec { return i.state.Spec }
func (i *instance) State() *task.Run { return i.state }

func (i *instance) Logs(file string) ([]string, error) {
	// todo: implement
	return nil, fmt.Errorf("not implemented")
}
