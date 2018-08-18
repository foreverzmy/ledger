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

var queryAccountField = graphql.Field{
	Description: "获取指定用户的某个资产信息",
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
				Description: "用户头像",
				Type:        graphql.String,
			},
			"asset":  &queryAssetField,
			"assets": &queryAssetsField,
			"bill":   &queryBillField,
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
}
