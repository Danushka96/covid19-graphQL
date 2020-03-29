package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"net/http"
)

// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		_ = json.NewEncoder(w).Encode(result)
	})
	// Display some basic instructions
	fmt.Println("Now server is running on port 8080")

	_ = http.ListenAndServe(":8080", nil)
}
