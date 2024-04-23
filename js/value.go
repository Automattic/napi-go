package js

import (
	"github.com/Automattic/napi-go"
)

type Value struct {
	Env   Env
	Value napi.Value
}

func (v Value) GetEnv() Env {
	return v.Env
}
