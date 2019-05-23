package fields

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Variables map[string]interface{}

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

	generatedTreeA := BuildTree(graphqlQueryA, Variables{})

	assert.Equal(t, expectedTreeA[""], generatedTreeA[""])

	assert.Equal(t, expectedTreeA, generatedTreeA)

	expectedTreeB := map[string][]string{
		"":      []string{"users"},
		"users": []string{"id", "title"},
	}

	generatedTreeB := BuildTree(graphqlQueryB, Variables{})

	assert.Equal(t, expectedTreeB, generatedTreeB)

	expectedTreeC := map[string][]string{
		"": []string{"hello"},
	}

	generatedTreeC := BuildTree(graphqlQueryC, Variables{})

	assert.Equal(t, expectedTreeC, generatedTreeC)

	expectedTreeD := map[string][]string{
		"":     []string{"user"},
		"user": []string{"id", "name"},
	}

	generatedTreeD := BuildTree(graphqlQueryD, Variables{})

	assert.Equal(t, expectedTreeD, generatedTreeD)

	expectedTreeE := map[string][]string{
		"":     []string{"user"},
		"user": []string{"id", "name", "age"},
	}

	generatedTreeE := BuildTree(graphqlQueryE, Variables{})

	assert.Equal(t, expectedTreeE, generatedTreeE)

	expectedTreeF := map[string][]string{
		"":                  []string{"users"},
		"users":             []string{"users"},
		"users.users":       []string{"users"},
		"users.users.users": []string{"name"}}

	generatedTreeF := BuildTree(graphqlQueryF, Variables{})

	assert.Equal(t, expectedTreeF, generatedTreeF)

	expectedTreeG := map[string][]string{
		"":      []string{"field"},
		"field": []string{"sub_field"}}

	generatedTreeG := BuildTree(graphqlQueryG, Variables{})

	assert.Equal(t, expectedTreeG, generatedTreeG)

	var graphqlQueryH string = `
		query (
		  $product_id: ID!,
		  $first: Int!
		){
		  product(id: $product_id) {
		    id
		  }

		  search(products: $search) {
		    term
		  }

		  other_a: search(products: $search) {
		    products(
		      first: $first,
		      sort: $sort,
		    ) {
		      total
		    }
		  }

		  other_b: search(products: $search) {
		    products(
		      first: $first,
		      sort: $sort
		    ) {
		      total
		    }
		  }
		}
	`
	expectedTreeH := map[string][]string{
		"":                 []string{"product", "search", "other_a", "other_b"},
		"other_a":          []string{"products"},
		"other_a.products": []string{"total"},
		"other_b":          []string{"products"},
		"other_b.products": []string{"total"},
		"product":          []string{"id"},
		"search":           []string{"term"}}

	generatedTreeH := BuildTreeUsingAliases(graphqlQueryH, Variables{})

	assert.Equal(t, expectedTreeH, generatedTreeH)

	// -------------------------------------

	var graphqlQueryI string = `
		query ProductsSearchPage($include_aggregations: Boolean!) {
		  search {
		    aggregations {
		      departments {
		        ...departmentAggregationFields
		        __typename
		      }
		    }
		    some_field
		  }
		}

		fragment departmentAggregationFields on DepartmentAggregation {
		  slug
		  name
		  __typename
		}

	`
	expectedTreeI := map[string][]string{
		"":                                []string{"search"},
		"search":                          []string{"aggregations", "some_field"},
		"search.aggregations":             []string{"departments"},
		"search.aggregations.departments": []string{"slug", "name", "__typename"}}

	generatedTreeI := BuildTreeUsingAliases(graphqlQueryI, Variables{})

	assert.Equal(t, expectedTreeI, generatedTreeI)

	// -------------

	var graphqlQueryJ string = `
		query ProductsSearchPage($include_aggregations: Boolean!) {
		  search {
		    aggregations @include(if: true) {
		      departments {
		        ...departmentAggregationFields
		        __typename
		      }
		    }
		    some_field
		  }
		}

		fragment departmentAggregationFields on DepartmentAggregation {
		  slug
		  name
		  __typename
		}

	`
	expectedTreeJ := map[string][]string{
		"": []string{"search"},
		"search": []string{
			"aggregations", "some_field"},
		"search.aggregations": []string{"departments"},
		"search.aggregations.departments": []string{
			"slug", "name", "__typename"}}

	generatedTreeJ := BuildTreeUsingAliases(graphqlQueryJ, Variables{})

	assert.Equal(t, expectedTreeJ, generatedTreeJ)

	var graphqlQueryK string = `
		query ProductsSearchPage($include_aggregations: Boolean!) {
		  search {
		    aggregations @include(if:false) {
		      departments {
		        ...departmentAggregationFields
		        __typename
		      }
		    }
		    some_field
		  }
		}

		fragment departmentAggregationFields on DepartmentAggregation {
		  slug
		  name
		  __typename
		}

	`
	expectedTreeK := map[string][]string{
		"": []string{"search"},
		"search": []string{
			"aggregations_FALSE", "some_field"},
		"search.aggregations_FALSE": []string{"departments"},
		"search.aggregations_FALSE.departments": []string{
			"slug", "name", "__typename"}}

	generatedTreeK := BuildTreeUsingAliases(graphqlQueryK, Variables{})

	assert.Equal(t, expectedTreeK, generatedTreeK)

	var graphqlQueryL string = `
		query ProductsSearchPage($include_aggregations: Boolean!) {
		  search {
		    aggregations @include(if: $include_aggregations) {
		      departments {
		        ...departmentAggregationFields
		        __typename
		      }
		    }
		    some_field
		  }
		}

		fragment departmentAggregationFields on DepartmentAggregation {
		  slug
		  name
		  __typename
		}

	`
	expectedTreeL := map[string][]string{
		"": []string{"search"},
		"search": []string{
			"aggregations_FALSE", "some_field"},
		"search.aggregations_FALSE": []string{"departments"},
		"search.aggregations_FALSE.departments": []string{
			"slug", "name", "__typename"}}

	generatedTreeL := BuildTreeUsingAliases(graphqlQueryL, Variables{
		"include_aggregations": false})

	assert.Equal(t, expectedTreeL, generatedTreeL)

	var graphqlQueryM string = `
		query ProductsSearchPage($include_aggregations: Boolean!) {
		  search {
		    aggregations @include(if: $include_aggregations) {
		      departments {
		        ...departmentAggregationFields
		        __typename
		      }
		    }
		    lorem @include(if: $blabla) {

		    }
		    some_field
		  }
		}

		fragment departmentAggregationFields on DepartmentAggregation {
		  slug
		  name
		  __typename
		}

	`
	expectedTreeM := map[string][]string{
		"":                    []string{"search"},
		"search":              []string{"aggregations", "lorem", "some_field"},
		"search.aggregations": []string{"departments"},
		"search.aggregations.departments": []string{
			"slug", "name", "__typename"}}

	generatedTreeM := BuildTreeUsingAliases(graphqlQueryM, Variables{
		"include_aggregations": true,
		"blabla":               "lorem"})

	assert.Equal(t, expectedTreeM, generatedTreeM)
}
