package simple

type Database struct {
	Name string
}

type DatabasePostgre Database
type DatabaseMongoDb Database

type DatabaseRepository struct {
	*DatabasePostgre
	*DatabaseMongoDb
}

func NewDatabasePostgre() *DatabasePostgre {
	database := &Database{
		Name: "Database Postgre",
	}
	return (*DatabasePostgre)(database)
}

func NewDatabaseMongoDB() *DatabaseMongoDb {
	database := &Database{
		Name: "Database MongoDB",
	}
	return (*DatabaseMongoDb)(database)
}

func NewDatabaseRepository(databasePostgre *DatabasePostgre, databaseMongoDB *DatabaseMongoDb) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostgre: databasePostgre,
		DatabaseMongoDb: databaseMongoDB,
	}
}
