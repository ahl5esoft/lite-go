package contract

import contract "github.com/ahl5esoft/lite-go/model/contract"

// 数据仓库
type IDbRepository interface {
	Add(entry contract.IDbModel) error
	Query() IDbQuery
	Remove(entry contract.IDbModel) error
	Save(entry contract.IDbModel) error
}
