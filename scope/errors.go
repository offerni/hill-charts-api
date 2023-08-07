package scope

import "errors"

var (
	ErrNoColour                        error = errors.New("Colour is Required")
	ErrNoID                            error = errors.New("ID is Required")
	ErrNoName                          error = errors.New("Name is Required")
	ErrNoOrganizationID                error = errors.New("OrganizationID is Required")
	ErrNoOrganizationRepository        error = errors.New("ScopeRepository Is Required")
	ErrNoScopeRepository               error = errors.New("ScopeRepository Is Required")
	ErrNoSquadID                       error = errors.New("SquadID is Required")
	ErrNoUserRepository                error = errors.New("UserRepository Is Required")
	ErrOganizationNotFound             error = errors.New("Organization Not found")
	ErrUserDoesNotBelongToOrganization error = errors.New("User Does Not Belong To Organization")
	ErrUserNotFound                    error = errors.New("User Not Found")
)
