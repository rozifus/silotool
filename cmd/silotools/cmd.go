package main

import (
	kcli "github.com/alecthomas/kong"
)


type CliContext struct { }

type CommandLine struct {
	Acr AcrCmd `cmd`
	Chip ChipCmd `cmd`
	Halo HaloCmd `cmd`
	Silo SiloCmd `cmd`
}

func (commandLine *CommandLine) Run() error {
	return nil
}

func main() {
	commandLine := &CommandLine{}

	ktx := kcli.Parse(commandLine)
	ktx.Run(&CliContext{})
}
