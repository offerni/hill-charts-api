package squad

import (
	hillchartsapi "github.com/offerni/hill-charts-api"
	"github.com/offerni/hill-charts-api/scope"
)

type FetchResponse struct {
	CurrentCycleName string
	ID               hillchartsapi.SquadID
	Name             string
	OrganizationID   hillchartsapi.OrganizationID
	Scopes           []*scope.FetchResponse
}
