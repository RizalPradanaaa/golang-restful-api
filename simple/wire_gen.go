// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package simple

import (
	"github.com/google/wire"
)

// Injectors from injector.go:

func InitializedService(isError bool) (*SimpleService, error) {
	simpleRepository := NewSimpleRepository(isError)
	simpleService, err := NewSimpleService(simpleRepository)
	if err != nil {
		return nil, err
	}
	return simpleService, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	databasePostgre := NewDatabasePostgre()
	databaseMongoDb := NewDatabaseMongoDB()
	databaseRepository := NewDatabaseRepository(databasePostgre, databaseMongoDb)
	return databaseRepository
}

func InitializedFooBarService() *FooBarService {
	fooRepository := NewFooRepository()
	fooService := NewFooService(fooRepository)
	barRepository := NewBarRepository()
	barService := NewBarService(barRepository)
	fooBarService := NewFooBarService(fooService, barService)
	return fooBarService
}

func InitializedSayHelloervice() *HelloService {
	sayHelloImpl := NewSayHelloImpl()
	helloService := NewSayHelloService(sayHelloImpl)
	return helloService
}

func InitializedFooBar() *FooBar {
	foo := NewFoo()
	bar := NewBar()
	fooBar := &FooBar{
		Foo: foo,
		Bar: bar,
	}
	return fooBar
}

func InitializedFooBarUsingValue() *FooBar {
	foo := _wireFooValue
	bar := _wireBarValue
	fooBar := &FooBar{
		Foo: foo,
		Bar: bar,
	}
	return fooBar
}

var (
	_wireFooValue = &Foo{}
	_wireBarValue = &Bar{}
)

// injector.go:

var fooSet = wire.NewSet(NewFooRepository, NewFooService)

var barSet = wire.NewSet(NewBarRepository, NewBarService)

var helloSet = wire.NewSet(
	NewSayHelloImpl, wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

var FooBarSet = wire.NewSet(
	NewFoo,
	NewBar,
)

var fooBarValueSet = wire.NewSet(wire.Value(&Foo{}), wire.Value(&Bar{}))
