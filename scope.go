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
	AccountID      AccountID
	Colour         string
	Name           string
	OrganizationID OrganizationID
	Progress       float32
	SquadID        SquadID
}

type ScopeUpdateOpts struct {
	AccountID      AccountID
	Colour         *string
	ID             ScopeID
	Name           *string
	OrganizationID OrganizationID
	Progress       *float32
}

type ScopeList struct {
	*PaginatedList
	Scopes []*Scope
}

type ScopeFindOpts struct {
	AccountID      AccountID
	ID             ScopeID
	OrganizationID OrganizationID
}

type ScopeFindAllOpts struct {
	*PaginationOpts
	AccountID      AccountID
	OrganizationID OrganizationID
	SquadID        SquadID
}

type ScopeRepository interface {
	Create(ctx context.Context, opts ScopeCreateOpts) (*Scope, error)
	Find(ctx context.Context, opts ScopeFindOpts) (*Scope, error)
	FindAll(ctx context.Context, opts ScopeFindAllOpts) (*ScopeList, error)
	Update(ctx context.Context, opts ScopeUpdateOpts) error
}
