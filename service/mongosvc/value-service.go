package mongosvc

import (
	"reflect"

	"github.com/ahl5esoft/lite-go/contract"
	mcontract "github.com/ahl5esoft/lite-go/model/contract"
	"github.com/ahl5esoft/lite-go/model/message"
	"go.mongodb.org/mongo-driver/bson"
)

type valueService[T mcontract.IValue, TLog mcontract.IValueLog, TValueType mcontract.IValueType] struct {
	dbFactory               contract.IDbFactory
	enumFactory             contract.IEnumFactory
	nowTime                 contract.INowTime
	stringGenerator         contract.IStringGenerator
	valueInterceptorFactory contract.IValueInterceptorFactory
	filter                  bson.M
	createEntryFunc         func() T
	createLogEntryFunc      func(int, int64, string) TLog
	entries                 []T
}

func (m *valueService[T, TLog, TValueType]) CheckConditions(_ contract.IUnitOfWork, _ [][]mcontract.IValueCondition) (bool, error) {
	return false, nil
}

func (m *valueService[T, TLog, TValueType]) GetCount(uow contract.IUnitOfWork, valueType int) (res int64, err error) {
	var entry T
	if entry, err = m.getOrCreateEntry(uow); err != nil {
		return
	}

	if v, ok := entry.GetValue()[valueType]; ok {
		res = v
	}

	return
}

func (m *valueService[T, TLog, TValueType]) GetEntry(v any) (err error) {
	var entries []T
	if entries, err = m.findEntries(); err != nil {
		return
	} else if len(entries) > 0 {
		reflect.ValueOf(v).Elem().Set(
			reflect.ValueOf(entries[0]),
		)
	}

	return
}

func (m *valueService[T, TLog, TValueType]) Update(uow contract.IUnitOfWork, source string, changeValues []message.ChangeValue) (err error) {
	var entry T
	if entry, err = m.getOrCreateEntry(uow); err != nil {
		return
	}

	var logDb contract.IDbRepository
	var logEntry TLog
	for _, r := range changeValues {
		var enumService contract.IEnumService[TValueType]
		if err = m.enumFactory.Build("ValueType", &enumService); err != nil {
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

		logEntry = m.createLogEntryFunc(r.ValueType, value[r.ValueType], r.Source)
		if logDb == nil {
			logDb = m.dbFactory.Db(logEntry, uow)
		}

		if valueTypeItem.GetIsReplace() {
			value[r.ValueType] = 0
		} else if valueTypeItem.GetDailyTime() != 0 {
			ok := m.nowTime.IsSame(
				value[valueTypeItem.GetDailyTime()],
				"week",
			)
			if !ok {
				value[valueTypeItem.GetDailyTime()] = m.nowTime.Unix()
				value[r.ValueType] = 0
			}
		} else if r.Count == 0 {
			continue
		}

		var interceptor contract.IValueInterceptor
		if interceptor, err = m.valueInterceptorFactory.Build(r); err != nil {
			return
		}

		var ok bool
		if ok, err = interceptor.Before(uow, m, r); err != nil {
			return
		} else if ok {
			continue
		}

		value[r.ValueType] += r.Count
		logEntry.SetCount(value[r.ValueType])
		if logEntry.IsChange() {
			logEntry.SetID(
				m.stringGenerator.Generate(),
			)
			logDb.Add(logEntry)
		}

		if err = interceptor.After(uow, m, r); err != nil {
			return
		}
	}

	err = m.dbFactory.Db(entry, uow).Save(entry)
	return
}

func (m *valueService[T, TLog, TValueType]) getOrCreateEntry(uow contract.IUnitOfWork) (res T, err error) {
	var entries []T
	if entries, err = m.findEntries(); err != nil {
		return
	}

	if len(entries) == 0 {
		res = m.createEntryFunc()
		err = m.dbFactory.Db(res, uow).Add(res)

		m.entries = append(m.entries, res)
	} else {
		res = entries[0]
	}

	return
}

func (m *valueService[T, TLog, TValueType]) findEntries() (res []T, err error) {
	if m.entries == nil {
		err = m.dbFactory.Db(
			m.createEntryFunc(),
		).Query().Where(m.filter).ToArray(&(m.entries))
		if err != nil {
			return
		}
	}

	res = m.entries
	return
}

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
	return &valueService[T, TLog, TValueType]{
		createEntryFunc:         createEntryFunc,
		createLogEntryFunc:      createLogEntryFunc,
		dbFactory:               dbFactory,
		enumFactory:             enumFactory,
		filter:                  filter,
		nowTime:                 nowTime,
		stringGenerator:         stringGenerator,
		valueInterceptorFactory: valueInterceptorFactory,
	}
}
