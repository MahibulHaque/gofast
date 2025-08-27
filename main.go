package main

import (
	"fmt"

	"github.com/mahibulhaque/gofast/tui/components/logo"
)

func main() {
	fmt.Println(logo.Render("v1.0.0", false, logo.DefaultOpts()))
}
