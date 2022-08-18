package mongosvc

import (
	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/model/contract"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type unitOfWork struct {
	dbPool     *dbPool
	writeModel map[string][]mongo.WriteModel
}

func (m *unitOfWork) Commit() error {
	if len(m.writeModel) == 0 {
		return nil
	}

	client, db, err := m.dbPool.GetClientAndDb()
	if err != nil {
		return err
	}

	return client.UseSession(m.dbPool.Ctx, func(ctx mongo.SessionContext) (err error) {
		defer func() {
			m.writeModel = make(map[string][]mongo.WriteModel)
		}()

		for k, v := range m.writeModel {
			if _, err = db.Collection(k).BulkWrite(ctx, v); err != nil {
				return
			}
		}

		return
	})
}

func (m *unitOfWork) RegisterAdd(entry contract.IDbModel) {
	model := getModelMetadata(entry)
	doc := make(bson.M)
	underscore.Chain(
		model.FindFields(),
	).Each(func(r *fieldMetadata, _ int) {
		doc[r.GetColumnName()] = r.GetValue(entry)
	})
	m.appendWriteModel(
		entry,
		mongo.NewInsertOneModel().SetDocument(doc),
	)
}

func (m *unitOfWork) RegisterRemove(entry contract.IDbModel) {
	m.appendWriteModel(
		entry,
		mongo.NewDeleteOneModel().SetFilter(bson.M{
			"_id": entry.GetID(),
		}),
	)
}

func (m *unitOfWork) RegisterSave(entry contract.IDbModel) {
	model := getModelMetadata(entry)
	writeModel := mongo.NewUpdateOneModel()
	doc := make(bson.M)
	underscore.Chain(
		model.FindFields(),
	).Each(func(r *fieldMetadata, _ int) {
		if r.GetTableName() != "" {
			writeModel.SetFilter(bson.M{
				"_id": r.GetValue(entry),
			})
		} else {
			doc[r.GetColumnName()] = r.GetValue(entry)
		}
	})
	writeModel.SetUpdate(bson.M{
		"$set": doc,
	})
	m.appendWriteModel(entry, writeModel)
}

func (m *unitOfWork) appendWriteModel(entry contract.IDbModel, writeModel mongo.WriteModel) {
	table, _ := getModelMetadata(entry).GetTableName()
	if _, ok := m.writeModel[table]; !ok {
		m.writeModel[table] = make([]mongo.WriteModel, 0)
	}

	m.writeModel[table] = append(m.writeModel[table], writeModel)
}

func newUnitOfWork(dbPool *dbPool) *unitOfWork {
	return &unitOfWork{
		dbPool:     dbPool,
		writeModel: make(map[string][]mongo.WriteModel),
	}
}
