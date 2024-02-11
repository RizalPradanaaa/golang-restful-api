// go:build wireinject
//  +bulid wireinject

package simple

import "github.com/google/wire"

func InitializedService() *SimpleService {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil
}
