package args

import "fmt"

func ParseArgs(args []string) (string, string, error) {
	if len(args) != 2 {
		return "", "", fmt.Errorf("Needs launch URL and EXE Name")
	}
	return args[0], args[1], nil
}
