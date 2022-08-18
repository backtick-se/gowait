package pb

import (
	"cowait/core"
	"cowait/core/msg"
)

func ParseHeader(h *Header) msg.Header {
	return msg.Header{
		ID:      core.TaskID(h.TaskId),
		Version: h.Version,
		Time:    h.Time.AsTime(),
	}
}
