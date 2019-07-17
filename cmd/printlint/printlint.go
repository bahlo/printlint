package main

import (
	"github.com/bahlo/printlint/printcheck"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(printcheck.Analyzer)
}
