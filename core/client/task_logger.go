package client

import "cowait/core/pb"

type TaskLogger interface {
	Log(file, data string) error
	Close() error
}

type taskLog struct {
	client *taskClient
	stream pb.Executor_TaskLogClient
}

func (t *taskLog) Log(file, data string) error {
	return t.stream.Send(&pb.LogEntry{
		Header: t.client.header(),
		File:   file,
		Data:   data,
	})
}

func (t *taskLog) Close() error {
	_, err := t.stream.CloseAndRecv()
	return err
}
