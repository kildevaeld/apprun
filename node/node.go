package node

import (
	"os/exec"

	"github.com/kildevaeld/apprun"
)

type Node struct {
}

func (self *Node) Run(ctx apprun.RunContext) error {

	bin := ctx.Command
	args := ctx.Args
	if bin == "" {
		bin = "npm"
		args = append([]string{"start"}, args...)
	}

	cmd := exec.Command(bin, args...)

	cmd.Stdout = ctx.Stdout
	cmd.Stderr = ctx.Stderr
	cmd.Dir = ctx.Workspace
	cmd.Env = ctx.Environ

	return cmd.Run()
}

func (self *Node) Build(ctx apprun.RunContext) error {

	cmd := exec.Command("npm", "i", "--production")
	cmd.Stdout = ctx.Stdout
	cmd.Stderr = ctx.Stderr
	cmd.Dir = ctx.Workspace
	cmd.Env = ctx.Environ

	return cmd.Run()
}

func init() {
	apprun.Register("node", &Node{})
}
