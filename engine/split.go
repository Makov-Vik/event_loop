package engine

import (
	"strings"
)

func NewSplitCommand(arg1 string, arg2 string) *splitCommand {
	return &splitCommand{
		arg1: arg1,
		arg2: arg2,
	}
}

type splitCommand struct {
	arg1, arg2 string
}

func (split *splitCommand) Execute(h Handler) {
	arrPart := strings.Split(split.arg1, split.arg2)
	res := strings.Join(arrPart, "\n")
	h.Post(NewPrintCommand(res))
}