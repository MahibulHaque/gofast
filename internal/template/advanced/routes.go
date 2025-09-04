package advanced

import (
	_ "embed"
)

//go:embed files/websocket/imports/standard_library.tmpl
var stdLibWebsocketImports []byte

//go:embed files/websocket/imports/fiber.tmpl
var fiberWebsocketTemplImports []byte

//go:embed files/react/vite.config.ts.tmpl
var reactViteConfigFile []byte

//go:embed files/react/components.json.tmpl
var reactComponentsJsonFile []byte

//go:embed files/react/src/styles.css.tmpl
var reactStylesCssTemplate []byte

//go:embed files/react/src/main.tsx.tmpl
var reactMainFile []byte

//go:embed files/react/src/routes/root.tsx.tmpl
var reactRootRouteFile []byte

//go:embed files/react/src/routes/index.tsx.tmpl
var reactIndexRouteFile []byte

//go:embed files/react/src/routes/demo.tanstack-query.tsx.tmpl
var reactDemoTanstackQueryRouteFile []byte

//go:embed files/react/src/components/Header.tsx.tmpl
var reactHeaderComponentFile []byte

//go:embed files/react/src/lib/utils.ts.tmpl
var reactUtilsFile []byte

//go:embed files/react/package.json.tmpl
var reactPackageJsonFile []byte

//go:embed files/react/tsconfig.app.json.tmpl
var reactTsConfigAppJsonFile []byte

//go:embed files/react/tsconfig.json.tmpl
var reactTsConfigJsonFile []byte



func StdLibWebsocketTemplImportsTemplate() []byte {
	return stdLibWebsocketImports
}

func FiberWebsocketTemplImportsTemplate() []byte {
	return fiberWebsocketTemplImports
}

func ReactViteConfigFile() []byte {
	return reactViteConfigFile
}

func ReactComponentsJsonFile() []byte {
	return reactComponentsJsonFile
}

func ReactStylesCssFile() []byte {
	return reactStylesCssTemplate
}

func ReactMainFile() []byte {
	return reactMainFile
}

func ReactPackageJsonFile() []byte {
	return reactPackageJsonFile
}

func ReactTsConfigAppJsonFile() []byte {
	return reactTsConfigAppJsonFile
}

func ReactTsConfigJsonFile() []byte{
	return reactTsConfigJsonFile
}

func ReactRootRouteFile() []byte {
	return reactRootRouteFile
}

func ReactIndexRouteFile()[]byte{
	return reactIndexRouteFile
}

func ReactDemoTanstackQueryRouteFile()[]byte{
	return reactDemoTanstackQueryRouteFile
}

func ReactHeaderComponentFile()[]byte{
	return reactHeaderComponentFile
}

func ReactUtilsFile() []byte {
	return reactUtilsFile
}