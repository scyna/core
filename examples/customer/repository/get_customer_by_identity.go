package repository

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/customer/model"
)

func (r *customerRepository) GetCustomerByIdentity(identity model.Identity) (*model.Customer, scyna.Error) {
	/*TODO*/
	return &model.Customer{}, nil
}
