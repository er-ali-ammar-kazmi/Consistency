package app

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func QueryType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"blogs": &graphql.Field{
					Type: graphql.NewList(graphql.NewObject(
						graphql.ObjectConfig{
							Name: "Blog",
							Fields: graphql.Fields{
								"id": &graphql.Field{
									Type: graphql.Int,
								},
								"title": &graphql.Field{
									Type: graphql.String,
								},
								"content": &graphql.Field{
									Type: graphql.String,
								},
							},
						},
					)),
					Resolve: resolver,
				},
			},
		},
	)
}

func Handler() *handler.Handler {

	queryObject := QueryType()
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryObject,
		},
	)
	if err != nil {
		fmt.Println("Error at schema creation ", err.Error())
	}

	handler := handler.New(
		&handler.Config{
			Schema:   &schema,
			Pretty:   true,
			GraphiQL: true,
		})
	return handler
}

func resolver(p graphql.ResolveParams) (any, error) {
	db := DbService()
	blogs, err := db.RetrieveAllBlogs()
	if err != nil {
		return nil, err
	}
	return blogs, nil
}
