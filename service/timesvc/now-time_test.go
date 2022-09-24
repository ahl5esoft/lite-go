package timesvc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewNowTime_Unix(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		res := NewNowTime().Unix()
		assert.Equal(
			t,
			res,
			time.Now().Unix(),
		)
	})
}
