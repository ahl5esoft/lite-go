package contract

import (
	"testing"

	"github.com/ahl5esoft/lite-go/model/global"
	"github.com/ahl5esoft/lite-go/model/message"
	"github.com/stretchr/testify/assert"
)

func Test_ValueServiceBase_Update(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		self := &ValueServiceBase[global.UserValue]{
			FindEntriesFunc: func() ([]global.UserValue, error) {
				return make([]global.UserValue, 0), nil
			},
			UpdateFunc: func(_ IUnitOfWork, _ string, _ []message.ChangeValue, entries *[]global.UserValue) error {
				*entries = append(*entries, global.UserValue{
					Value: make(map[int]int64),
				})
				return nil
			},
		}

		err := self.Update(nil, "", nil)
		assert.Nil(t, err)

		assert.EqualValues(t, self.entries, []global.UserValue{
			{
				Value: make(map[int]int64),
			},
		})
	})
}
