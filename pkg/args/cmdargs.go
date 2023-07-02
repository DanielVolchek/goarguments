package args

import (
	"errors"
	"os"
)

type CmdArg struct {
	arg   string
	value string
}

func LoadCmdArgs() ([]CmdArg, error) {

	args := os.Args[1:]

	parsed := []CmdArg{}

	current := CmdArg{}

	for _, arg := range args {
		if arg[:1] == "-" {
			if current.arg != "" {
				parsed = append(parsed, current)
			}

			current = CmdArg{arg: arg}
			continue
		}

		if current.value != "" {
			return nil, errors.New("Arguments formatted incorrectly")
		}

		current.value = arg
	}

	if current.arg != "" {
		parsed = append(parsed, current)
	}

	return parsed, nil
}
