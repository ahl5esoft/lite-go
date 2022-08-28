package genericsvc

import (
	"testing"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/stretchr/testify/assert"
)

func Test_factory_Build(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		res := NewFactory[contract.ILogFactory](func() contract.ILog {
			return nil
		}).Build()
		assert.Nil(t, res)
	})
}
