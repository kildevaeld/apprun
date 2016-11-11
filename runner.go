package apprun

import "io"

type RunContext struct {
	Config
	Stdout  io.Writer
	Stderr  io.Writer
	Environ []string
}

type Runner interface {
	Build(RunContext) error
	Run(RunContext) error
}

var runners map[string]Runner

func Register(name string, runner Runner) {
	runners[name] = runner
}

func init() {
	runners = make(map[string]Runner)
}
