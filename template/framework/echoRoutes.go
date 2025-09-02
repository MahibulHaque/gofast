package framework

import (
	_ "embed"
	"github.com/mahibulhaque/gofast/template/advanced"
)

//go:embed files/routes/echo.go.tmpl
var echoRoutesTemplate []byte


// EchoTemplates contains the methods used for building
// an app that uses [github.com/labstack/echo]
type EchoTemplates struct{}

func (e EchoTemplates) Main() []byte {
	return mainTemplate
}
func (e EchoTemplates) Server() []byte {
	return standardServerTemplate
}

func (e EchoTemplates) Routes() []byte {
	return echoRoutesTemplate
}


func (e EchoTemplates) WebsocketImports() []byte {
	return advanced.StdLibWebsocketTemplImportsTemplate()
}
