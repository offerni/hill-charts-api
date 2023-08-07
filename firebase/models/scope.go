package models

import (
	"time"

	hillchartsapi "github.com/offerni/hill-charts-api"
)

type Scope struct {
	AccountID  string     `firestore:"account_id"`
	Colour     string     `firestore:"colour"`
	CreatedAt  *time.Time `firestore:"created_at"`
	ID         string     `firestore:"id"`
	ModifiedAt *time.Time `firestore:"modified_at"`
	Name       string     `firestore:"name"`
	Progress   float32    `firestore:"progress"`
	SquadID    string     `firestore:"squad_id"`
}

func ConvertScopeModelToDomain(s *Scope) *hillchartsapi.Scope {
	return &hillchartsapi.Scope{
		AccountID:  hillchartsapi.AccountID(s.AccountID),
		Colour:     s.Colour,
		CreatedAt:  *s.CreatedAt,
		ID:         hillchartsapi.ScopeID(s.ID),
		ModifiedAt: *s.ModifiedAt,
		Name:       s.Name,
		Progress:   s.Progress,
		SquadID:    hillchartsapi.SquadID(s.SquadID),
	}
}

func ConvertScopeDomainToModel(s *hillchartsapi.Scope) *Scope {
	return &Scope{
		AccountID:  string(s.AccountID),
		Colour:     s.Colour,
		CreatedAt:  &s.CreatedAt,
		ID:         string(s.ID),
		ModifiedAt: &s.ModifiedAt,
		Name:       s.Name,
		Progress:   s.Progress,
		SquadID:    string(s.SquadID),
	}
}
