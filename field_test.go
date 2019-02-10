package fields

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ResolverA struct {
	Field Field `graphql:"a"`
}

type ResolverB struct {
	Field Field `graphql:"b"`
}

type ResolverC struct {
	Field Field `graphql:"c"`
}

type ResolverD struct {
	Field Field `graphql:"d"`
}

func TestField(t *testing.T) {
	resolver_a := &ResolverA{}

	var empty_string_array []string

	assert.Equal(t, 0, resolver_a.Field.Depth)
	assert.Equal(t, empty_string_array, resolver_a.Field.ParentTree)

	resolver_b := &ResolverB{}
	resolver_b.Field.SetCustomName("my_b")
	resolver_b.Field.SetParent(resolver_a)

	assert.Equal(t, 1, resolver_b.Field.Depth)
	assert.Equal(t, []string{"a"}, resolver_b.Field.ParentTree)

	assert.Equal(t, []string{"a"}, resolver_b.Field.ParentTree)

	resolver_c := &ResolverC{}
	resolver_c.Field.SetParent(resolver_b)

	assert.Equal(t, 2, resolver_c.Field.Depth)
	assert.Equal(t, []string{"a", "my_b"}, resolver_c.Field.ParentTree)

	resolver_d := &ResolverD{}
	resolver_d.Field.SetParent(resolver_c)

	assert.Equal(t, 3, resolver_d.Field.Depth)
	assert.Equal(t, []string{"a", "my_b", "c"}, resolver_d.Field.ParentTree)
}
