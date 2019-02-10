// Package fields implements a simple library to
// extract requested fields from a GraphQL request.
package fields

type key int

// ContextKey type to be used in Context for Request Tree.
const ContextKey key = iota
