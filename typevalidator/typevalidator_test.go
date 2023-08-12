package typevalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	var value bool = true
	assert.Equal(t, "bool", GetTypeOf(value))
}

func TestInt(t *testing.T) {
	var value int = 1
	assert.Equal(t, "int", GetTypeOf(value))
}

func TestInt8(t *testing.T) {
	var value int8 = 1
	assert.Equal(t, "int8", GetTypeOf(value))
}

func TestInt16(t *testing.T) {
	var value int16 = 1
	assert.Equal(t, "int16", GetTypeOf(value))
}

func TestInt32(t *testing.T) {
	var value int32 = 1
	assert.Equal(t, "int32", GetTypeOf(value))
}

func TestInt64(t *testing.T) {
	var value int64 = 1
	assert.Equal(t, "int64", GetTypeOf(value))
}

func TestUint(t *testing.T) {
	var value uint = 1
	assert.Equal(t, "uint", GetTypeOf(value))
}

func TestUint8(t *testing.T) {
	var value uint8 = 1
	assert.Equal(t, "uint8", GetTypeOf(value))
}

func TestUint16(t *testing.T) {
	var value uint16 = 1
	assert.Equal(t, "uint16", GetTypeOf(value))
}

func TestUint32(t *testing.T) {
	var value uint32 = 1
	assert.Equal(t, "uint32", GetTypeOf(value))
}

func TestUint64(t *testing.T) {
	var value uint64 = 1
	assert.Equal(t, "uint64", GetTypeOf(value))
}

func TestUintptr(t *testing.T) {
	var value uintptr = 1
	assert.Equal(t, "uintptr", GetTypeOf(value))
}

func TestFloat32(t *testing.T) {
	var value float32 = 1.0
	assert.Equal(t, "float32", GetTypeOf(value))
}

func TestFloat64(t *testing.T) {
	var value float64 = 1.0
	assert.Equal(t, "float64", GetTypeOf(value))
}

func TestComplex64(t *testing.T) {
	var value complex64 = 1.0
	assert.Equal(t, "complex64", GetTypeOf(value))
}

func TestComplex128(t *testing.T) {
	var value complex128 = 1.0
	assert.Equal(t, "complex128", GetTypeOf(value))
}

func TestArray(t *testing.T) {
	var value [1]int
	assert.Equal(t, "[1]int", GetTypeOf(value))
}

func TestSlice(t *testing.T) {
	var value []int
	assert.Equal(t, "[]int", GetTypeOf(value))
}

func TestString(t *testing.T) {
	var value string = "test"
	assert.Equal(t, "string", GetTypeOf(value))
}

func TestMap(t *testing.T) {
	var value map[string]int
	assert.Equal(t, "map[string]int", GetTypeOf(value))
}

func TestStruct(t *testing.T) {
	var value struct{} = struct{}{}
	assert.Equal(t, "struct {}", GetTypeOf(value))
}

func TestPointer(t *testing.T) {
	var value *int
	assert.Equal(t, "*int", GetTypeOf(value))
}

func TestInterface(t *testing.T) {
	var value interface{} = 1
	assert.Equal(t, "int", GetTypeOf(value))
}

func TestFunc(t *testing.T) {
	var value func()
	assert.Equal(t, "func()", GetTypeOf(value))
}

func TestChan(t *testing.T) {
	var value chan int
	assert.Equal(t, "chan int", GetTypeOf(value))
}

func TestNil(t *testing.T) {
	var value interface{} = nil
	assert.Equal(t, "<nil>", GetTypeOf(value))
}

func TestRune(t *testing.T) {
	var value rune = 'a'
	assert.Equal(t, "int32", GetTypeOf(value))
}

func TestByte(t *testing.T) {
	var value byte = 'a'
	assert.Equal(t, "uint8", GetTypeOf(value))
}
