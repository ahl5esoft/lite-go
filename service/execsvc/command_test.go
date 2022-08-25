package execsvc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_command_Exec(t *testing.T) {
	t.Run("timeout", func(t *testing.T) {
		_, _, err := NewCommand("ping", []string{"baidu.com", "-t"}).SetExpires(time.Second * 5).Exec()
		assert.Equal(t, err, errExecTimeout)
	})
}
