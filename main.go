package main

import (
	"flag"

	"github.com/bb-ben/termdict/core"
)

func main() {
	// positional arguments
	var word = flag.String("find", "", "query dictionary")
	flag.Parse()
	if word != nil {
		ret := core.Query(*word)
		core.TerminalFormatPrint(*ret)
	}
}
