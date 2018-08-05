package object

import (
	"github.com/foreverzmy/ledger/db"
	"github.com/foreverzmy/ledger/model"
	"github.com/globalsign/mgo/bson"
	"github.com/graphql-go/graphql"
	"time"
)

var addBillField = graphql.Field{
	Type:        graphql.Boolean,
	Description: "添加账单",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Description: "账户 id",
			Type:        graphql.NewNonNull(graphql.String),
		},
		"time": &graphql.ArgumentConfig{
			Description:  "站单时间",
			Type:         graphql.Int,
			DefaultValue: time.Now().UnixNano() / 1e6,
		},
		"amount": &graphql.ArgumentConfig{
			Description: "账单金额",
			Type:        graphql.NewNonNull(graphql.Int),
		},
		"class": &graphql.ArgumentConfig{
			Description: "支出类型",
			Type:        graphql.NewNonNull(graphql.Int),
		},
		"remark": &graphql.ArgumentConfig{
			Description:  "备注",
			Type:         graphql.String,
			DefaultValue: "",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := p.Args["id"].(string)
		time := p.Args["time"].(int64)
		amount := p.Args["amount"].(int64)
		class := p.Args["class"].(int)
		remark := p.Args["remark"].(string)

		return nil, nil
	},
}
