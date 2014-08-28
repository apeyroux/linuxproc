package linuxproc

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type Process struct {
	Name string
	Pid  int
}

func processStatus(p Process) (status []byte, err error) {
	statuspath := fmt.Sprintf("/proc/%d/status", p.Pid)
	status, err = ioutil.ReadFile(statuspath)
	if err != nil {
		return
	}
	return
}

func FindProcess(pid int) (p *Process, err error) {
	p = new(Process)
	p.Pid = pid
	status, err := processStatus(*p)
	if err != nil {
		return
	}
	rxName := regexp.MustCompile("Name:(.*)")
	matchName := rxName.FindAllSubmatch(status, -1)
	if len(matchName) != 0 {
		p.Name = strings.TrimSpace(string(matchName[0][1]))
	}
	return
}

func (p Process) State() (state string, err error) {
	status, err := processStatus(p)
	if err != nil {
		return
	}
	rxState := regexp.MustCompile("State:(.*)")
	matchState := rxState.FindAllSubmatch(status, -1)
	if len(matchState) != 0 {
		state = strings.TrimSpace(string(matchState[0][1]))
	}
	return
}

func (p Process) VmSize() (vmSize string, err error) {
	status, err := processStatus(p)
	if err != nil {
		return
	}
	rxVmSize := regexp.MustCompile("VmSize:(.*)")
	matchVmSize := rxVmSize.FindAllSubmatch(status, -1)
	if len(matchVmSize) != 0 {
		vmSize = strings.TrimSpace(string(matchVmSize[0][1]))
	}
	return
}
