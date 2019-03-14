package fields

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var graphql_query_a string = `

query productsSearch(

  $products_search: ProductsSearchInput!) {

  search(products: $products_search)


  {
    term
    products {
      edges {
        node {
          id           custom_title: title
          seller {
            ...SellerData
          }
        }
        cursor
      }
    }
  }

  search_users {
    term
    users {
      edges {
        node {
          id           custom_title: title
          seller {...SellerData}
        }
        cursor
      }
    }
  }
}

fragment SellerData on User {
  id, ...SellerDataB
} 

fragment

SellerDataB

on

  User {
  name
} `

// search {
//   term
//   products {
//     edges {
//       node {
//         id
//         title
//         seller {
//           id
//           name
//         }
//       }
//       cursor
//     }
//   }
// }
// search_users {
//   term
//   users {
//     edges {
//       node {
//         id
//         title
//         seller {
//           id
//           name
//         }
//       }
//       cursor
//     }
//   }
// }

var graphql_query_b string = `
query {
  users {
    id
    title
  }
}`

var graphql_query_c string = `
query {
  hello
}`

var graphql_query_d string = `
{
  user(id: 3) {
    id
    name
  }
}`

var graphql_query_e string = `
{
  user(id: 3) {
    id
    name
  }

  custom_user: user(id: 4) {
    id
    name
    age
  }
}`

var graphql_query_f string = `{
  users {
    users {
      users {
        name
      }
    }
  }
}`

var graphql_query_g string = `
{
  ...Frag
}

fragment Frag on SomeType {
  field {
    sub_field
  }
}`

func TestBuildTree(t *testing.T) {
	expected_tree_a := map[string][]string{
		"":                                     []string{"search", "search_users"},
		"search":                               []string{"term", "products"},
		"search.products":                      []string{"edges"},
		"search.products.edges":                []string{"node", "cursor"},
		"search.products.edges.node":           []string{"id", "title", "seller"},
		"search.products.edges.node.seller":    []string{"id", "name"},
		"search_users":                         []string{"term", "users"},
		"search_users.users":                   []string{"edges"},
		"search_users.users.edges":             []string{"node", "cursor"},
		"search_users.users.edges.node":        []string{"id", "title", "seller"},
		"search_users.users.edges.node.seller": []string{"id", "name"},
	}

	generated_tree_a := BuildTree(graphql_query_a)

	assert.Equal(t, expected_tree_a[""], generated_tree_a[""])

	assert.Equal(t, expected_tree_a, generated_tree_a)

	expected_tree_b := map[string][]string{
		"":      []string{"users"},
		"users": []string{"id", "title"},
	}

	generated_tree_b := BuildTree(graphql_query_b)

	assert.Equal(t, expected_tree_b, generated_tree_b)

	expected_tree_c := map[string][]string{
		"": []string{"hello"},
	}

	generated_tree_c := BuildTree(graphql_query_c)

	assert.Equal(t, expected_tree_c, generated_tree_c)

	expected_tree_d := map[string][]string{
		"":     []string{"user"},
		"user": []string{"id", "name"},
	}

	generated_tree_d := BuildTree(graphql_query_d)

	assert.Equal(t, expected_tree_d, generated_tree_d)

	expected_tree_e := map[string][]string{
		"":     []string{"user"},
		"user": []string{"id", "name", "age"},
	}

	generated_tree_e := BuildTree(graphql_query_e)

	assert.Equal(t, expected_tree_e, generated_tree_e)

	expected_tree_f := map[string][]string{
		"":                  []string{"users"},
		"users":             []string{"users"},
		"users.users":       []string{"users"},
		"users.users.users": []string{"name"}}

	generated_tree_f := BuildTree(graphql_query_f)

	assert.Equal(t, expected_tree_f, generated_tree_f)

	expected_tree_g := map[string][]string{
		"":      []string{"field"},
		"field": []string{"sub_field"}}

	generated_tree_g := BuildTree(graphql_query_g)

	assert.Equal(t, expected_tree_g, generated_tree_g)
}
