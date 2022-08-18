package contract

import (
	"context"
	"reflect"
)

// 跟踪接口
type ITraceable interface {
	WithContext(ctx context.Context) reflect.Value
}
