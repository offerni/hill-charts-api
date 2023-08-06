package hillchartsapi

import "context"

type UserID string

type User struct {
	AccountID      AccountID
	ID             UserID
	Name           string
	OrganizationID OrganizationID
}

type UserFindOpts struct {
	AccountID      AccountID
	ID             UserID
	Name           string
	OrganizationID OrganizationID
}

type UserRepository interface {
	Find(ctx context.Context, opts UserFindOpts) (*User, error)
}
