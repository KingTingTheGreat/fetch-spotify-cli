package util

import (
	"os"
)

func EndWithErr(errorMessage string) {
	os.Stderr.WriteString("Error: " + errorMessage + "\n")

	os.Stdout.WriteString(OutputFileName())
	os.Exit(1)
}
