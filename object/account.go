package object

import (
	"github.com/foreverzmy/ledger/db"
	"github.com/graphql-go/graphql"
)

// addAccountField 创建账户
var addAccountField = graphql.Field{
	Type:        graphql.Boolean,
	Description: "添加用户",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Description: "用户昵称",
			Type:        graphql.String,
		},
		"avatar": &graphql.ArgumentConfig{
			Description: "用户头像",
			Type:        graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name := p.Args["name"].(string)
		avatar := p.Args["avatar"].(string)
		err := db.AddAccount(name, avatar)
		if err != nil {
			return false, err
		}
		return true, nil
	},
}
