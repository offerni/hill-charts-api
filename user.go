package hillchartsapi

type UserID string

type User struct {
	AccountID      AccountID
	ID             UserID
	Name           string
	OrganizationID string
}
