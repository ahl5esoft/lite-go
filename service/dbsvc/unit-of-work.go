package dbsvc

import "fmt"

type UnitOfWork struct {
	CommitAction func() error

	afterAction map[string]func() error
}

func (m *UnitOfWork) Commit() (err error) {
	if err = m.CommitAction(); err != nil {
		return
	}

	for _, v := range m.afterAction {
		if err = v(); err != nil {
			return
		}
	}

	return
}

func (m *UnitOfWork) RegisterAfter(action func() error, key string) {
	if m.afterAction == nil {
		m.afterAction = make(map[string]func() error)
	}

	if key == "" {
		key = fmt.Sprintf(
			"key-%d",
			len(m.afterAction),
		)
	}

	m.afterAction[key] = action
}
