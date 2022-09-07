package contract

import (
	"testing"

	"github.com/ahl5esoft/lite-go/model/enum"
	"github.com/stretchr/testify/assert"
)

func Test_EnumServiceBase_AllItems(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		self := &EnumServiceBase[enum.Item]{
			FindItemsFunc: func() ([]enum.Item, error) {
				return []enum.Item{
					{
						Value: 1,
					},
				}, nil
			},
		}

		res, err := self.AllItems()
		assert.Nil(t, err)
		assert.EqualValues(t, res, []enum.Item{
			{
				Value: 1,
			},
		})
	})
}

func Test_EnumServiceBase_GetItem(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		self := &EnumServiceBase[enum.Item]{
			FindItemsFunc: func() ([]enum.Item, error) {
				return []enum.Item{
					{
						Value: 1,
					},
				}, nil
			},
		}

		res, err := self.GetItem(func(r enum.Item) bool {
			return r.Value == 1
		})
		assert.Nil(t, err)
		assert.EqualValues(t, res, enum.Item{
			Value: 1,
		})
	})
}
