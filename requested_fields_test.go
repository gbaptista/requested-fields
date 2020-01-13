package fields

import (
	"context"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
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
	Field Field `graphql:"query"`
}

type SearchResolver struct {
	Field Field `graphql:"search"`
}

type ProductsResolver struct {
	Field Field `graphql:"products"`
}

func TestRequestedFieldsForProducts(t *testing.T) {
	query_resolver := &QueryResolver{}

	search_resolver := &SearchResolver{}
	search_resolver.Field.SetParent(query_resolver)

	products_resolver := &ProductsResolver{}
	products_resolver.Field.SetParent(search_resolver)

	ctx := context.WithValue(context.Background(),
		ContextKey, BuildTree(graphql_query_products, Variables{}))

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
	Field Field `graphql:"user"`
}

func TestRequestedFieldsForUser(t *testing.T) {
	query_resolver := &QueryResolver{}

	user_resolver := &UserResolver{}
	user_resolver.Field.SetParent(query_resolver)

	ctx := context.WithValue(context.Background(),
		ContextKey, BuildTree(graphql_query_user, Variables{}))

	expected_fields := []string{"id", "name"}
	requested_fields := RequestedFor(ctx, user_resolver)

	assert.Equal(t, expected_fields, requested_fields)
}

var graphql_query_user_nested string = `
{
  user(id: 3) {
    id
    name
    user {
        id
        age
        height
    }
  }
}
`

func TestRequestedFieldsForContainingUser(t *testing.T) {
	query_resolver := &QueryResolver{}

	user_resolver := &UserResolver{}
	user_resolver.Field.SetParent(query_resolver)

	ctx := context.WithValue(context.Background(),
		ContextKey, BuildTree(graphql_query_user_nested, Variables{}))

	expected_fields := []string{"id", "name", "age", "height", "user"}
	requested_fields := RequestedForContaining(ctx, "user")

	sort.Slice(expected_fields, func(i, j int) bool { return expected_fields[i] < expected_fields[j] })
	sort.Slice(requested_fields, func(i, j int) bool { return requested_fields[i] < requested_fields[j] })

	assert.Equal(t, expected_fields, requested_fields)
}
