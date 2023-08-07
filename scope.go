package hillchartsapi

import (
	"context"
	"time"
)

type ScopeID string

const ScopeIDPrefix = "scp"

type Scope struct {
	AccountID  AccountID
	Colour     string
	CreatedAt  time.Time
	ID         ScopeID
	ModifiedAt time.Time
	Name       string
	Progress   float32
	SquadID    SquadID
}

type ScopeCreateOpts struct {
	AccountID AccountID
	Colour    string
	Name      string
	Progress  float32
	SquadID   SquadID
}

type ScopeList struct {
	*PaginatedList
	Scopes []*Scope
}

type ScopeFindAllOpts struct {
	*PaginationOpts
	AccountID      AccountID
	OrganizationID OrganizationID
	SquadID        SquadID
}

type ScopeRepository interface {
	Create(ctx context.Context, opts ScopeCreateOpts) (*Scope, error)
	FindAll(ctx context.Context, opts ScopeFindAllOpts) (*ScopeList, error)
}
