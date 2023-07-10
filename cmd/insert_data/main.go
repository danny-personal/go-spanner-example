package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
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

	singerColumns := []string{"SingerId", "FirstName", "LastName"}
	albumColumns := []string{"SingerId", "AlbumId", "AlbumTitle"}
	m := []*spanner.Mutation{
		spanner.InsertOrUpdate("Singers", singerColumns, []interface{}{1, "Marc", "Richards"}),
		spanner.InsertOrUpdate("Singers", singerColumns, []interface{}{2, "Catalina", "Smith"}),
		spanner.InsertOrUpdate("Singers", singerColumns, []interface{}{3, "Alice", "Trentor"}),
		spanner.InsertOrUpdate("Singers", singerColumns, []interface{}{4, "Lea", "Martin"}),
		spanner.InsertOrUpdate("Singers", singerColumns, []interface{}{5, "David", "Lomond"}),
		spanner.InsertOrUpdate("Albums", albumColumns, []interface{}{1, 1, "Total Junk"}),
		spanner.InsertOrUpdate("Albums", albumColumns, []interface{}{1, 2, "Go, Go, Go"}),
		spanner.InsertOrUpdate("Albums", albumColumns, []interface{}{2, 1, "Green"}),
		spanner.InsertOrUpdate("Albums", albumColumns, []interface{}{2, 2, "Forever Hold Your Peace"}),
		spanner.InsertOrUpdate("Albums", albumColumns, []interface{}{2, 3, "Terrified"}),
	}
	_, err = client.Apply(ctx, m)
	if err != nil {
		fmt.Println("Failed to insert data:", err)
		return
	}
}
