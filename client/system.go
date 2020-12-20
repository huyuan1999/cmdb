package client

import (
	"cmdb/models"
	"cmdb/utils"
	"io/ioutil"
	"path"
	"strings"
)

type System struct {
	models.System
	readDir string
}

func NewSystem(readDir string) *System {
	system := &System{}
	system.readDir = readDir
	return system
}

func (s *System) GetHostName() {
	s.HostName = utils.Shell("hostname").Stdout
}

func (s *System) GetRelease() {
	s.Kernel = utils.Shell("uname", "-r").Stdout
}

func (s *System) GetKernel() {
	dir, err := ioutil.ReadDir(s.readDir)
	if err != nil {
		return
	}
	for _, item := range dir {
		releasePath := path.Join(s.readDir, item.Name())
		if strings.HasSuffix(item.Name(), "release") && utils.IsFile(releasePath) {
			text, err := ioutil.ReadFile(releasePath)
			if err != nil {
				continue
			}
			s.Release = strings.Trim(string(text), "\n")
			return
		}
	}
}
