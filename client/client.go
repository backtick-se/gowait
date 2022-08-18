package client

import (
	"cowait/core"
	"cowait/core/pb"

	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Client interface{}

type client struct {
	conn    *grpc.ClientConn
	tasks   pb.TaskClient
	version string
}

func NewClient() (Client, error) {
	conn, err := grpc.Dial("localhost:1337")
	if err != nil {
		return nil, err
	}

	tasks := pb.NewTaskClient(conn)

	return &client{
		conn:    conn,
		tasks:   tasks,
		version: "gowait/1.0",
	}, nil
}

func (c *client) header(id core.TaskID) *pb.Header {
	return &pb.Header{
		TaskId:  "",
		Version: c.version,
		Time:    timestamppb.New(time.Now()),
	}
}

func (c *client) Init(ctx context.Context, task *core.TaskDef) error {
	_, err := c.tasks.TaskInit(ctx, &pb.TaskInitReq{
		Header: c.header(task.ID),
		Taskdef: &pb.TaskDef{
			Id:      string(task.ID),
			Name:    task.Name,
			Image:   task.Image,
			Command: task.Command,
			Input:   string(task.Input),
		},
	})
	return err
}
