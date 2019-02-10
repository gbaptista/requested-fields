package fields

import (
	"context"
	"strings"
)

func RequestedForAt(ctx context.Context, resolver interface{}, path_to_append string) []string {
	tree := ctx.Value("graphql_request_tree").(map[string][]string)

	name := NameFromResolver(resolver)
	field := FromResolver(resolver)

	path := append(field.ParentTree, name)

	// Remove the first "query" path
	_, path = path[0], path[1:]

	pathTree := strings.Join(path, ".")

	if path_to_append != "" {
		pathTree += "." + path_to_append
	}

	return tree[pathTree]
}

func RequestedFor(ctx context.Context, resolver interface{}) []string {
	return RequestedForAt(ctx, resolver, "")
}
