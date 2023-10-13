//go:build tools
// +build tools

package graphql_hackernews

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)
