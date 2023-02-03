package scyna_utils

import (
	"reflect"

	"google.golang.org/protobuf/proto"
)

func NewMessageForType[R proto.Message]() R {
	var msg R
	ref := reflect.New(reflect.TypeOf(msg).Elem())
	return ref.Interface().(R)
}
