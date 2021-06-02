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
	ns := env("OPS_NAMESPACE")

	metadata := fmt.Sprintf("JobID: %s commit Title:[%s] commit ID: [%s] user [%s] ",
		env("CI_JOB_ID"), env("CI_COMMIT_TITLE"),
		env("CI_COMMIT_SHA"), env("CI_DEPLOY_USER"))

	fmt.Printf("deploy  %s ,project %s   workload ID %s \n", ns, project, workloadID)

	resp, err := client.Deploy(context.Background(), &ops.DeployRequest{
		Project:    project,
		WorkloadID: workloadID,
		Token:      token,
		Metadata:   metadata,
		Namespace:  ns,
	})
	if err != nil {
		fmt.Println("deploy error ", err)
		TelegramNotify(fmt.Sprintf("❌ %s deploy fail %s metadata %s", workloadID,
			err.Error(), metadata))
		os.Exit(1)
	}

	fmt.Println("deploy  result ", resp)

	var icon string
	switch resp.Result {
	case ops.DeployResult_DeployResultSuccess:
		icon = "✅"
	case ops.DeployResult_DeployResultFail:
		icon = "❌"
	}

	format := `
*%s [%s/%s] Deploy Result*

- 🍊 Project  *%s*
- 🍉 Workload ID   *%s*   Result  %v 
- 🍭 Metadata  %s
`

	msg := fmt.Sprintf(format, icon, env("CI_PROJECT_NAMESPACE"), env("CI_PROJECT_NAME"),
		project, workloadID, resp, metadata)
	TelegramNotify(msg)
}
