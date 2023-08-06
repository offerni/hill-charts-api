package hillchartsapi

import "github.com/labstack/echo"

type ScopeID string

const ScopeIDPrefix = "scp"

type Scope struct {
	Colour   string
	ID       ScopeID
	Name     string
	Progress float32
	SquadID  SquadID
}

type ScopeCreateOpts struct {
	AccountID AccountID
	Colour    string
	Name      string
	Progress  float32
}

type ScopeRepository interface {
	Create(ctx echo.Context, opts ScopeCreateOpts) (*Scope, error)
}
