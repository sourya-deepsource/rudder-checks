package main

import (
	"github.com/sourya-deepsource/rudder-checks/checks"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(checks.LogAnalyzer)
}
