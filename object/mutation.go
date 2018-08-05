package object

import "github.com/graphql-go/graphql"

// mutationType 修改
var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "mutation",
	Fields: graphql.Fields{
		"addAccount": &addAccountField,
		"addAssets":  &addAssetsField,
	},
})
