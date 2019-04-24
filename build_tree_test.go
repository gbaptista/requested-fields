package fields

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildTree(t *testing.T) {
	var graphqlQueryA string = `

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

	var graphqlQueryB string = `
		query {
		  users {
		    id
		    title
		  }
		}`

	var graphqlQueryC string = `
		query {
		  hello
		}`

	var graphqlQueryD string = `
		{
		  user(id: 3) {
		    id
		    name
		  }
		}`

	var graphqlQueryE string = `
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

	var graphqlQueryF string = `{
		  users {
		    users {
		      users {
		        name
		      }
		    }
		  }
		}`

	var graphqlQueryG string = `
		{
		  ...Frag
		}

		fragment Frag on SomeType {
		  field {
		    sub_field
		  }
		}`

	expectedTreeA := map[string][]string{
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

	generatedTreeA := BuildTree(graphqlQueryA)

	assert.Equal(t, expectedTreeA[""], generatedTreeA[""])

	assert.Equal(t, expectedTreeA, generatedTreeA)

	expectedTreeB := map[string][]string{
		"":      []string{"users"},
		"users": []string{"id", "title"},
	}

	generatedTreeB := BuildTree(graphqlQueryB)

	assert.Equal(t, expectedTreeB, generatedTreeB)

	expectedTreeC := map[string][]string{
		"": []string{"hello"},
	}

	generatedTreeC := BuildTree(graphqlQueryC)

	assert.Equal(t, expectedTreeC, generatedTreeC)

	expectedTreeD := map[string][]string{
		"":     []string{"user"},
		"user": []string{"id", "name"},
	}

	generatedTreeD := BuildTree(graphqlQueryD)

	assert.Equal(t, expectedTreeD, generatedTreeD)

	expectedTreeE := map[string][]string{
		"":     []string{"user"},
		"user": []string{"id", "name", "age"},
	}

	generatedTreeE := BuildTree(graphqlQueryE)

	assert.Equal(t, expectedTreeE, generatedTreeE)

	expectedTreeF := map[string][]string{
		"":                  []string{"users"},
		"users":             []string{"users"},
		"users.users":       []string{"users"},
		"users.users.users": []string{"name"}}

	generatedTreeF := BuildTree(graphqlQueryF)

	assert.Equal(t, expectedTreeF, generatedTreeF)

	expectedTreeG := map[string][]string{
		"":      []string{"field"},
		"field": []string{"sub_field"}}

	generatedTreeG := BuildTree(graphqlQueryG)

	assert.Equal(t, expectedTreeG, generatedTreeG)
}
