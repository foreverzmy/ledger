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
			Description:  "账单时间",
			Type:         graphql.Int,
			DefaultValue: time.Now().UnixNano() / 1e6,
		},
		"payment": &graphql.ArgumentConfig{
			Description: "支出账户",
			Type:        graphql.NewNonNull(graphql.String),
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
		payment := p.Args["payment"].(string)
		amount := int64(p.Args["amount"].(int))
		class := p.Args["class"].(int)
		remark := p.Args["remark"].(string)

		bill := &model.Bill{
			ID:      bson.NewObjectId(),
			Payment: payment,
			Time:    time,
			Amount:  amount,
			Class:   class,
			Remake:  remark,
		}

		err := db.AddBill(id, bill)
		if err != nil {
			return false, err
		}
		return true, nil
	},
}
