// Adapted from https://github.com/GoogleCloudPlatform/golang-samples/
package main

import (
	"context"
	"fmt"
	"os"

	database "cloud.google.com/go/spanner/admin/database/apiv1"
	adminpb "cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
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
	adminClient, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		fmt.Println("Failed to create instance client:", err)
		return
	}
	defer adminClient.Close()
	op, err := adminClient.CreateDatabase(ctx, &adminpb.CreateDatabaseRequest{
		Parent:          fmt.Sprintf("projects/%s/instances/%s", testProjectName, testInstanceName),
		CreateStatement: "CREATE DATABASE test_database",
		ExtraStatements: []string{
			`CREATE TABLE Singers (
				SingerId   INT64 NOT NULL,
				FirstName  STRING(1024),
				LastName   STRING(1024),
				SingerInfo BYTES(MAX),
				FullName   STRING(2048) AS (
					ARRAY_TO_STRING([FirstName, LastName], " ")
				) STORED
			) PRIMARY KEY (SingerId)`,
			`CREATE TABLE Albums (
				SingerId     INT64 NOT NULL,
				AlbumId      INT64 NOT NULL,
				AlbumTitle   STRING(MAX)
			) PRIMARY KEY (SingerId, AlbumId),
			INTERLEAVE IN PARENT Singers ON DELETE CASCADE`,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := op.Wait(ctx); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Created database")
}
