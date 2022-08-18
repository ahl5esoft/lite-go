package mongosvc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_factory_Db(t *testing.T) {
	t.Run("add - 非事务", func(t *testing.T) {
		defer func() {
			_, db, err := pool.GetClientAndDb()
			assert.NoError(t, err)

			db.Drop(pool.Ctx)
		}()

		self := new(dbFactory)
		self.pool = pool

		entry := testModel{
			ID:   "id",
			Name: "1",
			Age:  1,
		}
		db := self.Db(entry)
		err := db.Add(entry)
		assert.NoError(t, err)

		var res []testModel
		err = db.Query().ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry},
		)
	})

	t.Run("add - 事务", func(t *testing.T) {
		defer func() {
			_, db, err := pool.GetClientAndDb()
			assert.NoError(t, err)

			db.Drop(pool.Ctx)
		}()

		self := new(dbFactory)
		self.pool = pool

		uow := self.Uow()

		entry := testModel{
			ID:   "id",
			Name: "1",
			Age:  1,
		}
		db := self.Db(entry, uow)
		err := db.Add(entry)
		assert.NoError(t, err)

		var res []testModel
		err = db.Query().ToArray(&res)
		assert.NoError(t, err)
		assert.Len(t, res, 0)
	})

	t.Run("remove - 非事务", func(t *testing.T) {
		defer func() {
			_, db, err := pool.GetClientAndDb()
			assert.NoError(t, err)

			db.Drop(pool.Ctx)
		}()

		self := new(dbFactory)
		self.pool = pool

		entry := testModel{
			ID:   "id",
			Name: "1",
			Age:  1,
		}
		db := self.Db(entry)
		err := db.Add(entry)
		assert.NoError(t, err)

		err = db.Remove(entry)
		assert.NoError(t, err)

		var res []testModel
		err = db.Query().ToArray(&res)
		assert.NoError(t, err)
		assert.Len(t, res, 0)
	})

	t.Run("remove - 事务", func(t *testing.T) {
		defer func() {
			_, db, err := pool.GetClientAndDb()
			assert.NoError(t, err)

			db.Drop(pool.Ctx)
		}()

		self := new(dbFactory)
		self.pool = pool

		entry := testModel{
			ID:   "id",
			Name: "1",
			Age:  1,
		}
		db := self.Db(entry)
		err := db.Add(entry)
		assert.NoError(t, err)

		err = self.Db(
			entry,
			self.Uow(),
		).Remove(entry)
		assert.NoError(t, err)

		var res []testModel
		err = db.Query().ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry},
		)
	})

	t.Run("save - 非事务", func(t *testing.T) {
		defer func() {
			_, db, err := pool.GetClientAndDb()
			assert.NoError(t, err)

			db.Drop(pool.Ctx)
		}()

		self := new(dbFactory)
		self.pool = pool

		entry := testModel{
			ID:   "id",
			Name: "1",
			Age:  1,
		}
		db := self.Db(entry)
		err := db.Add(entry)
		assert.NoError(t, err)

		modifiedEntry := testModel{
			ID:   entry.ID,
			Name: "11",
			Age:  11,
		}
		err = db.Save(modifiedEntry)
		assert.NoError(t, err)

		var res []testModel
		err = db.Query().ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{modifiedEntry},
		)
	})

	t.Run("save - 非事务", func(t *testing.T) {
		defer func() {
			_, db, err := pool.GetClientAndDb()
			assert.NoError(t, err)

			db.Drop(pool.Ctx)
		}()

		self := new(dbFactory)
		self.pool = pool

		entry := testModel{
			ID:   "id",
			Name: "1",
			Age:  1,
		}
		db := self.Db(entry)
		err := db.Add(entry)
		assert.NoError(t, err)

		modifiedEntry := testModel{
			ID:   entry.ID,
			Name: "11",
			Age:  11,
		}
		err = self.Db(
			modifiedEntry,
			self.Uow(),
		).Save(modifiedEntry)
		assert.NoError(t, err)

		var res []testModel
		err = db.Query().ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry},
		)
	})
}
