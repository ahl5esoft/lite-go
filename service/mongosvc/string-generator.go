package mongosvc

import (
	"github.com/ahl5esoft/lite-go/contract"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type stringGenerator struct{}

func (m stringGenerator) Generate() string {
	return primitive.NewObjectID().Hex()
}

// 创建字符串生成器
func NewStringGenerator() contract.IStringGenerator {
	return new(stringGenerator)
}
