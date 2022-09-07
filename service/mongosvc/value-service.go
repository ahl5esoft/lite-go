package mongosvc

import (
	"github.com/ahl5esoft/lite-go/contract"
	mcontract "github.com/ahl5esoft/lite-go/model/contract"
	"github.com/ahl5esoft/lite-go/model/message"
	"go.mongodb.org/mongo-driver/bson"
)

// 创建数值服务
func NewValueService[T mcontract.IValue, TLog mcontract.IValueLog, TValueType mcontract.IValueType](
	dbFactory contract.IDbFactory,
	enumFactory contract.IEnumFactory,
	nowTime contract.INowTime,
	stringGenerator contract.IStringGenerator,
	valueInterceptorFactory contract.IValueInterceptorFactory,
	filter bson.M,
	createEntryFunc func() T,
	createLogEntryFunc func(int, int64, string) TLog,
) contract.IValueService {
	var valueService contract.IValueService
	valueService = &contract.ValueServiceBase[T]{
		FindEntriesFunc: func() (res []T, err error) {
			err = dbFactory.Db(
				createEntryFunc(),
			).Query().Where(filter).ToArray(&res)
			return
		},
		UpdateFunc: func(uow contract.IUnitOfWork, source string, changeValues []message.ChangeValue, entries *[]T) (err error) {
			var entry T
			if len(*entries) == 0 {
				entry = createEntryFunc()
				dbFactory.Db(entry, uow).Add(entry)

				*entries = append(*entries, entry)
			} else {
				entry = (*entries)[0]
			}

			var logDb contract.IDbRepository
			var logEntry TLog
			now := nowTime.Unix()
			for _, r := range changeValues {
				var enumService contract.IEnumService[TValueType]
				if err = enumFactory.Build("ValueType", &enumService); err != nil {
					return
				}

				var valueTypeItem TValueType
				valueTypeItem, err = enumService.GetItem(func(cr TValueType) bool {
					return cr.GetValue() == r.ValueType
				})
				if err != nil {
					return
				}

				value := entry.GetValue()
				if _, ok := value[r.ValueType]; !ok {
					value[r.ValueType] = 0
				}

				if r.Source == "" {
					r.Source = source
				}

				logEntry = createLogEntryFunc(r.ValueType, value[r.ValueType], r.Source)
				if logDb == nil {
					logDb = dbFactory.Db(logEntry, uow)
				}

				if valueTypeItem.GetIsReplace() {
					value[r.ValueType] = 0
				} else if valueTypeItem.GetDailyTime() != 0 {
					ok := nowTime.IsSame(
						value[valueTypeItem.GetDailyTime()],
						"week",
					)
					if !ok {
						value[valueTypeItem.GetDailyTime()] = now
						value[r.ValueType] = 0
					}
				} else if r.Count == 0 {
					continue
				}

				var interceptor contract.IValueInterceptor
				if interceptor, err = valueInterceptorFactory.Build(r); err != nil {
					return
				}

				var ok bool
				if ok, err = interceptor.Before(uow, valueService, r); err != nil {
					return
				} else if ok {
					continue
				}

				value[r.ValueType] += r.Count
				logEntry.SetCount(value[r.ValueType])
				if logEntry.IsChange() {
					logEntry.SetID(
						stringGenerator.Generate(),
					)
					logDb.Add(logEntry)
				}

				if err = interceptor.After(uow, valueService, r); err != nil {
					return
				}
			}

			err = dbFactory.Db(entry, uow).Save(entry)
			return
		},
	}
	return valueService
}
