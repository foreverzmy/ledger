package main

import (
	"context"
	"net/http"

	"github.com/graphql-go/handler"

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

	http.ListenAndServe("127.0.0.1:8099", nil)
}
