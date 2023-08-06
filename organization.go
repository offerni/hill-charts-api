package hillchartsapi

import "github.com/labstack/echo"

type OrganizationID string

type Organization struct {
	AccountID AccountID
	ID        OrganizationID
	Name      string
}

type OrganizationCreateOpts struct {
	AccountID AccountID
	Name      string
}

type OrganizationRepository interface {
	Create(ctx echo.Context, opts OrganizationCreateOpts) (*Organization, error)
}
