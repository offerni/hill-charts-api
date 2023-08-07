package graph

import (
	"fmt"
	"time"

	"github.com/offerni/hill-charts-api/graph/model"
	"github.com/offerni/hill-charts-api/scope"
	"github.com/offerni/hill-charts-api/squad"
)

func buildSquadsResponseFromList(squads []*squad.FetchResponse) []*model.Squad {
	var squadsResp = []*model.Squad{}
	for _, squad := range squads {

		squadsResp = append(squadsResp, &model.Squad{
			CreatedAt:        timeToStr(squad.CreatedAt),
			CurrentCycleName: squad.CurrentCycleName,
			ID:               string(squad.ID),
			ModifiedAt:       timeToStr(squad.ModifiedAt),
			Name:             squad.Name,
			Scope:            buildScopesResponseFromList(squad.Scopes),
		})
	}

	return squadsResp
}

func timeToStr(time time.Time) *string {
	formattedTime := time.Format("2006-01-02 15:04:05") // just a timestamp template

	return &formattedTime
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
