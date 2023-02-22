package yamlsvc

import (
	"os"
	"reflect"
	"sync"

	"github.com/ahl5esoft/lite-go/contract"
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v3"
)

var configLoaderMutex sync.Mutex

type configLoader struct {
	filePath string
	doc      map[interface{}]interface{}
}

func (m *configLoader) Load(v interface{}) (err error) {
	if m.doc == nil {
		configLoaderMutex.Lock()
		defer configLoaderMutex.Unlock()

		if m.doc == nil {
			var bf []byte
			if bf, err = os.ReadFile(m.filePath); err != nil {
				return
			}

			if err = yaml.Unmarshal(bf, &(m.doc)); err != nil {
				return
			}
		}
	}

	if cv, ok := m.doc[reflect.TypeOf(v).Elem().Name()]; ok {
		var bf []byte
		if bf, err = jsoniter.Marshal(cv); err != nil {
			return
		}

		err = jsoniter.Unmarshal(bf, v)
	}
	return
}

func NewConfigLoader(filePath contract.IFilePath, filename string) contract.IConfigLoader {
	return &configLoader{
		filePath: filePath.Join(
			filePath.Getwd(),
			filename,
		),
	}
}
