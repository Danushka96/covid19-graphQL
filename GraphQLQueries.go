package main

import (
	"github.com/graphql-go/graphql"
	"log"
	"time"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{

		"records_byDay": &graphql.Field{
			Type:        graphql.NewList(recordType),
			Description: "Get Records of a given date",
			Args: graphql.FieldConfigArgument{
				"date": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				dateQuery, isOK := params.Args["date"].(string)
				if isOK {
					result, e := readCSVFromUrl(dateQuery)
					if e != nil {
						log.Fatal("error reading content")
					}
					return result, nil
				}
				return []Todo{}, nil
			},
		},

		"records_latest": &graphql.Field{
			Type:        graphql.NewList(recordType),
			Description: "List of Records",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, e := readCSVFromUrl(time.Now().Add(-24 * time.Hour).Format("01-02-2006"))
				if e != nil {
					log.Fatal("error reading content")
				}
				return result, nil
			},
		},
	},
})
