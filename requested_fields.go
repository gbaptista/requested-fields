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

	path_tree := strings.Join(path, ".")

	if path_to_append != "" {
		path_tree += "." + path_to_append
	}

	return tree[path_tree]
}

func RequestedFor(ctx context.Context, resolver interface{}) []string {
	return RequestedForAt(ctx, resolver, "")
}
