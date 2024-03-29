package mongosvc

import (
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/contract"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbQuery struct {
	sorts      bson.D
	filter     bson.M
	pool       *dbPool
	model      *modelMetadata
	findOption *options.FindOptions
}

func (m dbQuery) Count() (int64, error) {
	defer m.Reset()

	c, err := m.pool.GetCollection(m.model)
	if err != nil {
		return 0, err
	}

	return c.CountDocuments(m.pool.Ctx, m.filter)
}

func (m *dbQuery) Order(fields ...string) contract.IDbQuery {
	m.sort(1, fields)
	return m
}

func (m *dbQuery) OrderByDesc(fields ...string) contract.IDbQuery {
	m.sort(-1, fields)
	return m
}

func (m *dbQuery) Reset() {
	m.filter = make(bson.M)
	m.findOption = options.Find()
	m.sorts = bson.D{}
}

func (m *dbQuery) Skip(v int) contract.IDbQuery {
	m.findOption = m.findOption.SetSkip(
		int64(v),
	)
	return m
}

func (m *dbQuery) Take(v int) contract.IDbQuery {
	m.findOption = m.findOption.SetLimit(
		int64(v),
	)
	return m
}

func (m *dbQuery) ToArray(dst interface{}) error {
	defer m.Reset()

	c, err := m.pool.GetCollection(m.model)
	if err != nil {
		return err
	}

	if len(m.sorts) > 0 {
		m.findOption = m.findOption.SetSort(m.sorts)
	}

	cur, err := c.Find(m.pool.Ctx, m.filter, m.findOption)
	if err != nil {
		return err
	}

	sliceType := reflect.SliceOf(
		m.model.GetType(),
	)
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)
	for cur.Next(m.pool.Ctx) {
		value := reflect.New(
			m.model.GetType(),
		)
		temp := value.Interface()
		cur.Decode(temp)
		sliceValue = reflect.Append(
			sliceValue,
			value.Elem(),
		)
	}
	reflect.ValueOf(dst).Elem().Set(sliceValue)
	return nil
}

func (m *dbQuery) Where(args ...interface{}) contract.IDbQuery {
	if len(args) == 0 {
		return m
	}

	if f, ok := args[0].(bson.M); ok {
		m.filter = f
	}
	return m
}

func (m *dbQuery) sort(flag int, fields []string) {
	underscore.Chain(fields).Each(func(r string, _ int) {
		m.sorts = append(m.sorts, bson.E{
			Key:   r,
			Value: flag,
		})
	})
}

func newDbQuery(pool *dbPool, model *modelMetadata) contract.IDbQuery {
	q := &dbQuery{
		model: model,
		pool:  pool,
	}
	q.Reset()
	return q
}
