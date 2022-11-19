package main

import (
	"flag"
	"fmt"

	"github.com/bb-ben/termdict/core"
)

func main() {
	// positional arguments
	var word = flag.String("find", "", "query dictionary")
	flag.Parse()
	if word != nil {
		ret := core.Query(*word)
		if ret != nil {
			core.TerminalFormatPrint(*ret)
		} else {
			fmt.Println("No word found")
		}
	}
}
