package types

import (
	"fmt"
	"os"
	"os/exec"
)

type Service struct {
	Name    string   `yaml:"name"`
	Command string   `yaml:"command"`
	Path    string   `yaml:"path"`
	Env     []string `yaml:"env"`
	Ports   []int64  `yaml:"ports"`

	Process *os.Process
	Stopped bool
}

func (s *Service) Kill() {
	if len(s.Ports) == 0 && s.Process != nil {
		cmd := exec.Command("kill", "-9", fmt.Sprintf("%d", s.Process.Pid))
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		_ = cmd.Run()
	}

	for _, port := range s.Ports {
		cmd := exec.Command("fuser", "-k", fmt.Sprintf("%d/tcp", port))
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		_ = cmd.Run()
	}
}

func (s *Service) GetEnv() (env string) {
	for i := range s.Env {
		env += "export " + s.Env[i] + " && "
	}
	return
}
