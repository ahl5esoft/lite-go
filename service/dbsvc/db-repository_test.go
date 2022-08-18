package dbsvc

import (
	"testing"

	"github.com/ahl5esoft/lite-go/contract"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testModel struct {
	ID string
}

func (m testModel) GetID() string {
	return m.ID
}

func Test_repositoryBase_Add(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(dbRepository)

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self.uow = mockUow

		entry := testModel{}
		mockUow.EXPECT().RegisterAdd(entry)

		mockUow.EXPECT().Commit().Return(nil)

		err := self.Add(entry)
		assert.NoError(t, err)
	})

	t.Run("事务", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(dbRepository)
		self.isTx = true

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self.uow = mockUow

		entry := testModel{}
		mockUow.EXPECT().RegisterAdd(entry)

		err := self.Add(entry)
		assert.NoError(t, err)
	})
}

func Test_repositoryBase_Remove(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(dbRepository)

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self.uow = mockUow

		entry := testModel{}
		mockUow.EXPECT().RegisterRemove(entry)

		mockUow.EXPECT().Commit().Return(nil)

		err := self.Remove(entry)
		assert.NoError(t, err)
	})

	t.Run("事务", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(dbRepository)
		self.isTx = true

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self.uow = mockUow

		entry := testModel{}
		mockUow.EXPECT().RegisterRemove(entry)

		err := self.Remove(entry)
		assert.NoError(t, err)
	})
}

func Test_repositoryBase_Save(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(dbRepository)

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self.uow = mockUow

		entry := testModel{}
		mockUow.EXPECT().RegisterSave(entry)

		mockUow.EXPECT().Commit().Return(nil)

		err := self.Save(entry)
		assert.NoError(t, err)
	})

	t.Run("事务", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(dbRepository)
		self.isTx = true

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self.uow = mockUow

		entry := testModel{}
		mockUow.EXPECT().RegisterSave(entry)

		err := self.Save(entry)
		assert.NoError(t, err)
	})
}
