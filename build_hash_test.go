package fields

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type QueryVariables map[string]interface{}

func TestBuildHash(t *testing.T) {
	var graphqlQueryA string = `

		query productsSearch(

		  $products: ProductsSearchInput!,

		$products_search: ProductsSearchInput!) {

		  search(
		  	products: $products_search, products: $products
		 )


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

		    search {
		    aggregations @include(if: $var_tal) {
		      departments {
		      	id
		      	title
		        __typename
		      }
		    }
		    some_field
		  }

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

	generatedHashA, mappedVariablesA := BuildHash(
		graphqlQueryA, QueryVariables{}, true)

	assert.Equal(t, "3bd56bea76d31743f808b07e4c708415", generatedHashA)
	assert.Equal(t, map[string]string{}, mappedVariablesA)

}
