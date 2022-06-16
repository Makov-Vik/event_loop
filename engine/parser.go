package engine

import (
	"strings"
)

func Parse(in string) Command {
	fields := strings.Fields(in)

	if len(fields) < 2 {
		return NewPrintCommand("SYNTAX ERROR: Not Enough Parameters")
	}

	name := fields[0]
	args := fields[1:]

	switch name {
	case "print":
		message := strings.Join(args, " ")
		return NewPrintCommand(message)
	case "split":
		arg1 := args[0]

		arg2 := args[1]

		return NewSplitCommand(arg1, arg2)
	default:
		return NewPrintCommand("ERROR: Unknown Command (" + in + ")")
	}
}