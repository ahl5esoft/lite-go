package logsvc

import (
	"fmt"

	"github.com/ahl5esoft/lite-go/contract"
)

type LogProxy struct {
	DebugAction   func([][2]string)
	ErrorAction   func(error, [][2]string)
	InfoAction    func([][2]string)
	WarningAction func([][2]string)

	labels [][2]string
}

func (m *LogProxy) AddLabel(k, f string, v ...interface{}) contract.ILog {
	if m.labels == nil {
		m.labels = make([][2]string, 0)
	}

	if len(v) > 0 {
		f = fmt.Sprintf(f, v...)
	}
	m.labels = append(m.labels, [2]string{k, f})
	return m
}

func (m *LogProxy) Debug() {
	if len(m.labels) == 0 {
		return
	}

	m.DebugAction(m.labels)
}

func (m *LogProxy) Error(err error) {
	if err == nil && len(m.labels) == 0 {
		return
	}

	m.ErrorAction(err, m.labels)
}

func (m *LogProxy) Info() {
	if len(m.labels) == 0 {
		return
	}

	m.InfoAction(m.labels)
}

func (m *LogProxy) Warning() {
	if len(m.labels) == 0 {
		return
	}

	m.WarningAction(m.labels)
}
