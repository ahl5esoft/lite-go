package enumsvc

import (
	"testing"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_enumFactory_Build(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockEnum := contract.NewMockIEnum(ctrl)
		self := NewEnumFactory(map[string]func() contract.IEnum{
			"": func() contract.IEnum {
				return mockEnum
			},
		})

		res := self.Build("")
		assert.Equal(t, res, mockEnum)
	})
}
