package entities

import (
	"errors"

	"github.com/kaiquecaires/hexagon-finances-server/domain/value_objects"
)

type User struct {
	Id         string              `json:"id"`
	Email      value_objects.Email `json:"email"`
	Name       string              `json:"name"`
	SocialName string              `json:"social_name"`
}

func CreateUser(email string, name string, socialName string) (User, error) {
	if email == "" {
		return User{}, errors.New("email is required!")
	}

	if name == "" {
		return User{}, errors.New("name is required!")
	}

	newEmail, err := value_objects.NewEmail(email)

	if err != nil {
		return User{}, err
	}

	return User{
		Id:         "",
		Name:       name,
		Email:      newEmail,
		SocialName: socialName,
	}, nil
}
