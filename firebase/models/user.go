package models

import hillchartsapi "github.com/offerni/hill-charts-api"

type User struct {
	AccountID      string `firestore:"account_id"`
	ID             string `firestore:"id"`
	Name           string `firestore:"name"`
	OrganizationID string `firestore:"organization_id"`
}

func ConvertUserModelToDomain(u *User) *hillchartsapi.User {
	return &hillchartsapi.User{
		AccountID:      hillchartsapi.AccountID(u.AccountID),
		ID:             hillchartsapi.UserID(u.ID),
		Name:           u.Name,
		OrganizationID: hillchartsapi.OrganizationID(u.OrganizationID),
	}
}

func ConvertUserDomainToModel(u *hillchartsapi.User) *User {
	return &User{
		AccountID:      string(u.AccountID),
		ID:             string(u.ID),
		Name:           u.Name,
		OrganizationID: string(u.OrganizationID),
	}
}
