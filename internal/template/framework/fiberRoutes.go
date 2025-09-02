package framework

import (
	_ "embed"

	"github.com/mahibulhaque/gofast/internal/template/advanced"
)

//go:embed files/routes/fiber.go.tmpl
var fiberRoutesTemplate []byte

//go:embed files/server/fiber_server.go.tmpl
var fiberServerTemplate []byte

//go:embed files/main/fiber_main.go.tmpl
var fiberMainTemplate []byte

type FiberTemplates struct{}

func (f FiberTemplates) Main() []byte {
	return fiberMainTemplate
}
func (f FiberTemplates) Server() []byte {
	return fiberServerTemplate
}

func (f FiberTemplates) Routes() []byte {
	return fiberRoutesTemplate
}

func (f FiberTemplates) WebsocketImports() []byte {
	return advanced.FiberWebsocketTemplImportsTemplate()
}
