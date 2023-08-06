package scope

import hillchartsapi "github.com/offerni/hill-charts-api"

type FetchResponse struct {
	Colour   string
	ID       hillchartsapi.ScopeID
	Name     string
	Progress float32
	SquadID  hillchartsapi.SquadID
}
