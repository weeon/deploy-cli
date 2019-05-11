package biz

import (
	"context"
	"fmt"
	"os"

	"github.com/weeon/proto/ops"
	"google.golang.org/grpc"
)

var (
	client ops.OpsSrvClient
)

func env(k string) string {
	return os.Getenv(k)
}

func Init() error {
	address := env("OPS_SRV_ADDR")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client = ops.NewOpsSrvClient(conn)
	return nil
}

func Deploy() {
	project := env("OPS_PROJECT")
	workloadID := env("OPS_WORKLOAD_ID")
	token := env("OPS_TOKEN")

	metadata := fmt.Sprintf("job id %s", env("CI_JOB_ID"))

	resp, err := client.Deploy(context.Background(), &ops.DeployRequest{
		Project:    project,
		WorkloadID: workloadID,
		Token:      token,
		Metadata:   metadata,
	})
	if err != nil {
		fmt.Println("deploy error ", err)
		return
	}

	fmt.Println("deploy result ", resp)
}
