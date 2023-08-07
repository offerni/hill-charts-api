package graph

import (
	"github.com/offerni/hill-charts-api/scope"
	"github.com/offerni/hill-charts-api/squad"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	scopeSvc *scope.Service
	squadSvc *squad.Service
}

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type NewResolverOpts struct {
	ScopeService *scope.Service
	SquadService *squad.Service
}

func NewResolver(opts NewResolverOpts) *Resolver {
	return &Resolver{
		scopeSvc: opts.ScopeService,
		squadSvc: opts.SquadService,
	}
}
