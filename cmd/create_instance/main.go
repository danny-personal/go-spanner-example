// Adapted from https://github.com/GoogleCloudPlatform/golang-samples/
package main

import (
	"context"
	"fmt"
	"os"

	instanceadmin "cloud.google.com/go/spanner/admin/instance/apiv1"
	instanceadminpb "cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	testProjectName  = "testing-project"
	testInstanceName = "testing-instance"
)

func main() {
	if v := os.Getenv("SPANNER_EMULATOR_HOST"); v == "" {
		fmt.Println("SPANNER_EMULATOR_HOST is not set")
		return
	}
	ctx := context.Background()
	client, err := instanceadmin.NewInstanceAdminClient(ctx)
	if err != nil {
		fmt.Println("Failed to create instance client:", err)
		return
	}
	defer client.Close()
	fmt.Println(err)

	op, err := client.CreateInstance(ctx, &instanceadminpb.CreateInstanceRequest{
		Parent:     fmt.Sprintf("projects/%s", testProjectName),
		InstanceId: testInstanceName,
	})
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			fmt.Println("Instance already exists", err)
			return
		}
		fmt.Println("Failed to create instance:", err)
		return
	}

	if _, err := op.Wait(ctx); err != nil {
		fmt.Println("Failed to wait operation:", err)
		return
	}
}
