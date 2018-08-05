package main

import (
	"context"
	"github.com/graphql-go/handler"
	"net/http"

	"github.com/foreverzmy/ledger/object"
)

func main() {

	h := handler.New(&handler.Config{
		Schema:   &object.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.HandleFunc("/gp", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		h.ContextHandler(ctx, w, r)
	})

	http.ListenAndServe(":8099", nil)
}
