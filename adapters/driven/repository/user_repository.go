package repository

import "github.com/kaiquecaires/hexagon-finances-server/domain/entities"

type UserRepository struct{}

func (u UserRepository) Save(user entities.User) (entities.User, error) {
	return user, nil
}
