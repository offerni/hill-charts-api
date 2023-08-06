package squad

import "errors"

var (
	ErrNoOrganizationRepository error = errors.New("OrganizationRepository Is Required")
	ErrNoSquadRepository        error = errors.New("SquadRepository Is Required")
)
