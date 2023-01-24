package repository

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/customer/domain"
)

const CUSTOMER_TABLE = "ddd_ex.customer"

type customerRepository struct {
	LOG scyna.Logger
}

func NewRepository(LOG scyna.Logger) domain.IRepository {
	return &customerRepository{LOG: LOG}
}
