package service

import (
	"ex/hello/proto"

	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
)

func HelloHandler(ctx *scyna.Endpoint, request *proto.HelloRequest) scyna.Error {
	ctx.Logger.Info("Receive HelloRequest")

	if err := validateHelloRequest(request); err != nil {
		return scyna.REQUEST_INVALID
	}

	ctx.Response(&proto.HelloResponse{Content: "Hello " + request.Name})
	return scyna.OK
}

func validateHelloRequest(request *proto.HelloRequest) error {
	return validation.Validate(request.Name, validation.Required, validation.Length(3, 40))
}
