package main

import (
	_ "github.com/proxystore11/Prx11/hiddify_extension"

	"github.com/hiddify/hiddify-core/extension/server"
)

func main() {
	server.StartTestExtensionServer()
}
