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

//go:embed files/react/src/lib/utils.ts.tmpl
var reactUtilsFile []byte

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

func ReactRootRouteFile() []byte {
	return reactRootRouteFile
}

func ReactUtilsFile() []byte {
	return reactUtilsFile
}
