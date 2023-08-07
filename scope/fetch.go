package scope

import (
	"time"

	hillchartsapi "github.com/offerni/hill-charts-api"
)

type FetchResponse struct {
	Colour     string
	CreatedAt  time.Time
	ID         hillchartsapi.ScopeID
	ModifiedAt time.Time
	Name       string
	Progress   float32
	SquadID    hillchartsapi.SquadID
}
