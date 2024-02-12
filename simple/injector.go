// go:build wireinject
//  +bulid wireinject

package simple

import (
	"github.com/google/wire"
)

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}
func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(NewDatabasePostgre, NewDatabaseMongoDB, NewDatabaseRepository)
	return nil
}
