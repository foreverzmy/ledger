package object

import "github.com/graphql-go/graphql"

// Schema 类型
var Schema graphql.Schema

func init() {
	var err error
	Schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Mutation: mutationType,
			Query:    queryType,
		},
	)

	if err != nil {
		panic(err)
	}
}
