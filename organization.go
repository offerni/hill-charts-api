package hillchartsapi

import "context"

type OrganizationID string

type Organization struct {
	AccountID AccountID
	ID        OrganizationID
	Name      string
}

type OrganizationFindOpts struct {
	AccountID AccountID
	ID        OrganizationID
	Name      string
}

type OrganizationRepository interface {
	Find(ctx context.Context, opts OrganizationFindOpts) (*Organization, error)
}
