package squad

import hillchartsapi "github.com/offerni/hill-charts-api"

type FetchResponse struct {
	CurrentCycleName string
	ID               hillchartsapi.SquadID
	Name             string
	OrganizationID   hillchartsapi.OrganizationID
	Scopes           []*hillchartsapi.Scope
}
