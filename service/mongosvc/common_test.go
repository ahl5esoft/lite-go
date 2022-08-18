package mongosvc

import (
	"reflect"
)

const (
	testUri    = "mongodb://localhost:27017"
	testDbName = "ahl5esoft-lite-go"
)

var (
	pool              = newDbPool(testDbName, testUri)
	testModelMetadata = getModelMetadata(testModelType)
	testModelType     = reflect.TypeOf(testModel{})
)

type testModel struct {
	ID   string `alias:"user" bson:"_id" db:"_id"`
	Name string `db:"name"`
	Age  int
}

func (m testModel) GetID() string {
	return m.ID
}
