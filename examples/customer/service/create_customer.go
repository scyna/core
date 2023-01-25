package service

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/customer/domain"
	"github.com/scyna/core/examples/customer/model"
	proto "github.com/scyna/core/examples/customer/proto/generated"
)

func CreateCustomerHandler(ctx *scyna.Endpoint, request *proto.CreateCustomerRequest) scyna.Error {
	var ret scyna.Error
	repository := domain.LoadRepository(ctx.Logger)
	customer := model.Customer{ID: model.CustomerID(domain.OneID.Next())}

	if customer.Identity, ret = model.NewIdentity(request.IDType, request.IDNumber); ret != nil {
		return ret
	}

	if ret = domain.IdentityExists(repository, customer.Identity); ret != nil {
		return ret
	}

	if ret := repository.CreateCustomerProfile(&customer); ret != nil {
		return ret
	}

	ctx.Response(&proto.CreateCustomerResponse{OneID: string(customer.ID)})

	return scyna.OK
}
