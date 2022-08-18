package mongosvc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_unitOfWork_registerAdd(t *testing.T) {
	t.Run("未提交", func(t *testing.T) {
		entry := testModel{
			ID:   "id",
			Name: "add",
			Age:  1,
		}
		uow := newUnitOfWork(pool)
		uow.RegisterAdd(entry)
		assert.Len(t, uow.writeModel, 1)
		assert.Len(t, uow.writeModel["user"], 1)
	})

	t.Run("提交", func(t *testing.T) {
		entry := testModel{
			ID:   "id",
			Name: "add",
			Age:  1,
		}
		uow := newUnitOfWork(pool)
		uow.RegisterAdd(entry)
		err := uow.Commit()
		assert.NoError(t, err)

		_, db, err := pool.GetClientAndDb()
		assert.NoError(t, err)

		defer db.Drop(pool.Ctx)

		name, err := testModelMetadata.GetTableName()
		assert.NoError(t, err)

		cur, err := db.Collection(name).Find(pool.Ctx, bson.D{})
		assert.NoError(t, err)

		entries := make([]testModel, 0)
		for cur.Next(pool.Ctx) {
			var temp testModel
			err = cur.Decode(&temp)
			assert.NoError(t, err)

			entries = append(entries, temp)
		}

		assert.EqualValues(
			t,
			entries,
			[]testModel{entry},
		)
	})
}

func Test_unitOfWork_registerRemove(t *testing.T) {
	t.Run("未提交", func(t *testing.T) {
		entry := testModel{
			ID:   "id",
			Name: "remove",
			Age:  1,
		}
		uow := newUnitOfWork(pool)
		uow.RegisterRemove(entry)
		assert.Len(t, uow.writeModel, 1)
		assert.Len(t, uow.writeModel["user"], 1)
	})

	t.Run("提交", func(t *testing.T) {
		entry := testModel{
			ID:   "id",
			Name: "remove",
			Age:  1,
		}
		uow := newUnitOfWork(pool)
		uow.RegisterAdd(entry)
		uow.RegisterRemove(entry)
		err := uow.Commit()
		assert.NoError(t, err)

		_, db, err := pool.GetClientAndDb()
		assert.NoError(t, err)

		defer db.Drop(pool.Ctx)

		name, err := testModelMetadata.GetTableName()
		assert.NoError(t, err)

		cur, err := db.Collection(name).Find(pool.Ctx, bson.D{})
		assert.NoError(t, err)

		entries := make([]testModel, 0)
		for cur.Next(pool.Ctx) {
			var temp testModel
			err = cur.Decode(&temp)
			assert.NoError(t, err)

			entries = append(entries, temp)
		}

		assert.Len(t, entries, 0)
	})
}

func Test_unitOfWork_registerSave(t *testing.T) {
	t.Run("未提交", func(t *testing.T) {
		uow := newUnitOfWork(pool)
		entry := testModel{
			ID:   "id-2",
			Name: "save",
			Age:  2,
		}
		uow.RegisterSave(entry)
		assert.Len(t, uow.writeModel, 1)
		assert.Len(t, uow.writeModel["user"], 1)
	})

	t.Run("提交", func(t *testing.T) {
		uow := newUnitOfWork(pool)
		uow.RegisterAdd(testModel{
			ID:   "id-2",
			Name: "add",
			Age:  1,
		})
		entry := testModel{
			ID:   "id-2",
			Name: "save",
			Age:  2,
		}
		uow.RegisterSave(entry)
		err := uow.Commit()
		assert.NoError(t, err)

		_, db, err := pool.GetClientAndDb()
		assert.NoError(t, err)

		defer db.Drop(pool.Ctx)

		name, err := testModelMetadata.GetTableName()
		assert.NoError(t, err)

		cur, err := db.Collection(name).Find(pool.Ctx, bson.D{})
		assert.NoError(t, err)

		entries := make([]testModel, 0)
		for cur.Next(pool.Ctx) {
			var temp testModel
			err = cur.Decode(&temp)
			assert.NoError(t, err)

			entries = append(entries, temp)
		}

		assert.EqualValues(
			t,
			entries,
			[]testModel{entry},
		)
	})
}
