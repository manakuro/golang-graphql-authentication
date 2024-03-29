package graphql

import (
	"golang-graphql-authentication/graphql/field"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

func newQuery(db *gorm.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": field.NewUsers(db),
		},
	})
}
