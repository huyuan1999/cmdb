package Linux

import (
	"cmdb/st"
	"cmdb/utils"
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
)

type System struct {
	st.System
	info *host.InfoStat
}

func NewSystem() (*System, error) {
	system := &System{}
	info, err := host.Info()
	if err != nil {
		return nil, err
	}
	system.info = info
	utils.Call(system)
	return system, nil
}

func (s *System) GetHostName() {
	s.HostName = s.info.Hostname
}

func (s *System) GetKernel() {
	s.Kernel = s.info.KernelVersion
}

func (s *System) GetRelease() {
	s.Release = fmt.Sprintf("%s %s", s.info.Platform, s.info.PlatformVersion)
}
