package fields

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var graphql_query_products string = `
query {
  search {
    products {
      edges {
        node { id title }
        cursor
      }
    }
  }
}
`

type QueryResolver struct {
	Field Field `name:"query"`
}

type SearchResolver struct {
	Field Field `name:"search"`
}

type ProductsResolver struct {
	Field Field `name:"products"`
}

func TestRequestedFieldsForProducts(t *testing.T) {
	query_resolver := &QueryResolver{}

	search_resolver := &SearchResolver{}
	search_resolver.Field.SetParent(query_resolver)

	products_resolver := &ProductsResolver{}
	products_resolver.Field.SetParent(search_resolver)

	ctx := context.WithValue(context.Background(),
		"graphqlRequestTree", BuildTree(graphql_query_products))

	expected_fields := []string{"edges"}
	requested_fields := RequestedFor(ctx, products_resolver)

	assert.Equal(t, expected_fields, requested_fields)

	expected_fields = []string{"node", "cursor"}
	requested_fields = RequestedForAt(ctx, products_resolver, "edges")

	assert.Equal(t, expected_fields, requested_fields)

	expected_fields = []string{"id", "title"}
	requested_fields = RequestedForAt(ctx, products_resolver, "edges.node")

	assert.Equal(t, expected_fields, requested_fields)
}

var graphql_query_user string = `
{
  user(id: 3) {
    id
    name
  }
}
`

type UserResolver struct {
	Field Field `name:"user"`
}

func TestRequestedFieldsForUser(t *testing.T) {
	query_resolver := &QueryResolver{}

	user_resolver := &UserResolver{}
	user_resolver.Field.SetParent(query_resolver)

	ctx := context.WithValue(context.Background(),
		"graphqlRequestTree", BuildTree(graphql_query_user))

	expected_fields := []string{"id", "name"}
	requested_fields := RequestedFor(ctx, user_resolver)

	assert.Equal(t, expected_fields, requested_fields)
}
