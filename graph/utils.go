package graph

import (
	"fmt"
	"strconv"
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

func float32ToStr(number float32) string {
	return fmt.Sprintf("%.2f", number)
}

func strToFloat32(str *string) *float32 {
	if str == nil || *str == "" {
		return nil
	}

	s, err := strconv.ParseFloat(*str, 32)
	if err != nil {
		return nil
	}

	newFloat := float32(s)
	return &newFloat
}

func buildScopesResponseFromList(scopes []*scope.FetchResponse) []*model.Scope {
	var scopesResp = []*model.Scope{}
	for _, scope := range scopes {
		scopesResp = append(scopesResp, &model.Scope{
			Colour:   scope.Colour,
			ID:       string(scope.ID),
			Name:     scope.Name,
			Progress: float32ToStr(scope.Progress),
			SquadID:  string(scope.SquadID),
		})
	}

	return scopesResp
}
