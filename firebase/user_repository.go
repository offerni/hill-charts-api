package firebase

import (
	"cloud.google.com/go/firestore"
	hillchartsapi "github.com/offerni/hill-charts-api"
)

type userRepo struct {
	db *firestore.Client
}

func NewUserRepository(db *firestore.Client) (hillchartsapi.UserRepository, error) {
	if db == nil {
		return nil, ErrNoDB
	}

	return &userRepo{db}, nil
}
