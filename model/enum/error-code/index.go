package errorcode

// 枚举值
type Value int

const (
	Null               Value = iota       // 无效
	API                Value = iota + 500 // api不存在
	Auth                                  // 认证失败
	Verify                                // 参数验证失败
	Timeout                               // 请求超时
	ValueTypeNotEnough                    // 数值不足
	Panic              Value = 599        // 异常错误码
)
