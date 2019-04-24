# requested-fields [![Build Status](https://travis-ci.org/gbaptista/requested-fields.svg?branch=master)](https://travis-ci.org/gbaptista/requested-fields) [![Maintainability](https://api.codeclimate.com/v1/badges/23ed3e32ab6688c28f6e/maintainability)](https://codeclimate.com/github/gbaptista/requested-fields/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/23ed3e32ab6688c28f6e/test_coverage)](https://codeclimate.com/github/gbaptista/requested-fields/test_coverage)

A simple library to extract requested fields from a GraphQL request. Supports fragments, aliases and all GraphQL features.

- [Quick Guide](#quick-guide)
  - [Query Tree](#query-tree)
  - [Resolvers Field](#resolvers-field)
  - [Requested Fields](#requested-fields)
  - [Accessing Fields on Deeper Levels](#accessing-fields-on-deeper-levels)
  - [Custom Names](#custom-names)
- [Complete Demo with Real Requests](#complete-demo-with-real-requests)
- [Known Issues](#known-issues)
  - [BuildTreeUsingAliases](#BuildTreeUsingAliases)

## Quick Guide

```go
import (
	fields "github.com/gbaptista/requested-fields"
)
```

Examples of how to use with [graphql-go](https://github.com/graph-gophers/graphql-go) and [chi](https://github.com/go-chi/chi):

### Query Tree

Create the query tree and pass it through context:

```go
query := `{ user { name } }`

ctx := context.WithValue(request.Context(),
		fields.ContextKey, fields.BuildTree(query))
```

### Resolvers Field

Include a `Field` field in all resolvers:
```go
type Query struct {
	Field fields.Field `graphql:"query"`
}

type UserResolver struct {
	Field fields.Field `graphql:"user"`
}
```

### Requested Fields
**Always** set parent resolver on all resolvers. To access the fields use the `fields.RequestedFor` function:
```go
func (queryResolver *Query) User(ctx context.Context) *UserResolver {
	userResolver := &UserResolver{}
	userResolver.Field.SetParent(queryResolver)
  
	log.Println(fmt.Sprintf(
		"Query.User Fields: %v", fields.RequestedFor(ctx, userResolver)))

	return userResolver
}

func (userResolver *UserResolver) Address(ctx context.Context) *AddressResolver {
	addressResolver := &AddressResolver{}
	addressResolver.Field.SetParent(userResolver)

	log.Println(fmt.Sprintf(
		"User.Address Fields: %v", fields.RequestedFor(ctx, addressResolver)))

	return addressResolver
}
```

### Accessing Fields on Deeper Levels
```go
fields.RequestedForAt(ctx, queryResolver, "user.address")
fields.RequestedForAt(ctx, userResolver, "address")
fields.RequestedForAt(ctx, userResolver, "address.country")
```

### Custom Names
For resources with different names:
```graphql
type Article {
  title: String
  author: User
}
```

Use the `SetCustomName` function:
```go
func (articleResolver *ArticleResolver) Author(ctx context.Context) *UserResolver {
	authorResolver := &UserResolver{}

	authorResolver.Field.SetCustomName("author")
	authorResolver.Field.SetParent(articleResolver)

	return authorResolver
}
```

## Complete Demo with Real Requests

For a complete use example go to [gbaptista/requested-fields-demo](https://github.com/gbaptista/requested-fields-demo).

## Known Issues

When aliases for the same resource are used at the same level:
```graphql
{
  user(id: 3) {
    id
    name
    birthday
  }

  custom_user: user(id: 4) {
    id
    name
    age
  }
}
```

The requested fields will be the fields of all equal resources at the same level:
```golang
[]string{"id", "name", "birthday", "age"}
```

an alternative to this behavior is to use the `BuildTreeUsingAliases` function:

### BuildTreeUsingAliases

```graphql
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
```

The requested fields will be:
```golang
map[string][]string{
    "": []string{"user", "custom_user"},
    "user": []string{"id", "custom_name", "birthday"},
    "custom_user": []string{"id", "name", "age"}}
```
