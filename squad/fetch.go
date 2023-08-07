package squad

import (
	"time"

	hillchartsapi "github.com/offerni/hill-charts-api"
	"github.com/offerni/hill-charts-api/scope"
)

type FetchResponse struct {
	CreatedAt        time.Time
	CurrentCycleName string
	ID               hillchartsapi.SquadID
	ModifiedAt       time.Time
	Name             string
	OrganizationID   hillchartsapi.OrganizationID
	Scopes           []*scope.FetchResponse
}
