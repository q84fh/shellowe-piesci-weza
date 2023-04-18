package validator

import (
	"regexp"
	"strconv"
)

// AreNumbers - checks if arguments are numbers
func AreNumbers(args []string) (isNum bool, invalidArg string) {
	for _, arg := range args {
		isNum = regexp.MustCompile(`\d`).MatchString(arg)
		if !isNum {
			return false, arg
		}
	}
	return true, ""
}

// AreInts - checks if arguments are integers
func AreInts(args []string) (isNum bool, invalidArg string) {
	for _, arg := range args {
		if _, err := strconv.Atoi(arg); err != nil {
			return false, arg
		}
	}
	return true, ""
}
