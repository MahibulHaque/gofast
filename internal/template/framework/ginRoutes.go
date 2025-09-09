package framework

import (
	_ "embed"

	"github.com/mahibulhaque/gofast/internal/template/advanced"
)

//go:embed files/routes/gin.go.tmpl
var ginRoutesTemplate []byte

// GinTemplates contains the methods used for building
// an app that uses [github.com/gin-gonic/gin]
type GinTemplates struct{}

func (g GinTemplates) Main() []byte {
	return mainTemplate
}

func (g GinTemplates) Server() []byte {
	return standardServerTemplate
}

func (g GinTemplates) Routes() []byte {
	return ginRoutesTemplate
}

func (g GinTemplates) WebsocketImports() []byte {
	return advanced.StdLibWebsocketTemplImportsTemplate()
}

func (g GinTemplates) RequestPackage() []byte {
	return ginRequestPackageTemplate
}

func (g GinTemplates) ResponsePackage() []byte {
	return ginResponsePackageTemplate
}
