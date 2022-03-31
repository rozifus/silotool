package main

import (
	kcli "github.com/alecthomas/kong"
)


type CliContext struct { }

type CommandLine struct {
	Silo SiloCmd `cmd`
	Acr AcrCmd `cmd`
}

func (commandLine *CommandLine) Run() error {
	return nil
}

func main() {
	commandLine := &CommandLine{}

	ktx := kcli.Parse(commandLine)
	ktx.Run(&CliContext{})
}
