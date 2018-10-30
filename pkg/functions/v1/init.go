package v1

import (
	coreV1 "github.com/golang-boilerplate/pkg/core/v1"
)

//IFuncV1 is an interface for function struct
type IFuncV1 interface {
}

//FuncsV1 is a struct that hold all dependencies of function
type FuncsV1 struct {
	Core coreV1.ICoreV1
}

//InitFunc is a function to instantiate function
func InitFunc(core *coreV1.CoreV1) *FuncsV1 {
	return &FuncsV1{
		Core: core,
	}
}
