package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	spin "github.com/tj/go-spin"
)

type Process struct {
	message string
	timer   *time.Ticker
	out     io.Writer
	spin    *spin.Spinner
}

func (self *Process) Start() {
	if self.timer != nil {
		return
	}

	self.timer = time.NewTicker(100 * time.Millisecond)
	spin := spin.New()

	go func() {
		for range self.timer.C {
			self.print(fmt.Sprintf("\r  \033[36m%s\033[m %s ", self.message, spin.Next()))
		}
	}()

}

func (self *Process) print(msg string) {
	self.out.Write([]byte(msg))
}

func (self *Process) stop(msg string) {
	if self.timer == nil {
		return
	}
	self.timer.Stop()
	self.print(fmt.Sprintf("\r  \033[36m%s\033[m %s\n", self.message, msg))
}

func (self *Process) Fail(msg string) {
	self.stop(msg)
}

func (self *Process) Sucess(msg string) {
	self.stop(msg)
}

func NewProcess(msg string, fn func() error) error {

	p := Process{
		message: msg,
		out:     os.Stdout,
	}

	p.Start()
	if err := fn(); err != nil {
		p.Fail("fail")
		return err
	}

	p.Sucess("ok")

	return nil
}
