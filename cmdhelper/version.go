package cmdhelper

import "fmt"

// Return a formatted version string.
func Version(version string, iteration string) string {
	var result string
	if iteration == "0" {
		result = version
	} else {
		result = fmt.Sprintf("%s-%s", version, iteration)
	}

	return result
}
