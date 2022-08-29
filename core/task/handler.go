package task

type Handler interface {
	OnInit(*MsgInit) error
	OnComplete(*MsgComplete) error
	OnFailure(*MsgFailure) error
	OnLog(*MsgLog) error
}
