package main_test

import (
	"reflect"
	"testing"
)

const (
	cl = 1000
)

func BenchmarkReflectC1(b *testing.B) {
	s := make([]int, cl)
	for i := 0; i < cl; i++ {
		s[i] = i
	}

	for n := 0; n < b.N; n++ {
		ContainsReflect(s, cl-1)
	}
}

func BenchmarkGenericC1(b *testing.B) {
	s := make([]int, cl)
	for i := 0; i < cl; i++ {
		s[i] = i
	}

	for n := 0; n < b.N; n++ {
		ContainsGeneric(s, cl-1)
	}
}

func BenchmarkNativeC1(b *testing.B) {
	s := make([]int, cl)
	for i := 0; i < cl; i++ {
		s[i] = i
	}

	for n := 0; n < b.N; n++ {
		ContainsGeneric(s, cl-1)
	}
}

func BenchmarkReflectC2(b *testing.B) {
	s := make([]int, cl)
	for i := 0; i < cl; i++ {
		s[i] = i
	}

	for n := 0; n < b.N; n++ {
		ContainsReflect(s, cl-1)
	}
}

func BenchmarkGenericC2(b *testing.B) {
	s := make([]int, cl)
	for i := 0; i < cl; i++ {
		s[i] = i
	}

	for n := 0; n < b.N; n++ {
		ContainsGeneric(s, cl-1)
	}
}

func BenchmarkNativeC2(b *testing.B) {
	s := make([]int, cl)
	for i := 0; i < cl; i++ {
		s[i] = i
	}

	for n := 0; n < b.N; n++ {
		ContainsGeneric(s, cl-1)
	}
}

func ContainsGeneric[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsReflect(in interface{}, elem interface{}) bool {
	inValue := reflect.ValueOf(in)
	if inValue.Type().Kind() != reflect.Slice {
		panic("'in' is not a Slice")
	}
	for i := 0; i < inValue.Len(); i++ {
		if equal(elem, inValue.Index(i)) {
			return true
		}
	}
	return false
}

func equal(e interface{}, val reflect.Value) bool {
	if val.IsZero() {
		return val.Interface() == e
	}
	return reflect.DeepEqual(val.Interface(), e)
}

func ContainsNative(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
