package yamlsvc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	Int int
	Str string
}

func Test_configLoader_Load(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		self := &configLoader{
			doc: map[interface{}]interface{}{
				"Test": map[string]interface{}{
					"Int": 1,
					"Str": "s",
				},
			},
		}

		v := Test{}
		err := self.Load(&v)
		assert.Nil(t, err)
		assert.EqualValues(t, v, Test{
			Int: 1,
			Str: "s",
		})
	})
}
