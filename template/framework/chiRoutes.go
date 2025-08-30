package framework

import (
	_ "embed"
	"github.com/mahibulhaque/gofast/template/advanced"
)

//go:embed files/routes/chi.go.tmpl
var chiRoutesTemplate []byte

//go:embed files/tests/chi.go.tmpl
var chiTestHandlerTemplate []byte

// ChiTemplates contains the methods used for building
// an app that uses [github.com/go-chi/chi]
type ChiTemplates struct{}

func (c ChiTemplates) Main() []byte {
	return mainTemplate
}

func (c ChiTemplates) Server() []byte {
	return standardServerTemplate
}

func (c ChiTemplates) Routes() []byte {
	return chiRoutesTemplate
}

func (e ChiTemplates) TestHandler() []byte {
	return chiTestHandlerTemplate
}

func (c ChiTemplates) WebsocketImports() []byte {
	return advanced.StdLibWebsocketTemplImportsTemplate()
}
