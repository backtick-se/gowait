package grpc

import (
	"github.com/backtick-se/gowait/adapter/api/grpc/pb"
	"github.com/backtick-se/gowait/core/task"
)

type execLog struct {
	id     task.ID
	client *executorClient
	stream pb.Executor_TaskLogClient
}

func (t *execLog) Log(file, data string) error {
	return t.stream.Send(&pb.LogEntry{
		Header: t.client.header(string(t.id)),
		File:   file,
		Data:   data,
	})
}

func (t *execLog) Close() error {
	_, err := t.stream.CloseAndRecv()
	return err
}
