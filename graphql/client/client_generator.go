package main

import (
	"fmt"
	"os"
	"strings"
)

func normalizeType(graphQLType string) string {
	switch graphQLType {
	case "String":
		return "string"
	case "Int":
		return "int"
	case "Float":
		return "float64"
	case "Boolean":
		return "bool"
	case "ID":
		return "string"
	default:
		return "model." + graphQLType
	}
}

func defaultOutputValue(outputType string) string {
	switch outputType {
	case "string":
		return "\"\""
	case "int":
		return "0"
	case "float64":
		return "0.0"
	case "bool":
		return "false"
	default:
		return outputType + "{}"
	}
}

func GenerateFunctionDefinitions(queryDefinitions []QueryDefinitionInfo, graphqlQueries []GraphQLQueryInfo) string {
	var sb strings.Builder

	// Write package declaration
	sb.WriteString("package graphql\n\n")
	sb.WriteString("import (\n")
	sb.WriteString("\t\"bytes\"\n")
	sb.WriteString("\t\"encoding/json\"\n")
	sb.WriteString("\t\"fmt\"\n")
	sb.WriteString("\t\"io/ioutil\"\n")
	sb.WriteString("\t\"net/http\"\n")
	sb.WriteString("\t\"os\"\n")
	sb.WriteString("\t\"github.com/edgar-care/edgarlib/graphql/model\"\n")
	sb.WriteString(")\n\n")

	queryDefMap := make(map[string]QueryDefinitionInfo)
	for _, queryDef := range queryDefinitions {
		queryDefMap[queryDef.Name] = queryDef
	}
	for _, gqlQuery := range graphqlQueries {
		queryDef, exists := queryDefMap[gqlQuery.QueryName]
		if !exists {
			fmt.Printf("Warning: No matching definition for %s\n", gqlQuery.QueryName)
			continue
		}

		funcSignature := fmt.Sprintf("func %s(", gqlQuery.GivenName)

		variablesMap := "\tvariables := map[string]interface{}{\n"

		for i, input := range gqlQuery.Inputs {
			inputName := input[1:] // Remove the $ from input name
			inputInfo, exists := queryDef.Inputs[inputName]
			if !exists {
				fmt.Printf("Warning: No matching input definition for %s in %s\n", input, gqlQuery.QueryName)
				continue
			}
			if i > 0 {
				funcSignature += ", "
			}
			inputType := normalizeType(inputInfo.Type)
			if inputInfo.Mandatory {
				funcSignature += fmt.Sprintf("%s %s", inputName, inputType)
				variablesMap += fmt.Sprintf("\t\t\"%s\": %s,\n", inputName, inputName)
			} else {
				funcSignature += fmt.Sprintf("%s *%s", inputName, inputType)
				variablesMap += fmt.Sprintf("\t\t\"%s\": %s,\n", inputName, inputName)
			}
		}

		funcSignature += ") ("

		outputType := normalizeType(queryDef.OutputType)
		if queryDef.IsList {
			funcSignature += fmt.Sprintf("[]%s, error", outputType)
		} else {
			funcSignature += fmt.Sprintf("%s, error", outputType)
		}

		funcSignature += ") {\n"

		queryString := strings.ReplaceAll(gqlQuery.FullQuery, "\n", "\n\t")
		funcBody := fmt.Sprintf("\tquery := `%s`\n", queryString)
		variablesMap += "\t}\n"
		funcBody += variablesMap
		funcBody += "\treqBody := map[string]interface{}{\n"
		funcBody += "\t\t\"query\": query,\n"
		funcBody += "\t\t\"variables\": variables,\n"
		funcBody += "\t}\n"

		funcBody += "\tbody, err := json.Marshal(reqBody)\n"
		funcBody += "\tif err != nil {\n"
		if queryDef.IsList {
			funcBody += fmt.Sprintf("\t\treturn nil, err\n")
		} else {
			funcBody += fmt.Sprintf("\t\treturn %s, err\n", defaultOutputValue(outputType))
		}
		funcBody += "\t}\n\n"

		funcBody += "\tresp, err := http.Post(os.Getenv(\"GRAPHQL_URL\"), \"application/json\", bytes.NewBuffer(body))\n"
		funcBody += "\tif err != nil {\n"
		if queryDef.IsList {
			funcBody += fmt.Sprintf("\t\treturn nil, err\n")
		} else {
			funcBody += fmt.Sprintf("\t\treturn %s, err\n", defaultOutputValue(outputType))
		}
		funcBody += "\t}\n"
		funcBody += "\tdefer resp.Body.Close()\n\n"

		funcBody += "\tif resp.StatusCode != http.StatusOK {\n"
		if queryDef.IsList {
			funcBody += fmt.Sprintf("\t\treturn nil, fmt.Errorf(\"failed to fetch data: %v\", resp.Status)\n", "%v")
		} else {
			funcBody += fmt.Sprintf("\t\treturn %s, fmt.Errorf(\"failed to fetch data: %v\", resp.Status)\n", defaultOutputValue(outputType), "%v")
		}
		funcBody += "\t}\n\n"

		funcBody += "\tresponseBody, err := ioutil.ReadAll(resp.Body)\n"
		funcBody += "\tif err != nil {\n"
		if queryDef.IsList {
			funcBody += fmt.Sprintf("\t\treturn nil, err\n")
		} else {
			funcBody += fmt.Sprintf("\t\treturn %s, err\n", defaultOutputValue(outputType))
		}
		funcBody += "\t}\n\n"

		if queryDef.IsList {
			funcBody += fmt.Sprintf("\tvar result struct {\n\t\tErrors []model.GraphQLError `json:\"errors\"`\n\t\tData struct {\n\t\t\t%s []%s `json:\"%s\"`\n\t\t} `json:\"data\"`\n\t}\n", gqlQuery.GivenName, normalizeType(queryDef.OutputType), gqlQuery.QueryName)
		} else {
			funcBody += fmt.Sprintf("\tvar result struct {\n\t\tErrors []model.GraphQLError `json:\"errors\"`\n\t\tData struct {\n\t\t\t%s %s `json:\"%s\"`\n\t\t} `json:\"data\"`\n\t}\n", gqlQuery.GivenName, normalizeType(queryDef.OutputType), gqlQuery.QueryName)
		}
		funcBody += "\terr = json.Unmarshal(responseBody, &result)\n"
		funcBody += "\tif err != nil {\n"
		if queryDef.IsList {
			funcBody += fmt.Sprintf("\t\treturn nil, err\n")
		} else {
			funcBody += fmt.Sprintf("\t\treturn %s, err\n", defaultOutputValue(outputType))
		}
		funcBody += "\t}\n\n"

		funcBody += "\tif len(result.Errors) > 0 {\n"
		if queryDef.IsList {
			funcBody += fmt.Sprintf("\t\treturn nil, fmt.Errorf(\"GraphQL error: %s\", result.Errors[0].Message)\n", "%s")
		} else {
			funcBody += fmt.Sprintf("\t\treturn %s, fmt.Errorf(\"GraphQL error: %s\", result.Errors[0].Message)\n", defaultOutputValue(outputType), "%s")
		}
		funcBody += "\t}\n\n"

		if queryDef.IsList {
			funcBody += fmt.Sprintf("\treturn result.Data.%s, nil\n", gqlQuery.GivenName)
		} else {
			funcBody += fmt.Sprintf("\treturn result.Data.%s, nil\n", gqlQuery.GivenName)
		}

		funcSignature += funcBody
		funcSignature += "}\n"

		sb.WriteString(funcSignature)
		sb.WriteString("\n")
	}

	return sb.String()
}

func WriteToFile(filename string, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
