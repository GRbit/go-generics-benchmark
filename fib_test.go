package main_test

import (
	"testing"
)

const (
	fl = int64(20)
)

func BenchmarkInterfaceF1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibonacciInterface(fl)
	}
}

func BenchmarkGenericF1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibonacciGeneric(fl)
	}
}

func BenchmarkNativeF1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibonacciInt64(fl)
	}
}

func BenchmarkInterfaceF2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibonacciInterface(fl)
	}
}

func BenchmarkGenericF2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibonacciGeneric(fl)
	}
}

func BenchmarkNativeF2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibonacciInt64(fl)
	}
}

func FibonacciGeneric[T int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64](a T) T {
	if a <= 1 {
		return a
	}
	return FibonacciGeneric(a-1) + FibonacciGeneric(a-2)
}

func FibonacciInt64(a int64) int64 {
	if a <= 1 {
		return a
	}
	return FibonacciInt64(a-1) + FibonacciInt64(a-2)
}

func FibonacciInterface(a interface{}) interface{} {
	switch x := a.(type) {
	case int:
		if x <= 1 {
			return x
		}
		return FibonacciInterface(x-1).(int) + FibonacciInterface(x-2).(int)
	case int16:
		if x <= 1 {
			return x
		}
		return FibonacciInterface(x-1).(int16) + FibonacciInterface(x-2).(int16)
	case int32:
		if x <= 1 {
			return x
		}
		return FibonacciInterface(x-1).(int32) + FibonacciInterface(x-2).(int32)
	case int64:
		if x <= 1 {
			return x
		}
		return FibonacciInterface(x-1).(int64) + FibonacciInterface(x-2).(int64)
	case uint:
		if x <= 1 {
			return x
		}
		return FibonacciInterface(x-1).(uint) + FibonacciInterface(x-2).(uint)
	case uint16:
		if x <= 1 {
			return x
		}
		return FibonacciInterface(x-1).(uint16) + FibonacciInterface(x-2).(uint16)
	case uint32:
		if x <= 1 {
			return x
		}
		return FibonacciInterface(x-1).(uint32) + FibonacciInterface(x-2).(uint32)
	case uint64:
		if x <= 1 {
			return x
		}
		return FibonacciInterface(x-1).(uint64) + FibonacciInterface(x-2).(uint64)
	default:
		panic("only int and uint types are supported")
	}

	return nil
}
