package enumsvc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_enumFactory_Build(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		self := NewEnumFactory(map[string]func() any{
			"": func() any {
				return "ok"
			},
		})

		var res string
		err := self.Build("", &res)
		assert.Nil(t, err)
		assert.Equal(t, res, "ok")
	})
}
