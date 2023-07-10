// Adapted from https://github.com/GoogleCloudPlatform/golang-samples/
package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

var (
	testProjectName  = "testing-project"
	testInstanceName = "testing-instance"
	testDatabase     = "test_database"
)

func main() {
	db := fmt.Sprintf("projects/%s/instances/%s/databases/%s", testProjectName, testInstanceName, testDatabase)
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, db)
	if err != nil {
		fmt.Println("Failed to create client:", err)
		return
	}
	defer client.Close()

	stmt := spanner.Statement{
		SQL: `SELECT SingerId, FirstName, LastName FROM Singers
				WHERE LastName = @lastName`,
		Params: map[string]interface{}{
			"lastName": "Richards",
		},
	}
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			fmt.Println("Failed to iterator.Done:", err)
			return
		}
		if err != nil {
			fmt.Println("Failed to iterate:", err)
			return
		}
		var singerID int64
		var firstName, lastName string
		if err := row.Columns(&singerID, &firstName, &lastName); err != nil {
			fmt.Println("Failed to Columns fetches:", err)
			return
		}
		fmt.Printf("%d %s %s\n", singerID, firstName, lastName)
	}
}
