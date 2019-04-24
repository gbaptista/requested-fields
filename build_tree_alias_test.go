package fields

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildTreeAlias(t *testing.T) {
	var graphqlQueryA string = `
	  {
		  search {
		    filters
		  }

		  best: search {
		    connection
		  }
		}
	`

	expectedTreeA := map[string][]string{
		"":       []string{"search"},
		"search": []string{"filters", "connection"}}

	generatedTreeA := BuildTree(graphqlQueryA)

	assert.Equal(t, expectedTreeA, generatedTreeA)

	var graphqlQueryB string = `
	  {
		  search {
		    filters
		  }

		  best: search {
		    connection
		  }

		  worst:search {
		    term
		  }
		}
	`

	expectedTreeB := map[string][]string{
		"":       []string{"search", "best", "worst"},
		"best":   []string{"connection"},
		"search": []string{"filters"},
		"worst":  []string{"term"}}

	generatedTreeB := BuildTreeUsingAliases(graphqlQueryB)

	assert.Equal(t, expectedTreeB, generatedTreeB)

	var graphqlQueryC string = `
	{
	  user(id: 3) {
	    id
	    custom_name: name
	    birthday
	  }

	  custom_user: user(id: 4) {
	    id
	    name
	    age
	  }
	}
	`
	expectedTreeC := map[string][]string{
		"":            []string{"user", "custom_user"},
		"user":        []string{"id", "custom_name", "birthday"},
		"custom_user": []string{"id", "name", "age"}}

	generatedTreeC := BuildTreeUsingAliases(graphqlQueryC)

	assert.Equal(t, expectedTreeC, generatedTreeC)
}
