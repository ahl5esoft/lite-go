package contract

// 数据查询
type IDbQuery interface {
	// 数量
	Count() (int64, error)
	// 排序(正序)
	Order(fields ...string) IDbQuery
	// 排序(倒序)
	OrderByDesc(fields ...string) IDbQuery
	// 跳过
	Skip(v int) IDbQuery
	// 限制
	Take(v int) IDbQuery
	// 数组结果
	ToArray(dst any) error
	// 条件
	Where(args ...any) IDbQuery
}
