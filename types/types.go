package types

import "os"

type Service struct {
	Name    string   `yaml:"name"`
	Command string   `yaml:"command"`
	Path    string   `yaml:"path"`
	Env     []string `yaml:"env"`

	Process *os.Process
}

func (s *Service) GetEnv() (env string) {
	for i := range s.Env {
		env += "export " + s.Env[i] + " && "
	}
	return
}
