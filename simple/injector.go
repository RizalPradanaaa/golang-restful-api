// go:build wireinject
//  +bulid wireinject

package simple

import (
	"io"
	"os"

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

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
	// siapapun yang membutuhkan sayHello, akan diberi sayHelloImpl
)

// Kode Salah
// func InitializedSayHelloervice() *HelloService {
// 	wire.Build(NewSayHelloImpl, NewSayHelloService)
// 	return nil
// }

func InitializedSayHelloervice() *HelloService {
	wire.Build(helloSet, NewSayHelloService)
	return nil
}

var FooBarSet = wire.NewSet(
	NewFoo,
	NewBar,
)

func InitializedFooBar() *FooBar {
	wire.Build(
		FooBarSet,
		wire.Struct(new(FooBar), "Foo", "Bar"),
	)
	return nil
}

var fooBarValueSet = wire.NewSet(
	wire.Value(&Foo{}),
	wire.Value(&Bar{}),
)

func InitializedFooBarUsingValue() *FooBar {
	wire.Build(fooBarValueSet, wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}
