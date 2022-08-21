package message

type Route struct {
	Api      string `uri:"api"`
	Endpoint string `uri:"endpoint"`
}
