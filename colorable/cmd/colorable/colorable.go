package main

import (
	"github.com/wondenge/rangi/colorable"
	"io"
	"os"


)

func main() {
	io.Copy(colorable.NewColorableStdout(), os.Stdin)
}
