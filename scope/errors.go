package scope

import "errors"

var (
	ErrNoColour          error = errors.New("Colour is Required")
	ErrNoID              error = errors.New("ID is Required")
	ErrNoName            error = errors.New("Name is Required")
	ErrNoOrganizationID  error = errors.New("OrganizationID is Required")
	ErrNoScopeRepository error = errors.New("ScopeRepository Is Required")
	ErrNoSquadID         error = errors.New("SquadID is Required")
)
