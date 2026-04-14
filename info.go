package main

func info() PluginInfo {
	return PluginInfo{
		Name:        "graphql",
		Version:     "0.1.0",
		Description: "GraphQL executor — queries, mutations, variables, error handling.",
		Protocols:   []string{"graphql"},
		Fields: map[string]FieldSpec{
			"query": {
				Type:        "string",
				Required:    false,
				Description: "GraphQL query string. Use this OR mutation.",
			},
			"mutation": {
				Type:        "string",
				Required:    false,
				Description: "GraphQL mutation string. Use this OR query.",
			},
			"variables": {
				Type:        "map",
				Required:    false,
				Description: "Variables passed with the query/mutation.",
			},
			"operation_name": {
				Type:        "string",
				Required:    false,
				Description: "Operation name when the document contains multiple operations.",
			},
			"fragments": {
				Type:        "array",
				Required:    false,
				Description: "Named fragments referenced by the query.",
			},
		},
	}
}
