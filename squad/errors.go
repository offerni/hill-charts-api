package squad

import "errors"

var (
	ErrNoOrganizationRepository        error = errors.New("OrganizationRepository Is Required")
	ErrNoScopeRepository               error = errors.New("ScopeRepository Is Required")
	ErrNoSquadRepository               error = errors.New("SquadRepository Is Required")
	ErrNoUserRepository                error = errors.New("UserRepository Is Required")
	ErrOganizationNotFound             error = errors.New("Organization Not found")
	ErrUserDoesNotBelongToOrganization error = errors.New("User Does Not Belong To Organization")
	ErrUserNotFound                    error = errors.New("User Not found")
)
