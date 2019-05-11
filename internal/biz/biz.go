package biz

import (
	"github.com/weeon/proto/ops"
	"google.golang.org/grpc"
	"os"
)

var (
	client ops.OpsSrvClient
)

func env(k string) string {
	return os.Getenv(k)
}

func Init() error {
	address := env("SRV_ADDR")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client = ops.NewOpsSrvClient(conn)
	return nil
}
