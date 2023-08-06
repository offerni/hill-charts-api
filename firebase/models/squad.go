package models

import hillchartsapi "github.com/offerni/hill-charts-api"

type Squad struct {
	AccountID      string `firestore:"account_id"`
	ID             string `firestore:"id"`
	Name           string `firestore:"name"`
	OrganizationID string `firestore:"organization_id"`
}

func ConvertSquadModelToDomain(s *Squad) *hillchartsapi.Squad {
	return &hillchartsapi.Squad{
		ID:             hillchartsapi.SquadID(s.ID),
		Name:           s.Name,
		OrganizationID: hillchartsapi.OrganizationID(s.OrganizationID),
		AccountID:      hillchartsapi.AccountID(s.AccountID),
	}
}

func ConvertSquadDomainToModel(s *hillchartsapi.Squad) *Squad {
	return &Squad{
		ID:             string(s.ID),
		Name:           s.Name,
		OrganizationID: string(s.OrganizationID),
	}
}
