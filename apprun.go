package apprun

import (
	"errors"
	"os"

	"github.com/Masterminds/vcs"
)

type Apprun struct {
	config Config
	runner Runner
	repo   vcs.Repo
}

func (self *Apprun) Init() error {
	if self.repo == nil {
		return nil
	}

	if !DirExists(self.config.Workspace) {
		err := self.repo.Get()
		if err != nil {
			return err
		}

		if self.config.Branch != "" {
			self.repo.UpdateVersion(self.config.Branch)
		}
	} else if empty, _ := IsEmpty(self.config.Workspace); empty {

	}

	return nil
}

func (self *Apprun) Update() error {

	if self.repo != nil {

		v, _ := self.repo.Version()

		if self.config.Branch != "" && v != self.config.Branch {
			self.repo.UpdateVersion(self.config.Branch)
		}

		if err := self.repo.Update(); err != nil {
			return err
		}
	}

	return nil
}

func (self *Apprun) Run() error {

	ctx := RunContext{
		Stderr:  os.Stderr,
		Stdout:  os.Stdout,
		Config:  self.config,
		Environ: self.config.Environ,
	}

	return self.runner.Run(ctx)
}

func (self *Apprun) Build() error {
	ctx := RunContext{
		Stderr:  os.Stderr,
		Stdout:  os.Stdout,
		Config:  self.config,
		Environ: self.config.Environ,
	}

	return self.runner.Build(ctx)

}

func New(config Config) (*Apprun, error) {
	var (
		runner Runner
		ok     bool
		repo   vcs.Repo
		err    error
	)

	if runner, ok = runners[config.Kind]; !ok {
		return nil, errors.New("unknown kind " + config.Kind)
	}

	if config.Remote != "" {
		if repo, err = vcs.NewRepo(config.Remote, config.Workspace); err != nil {
			return nil, err
		}
	}

	return &Apprun{config, runner, repo}, nil
}
