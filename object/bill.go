package object

import (
	"errors"
	"time"

	"github.com/foreverzmy/ledger/db"
	"github.com/foreverzmy/ledger/model"
	"github.com/globalsign/mgo/bson"
	"github.com/graphql-go/graphql"
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
			Remark:  remark,
		}

		err := db.AddBill(id, bill)
		if err != nil {
			return false, err
		}
		return true, nil
	},
}

var billField = graphql.Fields{
	"id": &graphql.Field{
		Description: "账单 ID",
		Type:        graphql.String,
	},
	"payment": &graphql.Field{
		Description: "支出账户",
		Type:        graphql.String,
	},
	"time": &graphql.Field{
		Description: "创建时间",
		Type:        graphql.Int,
	},
	"amount": &graphql.Field{
		Description: "支出金额",
		Type:        graphql.Int,
	},
	"class": &graphql.Field{
		Description: "账单类型",
		Type:        graphql.String,
	},
	"remark": &graphql.Field{
		Description: "备注",
		Type:        graphql.String,
	},
}

var queryBillField = graphql.Field{
	Description: "获取指定用户的某一条账单",
	Type: graphql.NewObject(graphql.ObjectConfig{
		Name:   "bill",
		Fields: billField,
	}),
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Description: "账单 ID",
			Type:        graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := p.Args["id"].(string)

		account := p.Source.(model.Account)

		billInAccount := false
		for _, billID := range account.Bills {
			if billID.Hex() == id {
				billInAccount = true
				break
			}
		}

		if !billInAccount {
			return nil, errors.New("该账单不属于用户")
		}

		bill, err := db.GetBill(id)

		return bill, err
	},
}

// var queryBillsField = graphql.Field{
// 	Description: "账单列表",
// 	Type:graphql.NewList(
// 		graphql.NewObject(graphql.ObjectConfig{
// 			Name:"bills",
// 			Fields:graphql.Fields{
// 				"id":&graphql.Field{
// 					Description:"账单 ID",

// 				},
// 			},
// 		})
// 	),
// 	Args:graphql.FieldArgument{

// 	}
// }
