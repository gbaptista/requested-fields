package fields

import (
	"context"
	"strings"
)

// RequestedFor returns all requested fields for some Resolver.
func RequestedFor(ctx context.Context, resolver interface{}) []string {
	return RequestedForAt(ctx, resolver, "")
}

// RequestedForAt returns all requested fields for
//some path from a reference Resolver.
func RequestedForAt(ctx context.Context, resolver interface{}, pathToAppend string) []string {
	tree := ctx.Value(ContextKey).(map[string][]string)

	name := nameFromResolver(resolver)
	field := fromResolver(resolver)

	path := append(field.ParentTree, name)

	// Remove the first "query" path
	_, path = path[0], path[1:]

	pathTree := strings.Join(path, ".")

	if pathToAppend != "" {
		if pathTree == "" {
			pathTree = pathToAppend
		} else {
			pathTree += "." + pathToAppend
		}
	}

	return tree[pathTree]
}
