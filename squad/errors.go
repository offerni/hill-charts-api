package squad

import "errors"

var (
	ErrNoAccountID                     error = errors.New("AccountID Is Required")
	ErrNoID                            error = errors.New("ID Is Required")
	ErrNoName                          error = errors.New("Name Is Required")
	ErrNoOrganizationID                error = errors.New("OrganizationID Is Required")
	ErrNoOrganizationRepository        error = errors.New("OrganizationRepository Is Required")
	ErrNoScopeRepository               error = errors.New("ScopeRepository Is Required")
	ErrNoSquadRepository               error = errors.New("SquadRepository Is Required")
	ErrNoUserRepository                error = errors.New("UserRepository Is Required")
	ErrOganizationNotFound             error = errors.New("Organization Not found")
	ErrUserDoesNotBelongToOrganization error = errors.New("User Does Not Belong To Organization")
	ErrUserNotFound                    error = errors.New("User Not found")
)
