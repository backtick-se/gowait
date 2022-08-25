package grpc

import "github.com/backtick-se/gowait/adapter/api/grpc/pb"

type execLog struct {
	client *executorClient
	stream pb.Executor_TaskLogClient
}

func (t *execLog) Log(file, data string) error {
	return t.stream.Send(&pb.LogEntry{
		Header: t.client.header(),
		File:   file,
		Data:   data,
	})
}

func (t *execLog) Close() error {
	_, err := t.stream.CloseAndRecv()
	return err
}
