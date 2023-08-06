package models

import hillchartsapi "github.com/offerni/hill-charts-api"

type Scope struct {
	AccountID string  `firestore:"account_id"`
	Colour    string  `firestore:"colour"`
	ID        string  `firestore:"id"`
	Name      string  `firestore:"name"`
	Progress  float32 `firestore:"progress"`
	SquadID   string  `firestore:"squad_id"`
}

func ConvertScopeModelToDomain(s *Scope) *hillchartsapi.Scope {
	return &hillchartsapi.Scope{
		AccountID: hillchartsapi.AccountID(s.AccountID),
		Colour:    s.Colour,
		ID:        hillchartsapi.ScopeID(s.ID),
		Name:      s.Name,
		Progress:  s.Progress,
		SquadID:   hillchartsapi.SquadID(s.SquadID),
	}
}

func ConvertScopeDomainToModel(s *hillchartsapi.Scope) *Scope {
	return &Scope{
		AccountID: string(s.AccountID),
		Colour:    s.Colour,
		ID:        string(s.ID),
		Name:      s.Name,
		Progress:  s.Progress,
		SquadID:   string(s.SquadID),
	}
}
