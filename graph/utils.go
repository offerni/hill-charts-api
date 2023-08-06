package graph

import (
	"fmt"

	"github.com/offerni/hill-charts-api/graph/model"
	"github.com/offerni/hill-charts-api/scope"
	"github.com/offerni/hill-charts-api/squad"
)

func buildSquadsResponseFromList(squads []*squad.FetchResponse) []*model.Squad {
	var squadsResp = []*model.Squad{}
	for _, squad := range squads {
		squadsResp = append(squadsResp, &model.Squad{
			CurrentCycleName: squad.CurrentCycleName,
			ID:               string(squad.ID),
			Name:             squad.Name,
			Scope:            buildScopesResponseFromList(squad.Scopes),
		})
	}

	return squadsResp
}

func buildScopesResponseFromList(scopes []*scope.FetchResponse) []*model.Scope {
	var scopesResp = []*model.Scope{}
	for _, scope := range scopes {
		scopesResp = append(scopesResp, &model.Scope{
			Colour:   scope.Colour,
			ID:       string(scope.ID),
			Name:     scope.Name,
			Progress: fmt.Sprintf("%.2f", scope.Progress),
			SquadID:  string(scope.SquadID),
		})
	}

	return scopesResp
}
