package mongosvc

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbPoolMutex sync.Mutex

type dbPool struct {
	Ctx context.Context

	name   string
	client *mongo.Client
	db     *mongo.Database
	option *options.ClientOptions
}

func (m *dbPool) GetClientAndDb() (client *mongo.Client, db *mongo.Database, err error) {
	if m.client == nil || m.db == nil {
		dbPoolMutex.Lock()
		defer dbPoolMutex.Unlock()

		if m.client == nil {
			if m.client, err = mongo.Connect(m.Ctx, m.option); err != nil {
				return
			}
		}

		if m.db == nil {
			m.db = m.client.Database(m.name)
		}
	}

	client = m.client
	db = m.db
	return
}

func (m *dbPool) GetCollection(v any) (col *mongo.Collection, err error) {
	var name string
	if modelMetadata, ok := v.(*modelMetadata); ok {
		if name, err = modelMetadata.GetTableName(); err != nil {
			return
		}
	} else {
		name = v.(string)
	}

	var db *mongo.Database
	if _, db, err = m.GetClientAndDb(); err != nil {
		return
	}

	col = db.Collection(name)
	return
}

func (m *dbPool) WithContext(ctx context.Context) *dbPool {
	return &dbPool{
		Ctx:    ctx,
		client: m.client,
		db:     m.db,
		name:   m.name,
		option: m.option,
	}
}

func newDbPool(name, uri string) *dbPool {
	monitor := &monitor{}
	return &dbPool{
		Ctx:  context.Background(),
		name: name,
		option: options.Client().ApplyURI(uri).SetMonitor(&event.CommandMonitor{
			Started:   monitor.started,
			Succeeded: monitor.succeeded,
			Failed:    monitor.failed,
		}),
	}
}
