package test

import (
	"reflect"
	"testing"

	scyna_proto "github.com/scyna/core/proto/generated"
	"google.golang.org/protobuf/proto"
)

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func TestGetType(t *testing.T) {
	var m proto.Message
	r := &scyna_proto.Request{}
	m = r

	if reflect.TypeOf(m).Elem().Name() == "Request" {
		t.Log("OK:", reflect.TypeOf(m).Name())
	} else {
		t.Error("Not OK", reflect.TypeOf(m).Name())
	}

	if getType(m) == "*Request" {
		t.Log("OK:", getType(m))
	} else {
		t.Error("Not OK", getType(m))
	}
}
