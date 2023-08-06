package squad

import "errors"

var (
	ErrNoOrganizationRepository error = errors.New("OrganizationRepository Is Required")
)
