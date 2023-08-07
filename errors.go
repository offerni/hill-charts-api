package hillchartsapi

import "errors"

var (
	ErrNoAccountID      error = errors.New("AccountID is required")
	ErrNoDB             error = errors.New("No Database Found")
	ErrNoOrganizationID error = errors.New("OrganizationID Is Required")
	ErrNoSquadID        error = errors.New("SquadID Is Required")
	ErrNoUserID         error = errors.New("UserID Is Required")
)
