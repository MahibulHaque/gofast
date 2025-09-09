package framework

import (
	_ "embed"
)

//go:embed files/request/standard_library_request.go.tmpl
var standardRequestPackageTemplate []byte

//go:embed files/response/standard_library_response.go.tmpl
var standardResponsePackageTemplate []byte

//go:embed files/request/echo_request.go.tmpl
var echoRequestPackageTemplate []byte

//go:embed files/response/echo_response.go.tmpl
var echoResponsePackageTemplate []byte

//go:embed files/request/fiber_request.go.tmpl
var fiberRequestPackageTemplate []byte

//go:embed files/response/fiber_response.go.tmpl
var fiberResponsePackageTemplate []byte

//go:embed files/request/gin_request.go.tmpl
var ginRequestPackageTemplate []byte

//go:embed files/response/gin_response.go.tmpl
var ginResponsePackageTemplate []byte

type UtilityPackage struct{}

func (up UtilityPackage) StandardRequestPackageTemplate() []byte {
	return standardRequestPackageTemplate
}

func (up UtilityPackage) StandardResponsePackageTemplate() []byte {
	return standardResponsePackageTemplate
}

func (up UtilityPackage) EchoRequestPackageTemplate() []byte {
	return echoRequestPackageTemplate
}

func (up UtilityPackage) EchoResponsePackageTemplate() []byte {
	return echoResponsePackageTemplate
}

func (up UtilityPackage) FiberRequestPackageTemplate() []byte {
	return fiberRequestPackageTemplate
}

func (up UtilityPackage) FiberResponsePackageTemplate() []byte {
	return fiberResponsePackageTemplate
}

func (up UtilityPackage) GinRequestPackageTemplate() []byte {
	return ginRequestPackageTemplate
}

func (up UtilityPackage) GinResponsePackageTemplate() []byte {
	return ginResponsePackageTemplate
}
