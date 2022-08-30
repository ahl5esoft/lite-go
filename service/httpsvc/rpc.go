package httpsvc

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/model/message"
	jsoniter "github.com/json-iterator/go"
)

type rpc struct {
	url    string
	body   any
	header map[string]string
}

func (m *rpc) Call(route string, res *message.ApiResponse) (err error) {
	routeParts := strings.Split(route, "/")
	if len(routeParts) == 3 {
		route = strings.Join([]string{
			"",
			routeParts[1],
			"ih",
			routeParts[2],
		}, "/")
	}
	var reader io.Reader
	if m.body != nil {
		var bf []byte
		if bf, err = jsoniter.Marshal(m.body); err != nil {
			return
		}

		reader = bytes.NewReader(bf)
	}
	var req *http.Request
	req, err = http.NewRequest(
		"post",
		fmt.Sprintf("%s/%s", m.url, route),
		reader,
	)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")

	var resp *http.Response
	if resp, err = (&http.Client{}).Do(req); err != nil {
		return
	}

	defer resp.Body.Close()

	var respBf []byte
	if respBf, err = io.ReadAll(resp.Body); err != nil {
		return
	}

	err = jsoniter.Unmarshal(respBf, res)
	return
}

func (m *rpc) SetBody(v any) contract.IRpc {
	m.body = v
	return m
}

func (m *rpc) SetHeader(v map[string]string) contract.IRpc {
	m.header = v
	return m
}

// 创建rpc
func NewRpc(url string) contract.IRpc {
	return &rpc{
		url: url,
	}
}
