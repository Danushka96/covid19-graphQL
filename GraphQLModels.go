package main

import "github.com/graphql-go/graphql"

var recordType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Records",
	Fields: graphql.Fields{
		"fips": &graphql.Field{
			Type: graphql.String,
		},
		"admin_2": &graphql.Field{
			Type: graphql.String,
		},
		"province_state": &graphql.Field{
			Type: graphql.String,
		},
		"country_region": &graphql.Field{
			Type: graphql.String,
		},
		"last_update": &graphql.Field{
			Type: graphql.String,
		},
		"lat": &graphql.Field{
			Type: graphql.String,
		},
		"long_": &graphql.Field{
			Type: graphql.String,
		},
		"confirmed": &graphql.Field{
			Type: graphql.String,
		},
		"deaths": &graphql.Field{
			Type: graphql.String,
		},
		"recovered": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.String,
		},
		"combined_key": &graphql.Field{
			Type: graphql.String,
		},
	},
})
