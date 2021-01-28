package resolver

// READ: https://gqlgen.com/getting-started/#finishing-touches

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/maapteh/graphql-golang/modules/zoo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver placeholder for dependencies
type Resolver struct {
}

func (r *Resolver) zoo() zoo.Zoo {
	return zoo.GetZoo(zoo.Artis)
}
