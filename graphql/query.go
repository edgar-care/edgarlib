package graphql

import (
	"context"
	"os"

	"github.com/machinebox/graphql"
)

func createClient() *graphql.Client {
	return graphql.NewClient(os.Getenv("GRAPHQL_URL"))
}

func Query(query string, variables map[string]interface{}, respData interface{}) error {
	var request = graphql.NewRequest(query)
	var ctx = context.Background()
	for key, value := range variables {
		request.Var(key, value)
	}
	request.Header.Set(os.Getenv("API_KEY"), os.Getenv("API_KEY_VALUE"))
	err := createClient().Run(ctx, request, respData)
	return err
}
