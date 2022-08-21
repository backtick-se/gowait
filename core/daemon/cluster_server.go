package daemon

import "cowait/core"

type ClusterServer interface {
	Get(core.TaskID) (core.Task, bool)
	Schedule(*core.TaskSpec) (core.Task, error)
}
