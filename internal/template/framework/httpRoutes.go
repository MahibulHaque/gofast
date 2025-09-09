package framework

import (
	_ "embed"

	"github.com/mahibulhaque/gofast/internal/template/advanced"
)

//go:embed files/routes/standard_library.go.tmpl
var standardRoutesTemplate []byte

//go:embed files/server/server.go.tmpl
var standardServerTemplate []byte

// StandardLibTemplate contains the methods used for building
// an app that uses [net/http]
type StandardLibTemplate struct{}

func (s StandardLibTemplate) Main() []byte {
	return mainTemplate
}

func (s StandardLibTemplate) Server() []byte {
	return standardServerTemplate
}

func (s StandardLibTemplate) Routes() []byte {
	return standardRoutesTemplate
}

func (s StandardLibTemplate) WebsocketImports() []byte {
	return advanced.StdLibWebsocketTemplImportsTemplate()
}

func (s StandardLibTemplate) RequestPackage() []byte {
	return standardRequestPackageTemplate
}

func (s StandardLibTemplate) ResponsePackage() []byte {
	return standardResponsePackageTemplate
}
