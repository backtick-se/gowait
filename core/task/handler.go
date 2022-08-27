package task

type Handler interface {
	Init(*MsgInit) error
	Complete(*MsgComplete) error
	Fail(*MsgFailure) error
	Log(*MsgLog) error
}
