package models

import (
	"time"

	hillchartsapi "github.com/offerni/hill-charts-api"
)

type Squad struct {
	AccountID        string     `firestore:"account_id"`
	CreatedAt        *time.Time `firestore:"created_at"`
	CurrentCycleName string     `firestore:"current_cycle_name"`
	ID               string     `firestore:"id"`
	ModifiedAt       *time.Time `firestore:"modified_at"`
	Name             string     `firestore:"name"`
	OrganizationID   string     `firestore:"organization_id"`
}

func ConvertSquadModelToDomain(s *Squad) *hillchartsapi.Squad {
	return &hillchartsapi.Squad{
		AccountID:        hillchartsapi.AccountID(s.AccountID),
		CreatedAt:        *s.CreatedAt,
		CurrentCycleName: s.CurrentCycleName,
		ID:               hillchartsapi.SquadID(s.ID),
		ModifiedAt:       *s.ModifiedAt,
		Name:             s.Name,
		OrganizationID:   hillchartsapi.OrganizationID(s.OrganizationID),
	}
}

func ConvertSquadDomainToModel(s *hillchartsapi.Squad) *Squad {
	return &Squad{
		AccountID:        string(s.AccountID),
		CreatedAt:        &s.CreatedAt,
		CurrentCycleName: s.CurrentCycleName,
		ID:               string(s.ID),
		ModifiedAt:       &s.ModifiedAt,
		Name:             s.Name,
		OrganizationID:   string(s.OrganizationID),
	}
}
