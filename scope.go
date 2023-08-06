package hillchartsapi

import "context"

type ScopeID string

const ScopeIDPrefix = "scp"

type Scope struct {
	Colour   string
	ID       ScopeID
	Name     string
	Progress float32
	SquadID  SquadID
}

type ScopeCreateOpts struct {
	AccountID AccountID
	Colour    string
	Name      string
	Progress  float32
}

type ScopeRepository interface {
	Create(ctx context.Context, opts ScopeCreateOpts) (*Scope, error)
}
