package main

import (
	"os"

	"github.com/taylormonacelli/fulltexas"
)

func main() {
	code := fulltexas.Execute()
	os.Exit(code)
}
