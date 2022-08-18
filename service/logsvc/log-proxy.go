package logsvc

import (
	"fmt"

	"github.com/ahl5esoft/lite-go/contract"
)

type LogProxy struct {
	DebugAction   func(map[string]string)
	ErrorAction   func(map[string]string, error)
	InfoAction    func(map[string]string)
	WarningAction func(map[string]string)

	label map[string]string
}

func (m *LogProxy) AddLabel(key, f string, v ...interface{}) contract.ILog {
	if m.label == nil {
		m.label = map[string]string{}
	}

	if len(v) > 0 {
		f = fmt.Sprintf(f, v...)
	}
	m.label[key] = f
	return m
}

func (m *LogProxy) Debug() {
	if len(m.label) == 0 {
		return
	}

	m.DebugAction(m.label)
}

func (m *LogProxy) Error(err error) {
	if len(m.label) == 0 && err == nil {
		return
	}

	m.ErrorAction(m.label, err)
}

func (m *LogProxy) Info() {
	if len(m.label) == 0 {
		return
	}

	m.InfoAction(m.label)
}

func (m *LogProxy) Warning() {
	if len(m.label) == 0 {
		return
	}

	m.WarningAction(m.label)
}
