package graph

//go:generate go run github.com/99designs/gqgen generate

import "github.com/kumin/GolangGraphQL/repos"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProductRepo repos.ProductRepo
}

func NewResolver(
	productRepo repos.ProductRepo,
) *Resolver {
	return &Resolver{
		ProductRepo: productRepo,
	}
}
