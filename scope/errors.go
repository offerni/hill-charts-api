package scope

import "errors"

var (
	ErrNoScopeRepository error = errors.New("ScopeRepository Is Required")
)
