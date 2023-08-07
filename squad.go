package hillchartsapi

import (
	"context"
	"time"
)

type SquadID string

const SquadIDPrefix = "sqd"

type Squad struct {
	AccountID        AccountID
	CreatedAt        time.Time
	CurrentCycleName string
	ID               SquadID
	ModifiedAt       time.Time
	Name             string
	OrganizationID   OrganizationID
}

type SquadCreateOpts struct {
	AccountID        AccountID
	CurrentCycleName string
	Name             string
	OrganizationID   OrganizationID
}

type SquadFindOpts struct {
	AccountID      AccountID
	ID             SquadID
	OrganizationID OrganizationID
}

type SquadUpdateOpts struct {
	CurrentCycleName *string
	ID               SquadID
	Name             *string
	OrganizationID   OrganizationID
}

type SquadList struct {
	*PaginatedList
	Squads []*Squad
}

type SquadFindAllOpts struct {
	*PaginationOpts
	AccountID      AccountID
	OrganizationID OrganizationID
}

type SquadRepository interface {
	Create(ctx context.Context, opts SquadCreateOpts) (*Squad, error)
	Find(ctx context.Context, opts SquadFindOpts) (*Squad, error)
	FindAll(ctx context.Context, opts SquadFindAllOpts) (*SquadList, error)
	Update(ctx context.Context, opts SquadUpdateOpts) error
}
