package ports

import "github.com/kaiquecaires/hexagon-finances-server/domain/entities"

type UserRepository interface {
	Save(user entities.User) (entities.User, error)
}
