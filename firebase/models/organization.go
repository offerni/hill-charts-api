package models

import hillchartsapi "github.com/offerni/hill-charts-api"

type Organization struct {
	AccountID string `firestore:"account_id"`
	ID        string `firestore:"id"`
	Name      string `firestore:"name"`
}

func ConvertOrganizationModelToDomain(o *Organization) *hillchartsapi.Organization {
	return &hillchartsapi.Organization{
		AccountID: hillchartsapi.AccountID(o.AccountID),
		ID:        hillchartsapi.OrganizationID(o.ID),
		Name:      o.Name,
	}
}

func ConvertOrganizationDomainToModel(o *hillchartsapi.Organization) *Organization {
	return &Organization{
		AccountID: string(o.AccountID),
		ID:        string(o.ID),
		Name:      o.Name,
	}
}
