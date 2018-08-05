package object

import (
	"github.com/foreverzmy/ledger/db"
	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "query",
	Fields: graphql.Fields{
		"account": &graphql.Field{
			Description: "[用户管理] 获取指定用户的信息",
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name:        "account",
				Description: "用户信息描述",
				Fields: graphql.Fields{
					"_id": &graphql.Field{
						Description: "用户名称",
						Type:        graphql.String,
					},
					"name": &graphql.Field{
						Description: "用户密码",
						Type:        graphql.String,
					},
					"avatar": &graphql.Field{
						Description: "用户手机号",
						Type:        graphql.Int,
					},
				},
			}),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "账户ID",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(string)
				res, err := db.GetAccount(id)
				if err != nil {
					return false, err
				}
				return res, nil
			},
		},
	},
})
