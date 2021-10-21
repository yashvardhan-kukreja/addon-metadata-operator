package validate

import (
	"fmt"
	"strings"

	"github.com/mt-sre/addon-metadata-operator/pkg/utils"
)

var (
	equalLines = strings.Repeat("=", 6)

	success = utils.Green("Success")
	failed  = utils.Red("Failed")
	err     = utils.IntenselyBoldRed("Error")
)

func printHeading() {
	fmt.Printf("\n%sRUNNING VALIDATORS%s\n\n", equalLines, equalLines)
}

func printSuccessMessage(msg string) {
	fmt.Printf("\r%s ... %s", msg, success)
}

func printFailureMessage(msg string) {
	fmt.Printf("\r%s ... %s", msg, failed)
}

func printErrorMessage(msg string) {
	fmt.Printf("\r%s ... %s", msg, err)
}
