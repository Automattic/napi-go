package napi

/*
#include <stdlib.h>
#include <node/node_api.h>
*/
import "C"

import (
	"unsafe"
)

func GetUndefined(env Env) (Value, Status) {
	var result Value
	status := Status(C.napi_get_undefined(
		C.napi_env(env),
		(*C.napi_value)(unsafe.Pointer(&result)),
	))
	return result, status
}

func GetNull(env Env) (Value, Status) {
	var result Value
	status := Status(C.napi_get_null(
		C.napi_env(env),
		(*C.napi_value)(unsafe.Pointer(&result)),
	))
	return result, status
}

func CreateObject(env Env) (Value, Status) {
	var result Value
	status := Status(C.napi_create_object(
		C.napi_env(env),
		(*C.napi_value)(unsafe.Pointer(&result)),
	))
	return result, status
}

func CreateArray(env Env) (Value, Status) {
	var result Value
	status := Status(C.napi_create_array(
		C.napi_env(env),
		(*C.napi_value)(unsafe.Pointer(&result)),
	))
	return result, status
}

func CreateArrayWithLength(env Env, length int) (Value, Status) {
	var result Value
	status := Status(C.napi_create_array_with_length(
		C.napi_env(env),
		C.size_t(length),
		(*C.napi_value)(unsafe.Pointer(&result)),
	))
	return result, status
}

func CreateStringUtf8(env Env, str string) (Value, Status) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	var result Value
	status := Status(C.napi_create_string_utf8(
		C.napi_env(env),
		cstr,
		C.size_t(len([]byte(str))), // must pass number of bytes
		(*C.napi_value)(unsafe.Pointer(&result)),
	))
	return result, status
}

func CreateError(env Env, code, msg Value) (Value, Status) {
	var result Value
	status := Status(C.napi_create_error(
		C.napi_env(env),
		C.napi_value(code),
		C.napi_value(msg),
		(*C.napi_value)(unsafe.Pointer(&result)),
	))
	return result, status
}

func Typeof(env Env, value Value) (ValueType, Status) {
	var result ValueType
	status := Status(C.napi_typeof(
		C.napi_env(env),
		C.napi_value(value),
		(*C.napi_valuetype)(unsafe.Pointer(&result)),
	))
	return result, status
}

func SetProperty(env Env, object, key, value Value) Status {
	return Status(C.napi_set_property(
		C.napi_env(env),
		C.napi_value(object),
		C.napi_value(key),
		C.napi_value(value),
	))
}

func SetElement(env Env, object Value, index int, value Value) Status {
	return Status(C.napi_set_element(
		C.napi_env(env),
		C.napi_value(object),
		C.uint32_t(index),
		C.napi_value(value),
	))
}