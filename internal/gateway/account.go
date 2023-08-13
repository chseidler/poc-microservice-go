package gateway

import "github.com/chseidler/poc-microservice-go/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
}
