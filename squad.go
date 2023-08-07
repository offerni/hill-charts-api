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
	FindAll(ctx context.Context, opts SquadFindAllOpts) (*SquadList, error)
}
