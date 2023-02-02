package scyna_test

import (
	"bytes"
	"math"

	"google.golang.org/protobuf/proto"
	pref "google.golang.org/protobuf/reflect/protoreflect"
)

func matchMessage(x, y proto.Message) bool {
	mx := x.ProtoReflect()
	my := y.ProtoReflect()

	if mx.Descriptor() != my.Descriptor() {
		return false
	}

	equal := true
	my.Range(func(fd pref.FieldDescriptor, vy pref.Value) bool {
		if mx.Has(fd) {
			vx := mx.Get(fd)
			equal = equalField(fd, vx, vy)
		}
		return equal
	})
	return equal
}

// equalField compares two fields.
func equalField(fd pref.FieldDescriptor, x, y pref.Value) bool {
	switch {
	case fd.IsList():
		return equalList(fd, x.List(), y.List())
	case fd.IsMap():
		return equalMap(fd, x.Map(), y.Map())
	default:
		return equalValue(fd, x, y)
	}
}

// equalMap compares two maps.
func equalMap(fd pref.FieldDescriptor, x, y pref.Map) bool {
	if x.Len() != y.Len() {
		return false
	}
	equal := true
	x.Range(func(k pref.MapKey, vx pref.Value) bool {
		vy := y.Get(k)
		equal = y.Has(k) && equalValue(fd.MapValue(), vx, vy)
		return equal
	})
	return equal
}

// equalList compares two lists.
func equalList(fd pref.FieldDescriptor, x, y pref.List) bool {
	if x.Len() != y.Len() {
		return false
	}
	for i := x.Len() - 1; i >= 0; i-- {
		if !equalValue(fd, x.Get(i), y.Get(i)) {
			return false
		}
	}
	return true
}

// equalValue compares two singular values.
func equalValue(fd pref.FieldDescriptor, x, y pref.Value) bool {
	switch fd.Kind() {
	case pref.BoolKind:
		return x.Bool() == y.Bool()
	case pref.EnumKind:
		return x.Enum() == y.Enum()
	case pref.Int32Kind, pref.Sint32Kind,
		pref.Int64Kind, pref.Sint64Kind,
		pref.Sfixed32Kind, pref.Sfixed64Kind:
		return x.Int() == y.Int()
	case pref.Uint32Kind, pref.Uint64Kind,
		pref.Fixed32Kind, pref.Fixed64Kind:
		return x.Uint() == y.Uint()
	case pref.FloatKind, pref.DoubleKind:
		fx := x.Float()
		fy := y.Float()
		if math.IsNaN(fx) || math.IsNaN(fy) {
			return math.IsNaN(fx) && math.IsNaN(fy)
		}
		return fx == fy
	case pref.StringKind:
		return x.String() == y.String()
	case pref.BytesKind:
		return bytes.Equal(x.Bytes(), y.Bytes())
	// case pref.MessageKind, pref.GroupKind:
	// 	return matchMessage(x.Message(), y.Message())
	default:
		return x.Interface() == y.Interface()
	}
}
