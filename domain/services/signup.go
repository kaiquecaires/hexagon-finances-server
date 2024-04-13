package services

import (
	"github.com/kaiquecaires/hexagon-finances-server/domain/entities"
	"github.com/kaiquecaires/hexagon-finances-server/domain/ports"
)

type Signup struct {
	UserRepository ports.UserRepository
}

func (s Signup) Signup(user entities.User) (entities.User, error) {
	return s.UserRepository.Save(user)
}
