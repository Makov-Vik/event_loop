package main

import (
	"reflect"
	"testing"

	"github.com/Makov-Vik/event_loop/engine"
	"github.com/stretchr/testify/assert"
)


func TestParserSplit(t *testing.T) {
	commandString := "split a:bc:d:ef :"
	command := engine.Parse(commandString)
	examplSplit := engine.NewSplitCommand("a:bc:d:ef", ":")

	if assert.NotNil(t, command) {
		assert.Equal(t, command, examplSplit)
	}
}

func TestParserPrint(t *testing.T) {
	commandString := "print aloha"
	command := engine.Parse(commandString)
	examplePrint := engine.NewPrintCommand("aloha")

	if assert.NotNil(t, examplePrint) {
		assert.Equal(t, examplePrint, command)
	}
}

func TestParserDefault(t *testing.T) {
	commandString := ""
	command := engine.Parse(commandString)
	examplePrint := engine.NewPrintCommand("")

	if assert.NotNil(t, command) {
		assert.Equal(t, reflect.TypeOf(command), reflect.TypeOf(examplePrint))
	}
}

func TestLoopPosting(t *testing.T) {
	command := engine.NewPrintCommand("aloha")

	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	eventLoop.AwaitFinish()

	err := eventLoop.Post(command)
	assert.Error(t, err)
}