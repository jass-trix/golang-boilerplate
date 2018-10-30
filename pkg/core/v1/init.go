package v1

//ICoreV1 is an interface for CoreV1 Struct
type ICoreV1 interface {
}

//CoreV1 is a struct that hold all dependencies of core module
type CoreV1 struct {
}

//InitCore is a function to instantiate core
func InitCore() *CoreV1 {
	return &CoreV1{}
}
