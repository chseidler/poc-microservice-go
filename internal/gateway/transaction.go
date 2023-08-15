package gateway

import "github.com/chseidler/poc-microservice-go/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
