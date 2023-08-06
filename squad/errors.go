package squad

import "errors"

var (
	ErrNoOrganizationRepository error = errors.New("OrganizationRepository Is Required")
	ErrNoScopeRepository        error = errors.New("ScopeRepository Is Required")
	ErrNoSquadRepository        error = errors.New("SquadRepository Is Required")
)
