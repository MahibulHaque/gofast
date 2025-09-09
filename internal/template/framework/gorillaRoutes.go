package framework

import (
	_ "embed"

	"github.com/mahibulhaque/gofast/internal/template/advanced"
)

//go:embed files/routes/gorilla.go.tmpl
var gorillaRoutesTemplate []byte

// GorillaTemplates contains the methods used for building
// an app that uses [github.com/gorilla/mux]
type GorillaTemplates struct{}

func (g GorillaTemplates) Main() []byte {
	return mainTemplate
}

func (g GorillaTemplates) Server() []byte {
	return standardServerTemplate
}

func (g GorillaTemplates) Routes() []byte {
	return gorillaRoutesTemplate
}

func (g GorillaTemplates) WebsocketImports() []byte {
	return advanced.StdLibWebsocketTemplImportsTemplate()
}

func (g GorillaTemplates) RequestPackage() []byte {
	return standardRequestPackageTemplate
}

func (g GorillaTemplates) ResponsePackage() []byte {
	return standardResponsePackageTemplate
}
