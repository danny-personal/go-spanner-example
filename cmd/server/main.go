package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	singerv1 "github.com/danny-yamamoto/go-spanner-example/gen/singer/v1"
	"github.com/danny-yamamoto/go-spanner-example/gen/singer/v1/singerv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

/*
var (

	testProjectName  = "testing-project"
	testInstanceName = "testing-instance"
	testDatabase     = "test_database"

)
*/
type SingerServer struct{}

func (s *SingerServer) Singer(
	ctx context.Context,
	req *connect.Request[singerv1.SingerRequest],
) (*connect.Response[singerv1.SingerResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&singerv1.SingerResponse{
		SingerId: req.Msg.SingerId,
	})
	/*
		db := fmt.Sprintf("projects/%s/instances/%s/databases/%s", testProjectName, testInstanceName, testDatabase)
		client, err := spanner.NewClient(ctx, db)
		if err != nil {
			log.Println(err)
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
			if err != nil {
				fmt.Println("Failed to iterate:", err)
			}
			var singerID int64
			var firstName, lastName string
			if err := row.Columns(&singerID, &firstName, &lastName); err != nil {
				fmt.Println("Failed to Columns fetches:", err)
			}
			fmt.Printf("%d %s %s\n", singerID, firstName, lastName)
		}
	*/
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	fmt.Println("hello")
	mux := http.NewServeMux()
	singer := &SingerServer{}
	path, handler := singerv1connect.NewSingerServiceHandler(singer)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}))
}
