package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

type InputInfo struct {
	Type      string
	Mandatory bool
}

type QueryDefinitionInfo struct {
	Name       string
	Inputs     map[string]InputInfo
	OutputType string
	IsList     bool
}

type GraphQLQueryInfo struct {
	Type      string
	GivenName string
	QueryName string
	Inputs    []string
	FullQuery string
}

// parseGraphQLQueriesAndMutations parses GraphQL queries and mutations from a file
func parseGraphQLQueriesAndMutations(filename string) ([]GraphQLQueryInfo, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	content := string(data)

	queryRegex := regexp.MustCompile(`(query|mutation)\s+(\w+)\s*\(([^)]*)\)\s*{\s*(\w+)\s*\(([^)]*)\)\s*}|(query|mutation)\s+(\w+)\s*\(([^)]*)\)\s*{\s*(\w+)\s*\(([^)]*)\)\s*{(?:[^{}]*|{(?:[^{}]*|{(?:[^{}]*|{(?:[^{}]*|{(?:[^{}]*|{[^{}]*})*})*})*})*})*}\s*}|(query|mutation)\s+(\w+)\s*\(([^)]*)\)\s*{\s*(\w+)\s*\(([^)]*)\)\s*{(?:[^{}]*|{(?:[^{}]*|{(?:[^{}]*|{(?:[^{}]*|{(?:[^{}]*|{(?:[^{}]*|{[^{}]*})*})*})*})*})*})*}\s*}`)
	matches := queryRegex.FindAllStringSubmatch(content, -1)

	//queryRegex := regexp.MustCompile(`(query|mutation)\s+(\w+)\s*\(([^)]*)\)\s*{\s*(\w+)\s*\(([^)]*)\)\s*{(?:[^{}]*|{(?:[^{}]*|{(?:[^{}]*|{(?:[^{}]*|{(?:[^{}]*|{[^{}]*})*})*})*})*})*}|(query|mutation)\s+(\w+)\s*\(([^)]*)\)\s*{\s*(\w+)\s*\(([^)]*)\)\s*{(?:[^{}]*|{(?:[^{}]*|{(?:[^{}]*|{(?:[^{}]*|{(?:[^{}]*|{(?:[^{}]*|{[^{}]*})*})*})*})*})*})*})*}`)
	//matches := queryRegex.FindAllStringSubmatch(content, -1)

	var graphqlInfos []GraphQLQueryInfo

	for _, match := range matches {
		var graphqlType string
		var givenName string
		var fullQuery string
		var queryName string
		var inputsString string

		if match[1] != "" {
			graphqlType = match[1]
			givenName = match[2]
			fullQuery = match[0]
			queryName = match[4]
			inputsString = match[3]
		} else {
			graphqlType = match[6]
			givenName = match[7]
			fullQuery = match[0]
			queryName = match[9]
			inputsString = match[8]
		}

		var inputs []string
		if inputsString != "" {
			inputPairs := strings.Split(inputsString, ",")
			for _, inputPair := range inputPairs {
				inputParts := strings.Split(strings.TrimSpace(inputPair), ":")
				inputName := strings.TrimSpace(inputParts[0])
				inputs = append(inputs, inputName)
			}
		}

		graphqlInfo := GraphQLQueryInfo{
			Type:      graphqlType,
			GivenName: givenName,
			QueryName: queryName,
			Inputs:    inputs,
			FullQuery: fullQuery,
		}

		graphqlInfos = append(graphqlInfos, graphqlInfo)
	}

	return graphqlInfos, nil
}

func parseGraphQLSchema(filename string) ([]QueryDefinitionInfo, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	content := string(data)

	queryRegex := regexp.MustCompile(`(\w+)\(([^)]*)\):\s*\[?(\w+)\!?]?\!?`)

	matches := queryRegex.FindAllStringSubmatch(content, -1)

	var queries []QueryDefinitionInfo

	for _, match := range matches {
		queryName := match[1]
		inputs := match[2]
		outputType := match[3]
		isList := strings.Contains(match[0], "[")

		inputMap := make(map[string]InputInfo)
		if inputs != "" {
			inputPairs := strings.Split(inputs, ",")
			for _, inputPair := range inputPairs {
				inputParts := strings.Split(strings.TrimSpace(inputPair), ":")
				inputName := strings.TrimSpace(inputParts[0])
				inputType := strings.TrimSpace(inputParts[1])
				mandatory := strings.HasSuffix(inputType, "!")
				inputType = strings.TrimSuffix(inputType, "!")
				inputMap[inputName] = InputInfo{
					Type:      inputType,
					Mandatory: mandatory,
				}
			}
		}

		queryInfo := QueryDefinitionInfo{
			Name:       queryName,
			Inputs:     inputMap,
			OutputType: outputType,
			IsList:     isList,
		}

		queries = append(queries, queryInfo)
	}

	return queries, nil
}
