package main

import (
	"fmt"
)

func main() {
	queryDefinitions, err := parseGraphQLSchema("schema.graphql")

	graphqlQueries, err := parseGraphQLQueriesAndMutations("queries.graphql")
	if err != nil {
		fmt.Printf("Error parsing GraphQL queries and mutations: %v\n", err)
		return
	}

	content := GenerateFunctionDefinitions(queryDefinitions, graphqlQueries)
	err = WriteToFile("generated_client.go", content)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	} else {
		fmt.Println("File generated successfully: generated_functions.go")
	}
}
