package model

type GraphQLError struct {
	Message string   `json:"message"`
	Path    []string `json:"path"`
}
