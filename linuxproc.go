package linuxproc

//ok

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Memory struct{}

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

func memoryStatus() (status []byte, err error) {
	statuspath := fmt.Sprintf("/proc/meminfo")
	status, err = ioutil.ReadFile(statuspath)
	if err != nil {
		return
	}
	return
}

func sectionString(regex string, status []byte) (result string) {
	rx := regexp.MustCompile(regex)
	match := rx.FindAllSubmatch(status, -1)
	if len(match) != 0 {
		result = strings.TrimSpace(string(match[0][1]))
	}
	return
}

func sectionInt(regex string, status []byte) (result int) {
	rx := regexp.MustCompile(regex)
	match := rx.FindAllSubmatch(status, -1)
	if len(match) != 0 {
		result, _ = strconv.Atoi(strings.TrimSpace(string(match[0][1])))
	}
	return
}

func sectionSInt(regex string, status []byte) (result []int) {
	rx := regexp.MustCompile(regex)
	match := rx.FindAllSubmatch(status, -1)
	if len(match) != 0 {
		sresult := strings.Fields(strings.TrimSpace(string(match[0][1])))
		for _, r := range sresult {
			rx, _ := strconv.Atoi(r)
			result = append(result, rx)
		}
	}
	return
}

func (m Memory) MemTotal() (memTotal int, err error) {
	status, err := memoryStatus()
	if err != nil {
		return
	}
	smemTotal := strings.Fields(strings.TrimSpace(sectionString("MemTotal:(.*)", status)))
	memTotal, err = strconv.Atoi(smemTotal[0])
	return
}

func (m Memory) MemFree() (memFree int, err error) {
	status, err := memoryStatus()
	if err != nil {
		return
	}
	smemFree := strings.Fields(strings.TrimSpace(sectionString("MemFree:(.*)", status)))
	memFree, err = strconv.Atoi(smemFree[0])
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
	state = sectionString("State:(.*)", status)
	return
}

func (p Process) VmSize() (vmSize int, err error) {
	status, err := processStatus(p)
	if err != nil {
		return
	}
	svmSize := strings.Fields(strings.TrimSpace(sectionString("VmSize:(.*)", status)))
	vmSize, err = strconv.Atoi(svmSize[0])
	return
}

func (p Process) VmPeak() (vmPeak int, err error) {
	status, err := processStatus(p)
	if err != nil {
		return
	}
	svmPeak := strings.Fields(strings.TrimSpace(sectionString("VmPeak:(.*)", status)))
	vmPeak, err = strconv.Atoi(svmPeak[0])
	return
}

func (p Process) VmData() (vmData int, err error) {
	status, err := processStatus(p)
	if err != nil {
		return
	}
	svmData := strings.Fields(strings.TrimSpace(sectionString("VmData:(.*)", status)))
	vmData, err = strconv.Atoi(svmData[0])
	return
}

func (p Process) Uid() (uid []int, err error) {
	status, err := processStatus(p)
	if err != nil {
		return
	}
	uid = sectionSInt("Uid:(.*)", status)
	return
}

func (p Process) Gid() (uid []int, err error) {
	status, err := processStatus(p)
	if err != nil {
		return
	}
	uid = sectionSInt("Gid:(.*)", status)
	return
}

func (p Process) PPid() (ppid int, err error) {
	status, err := processStatus(p)
	if err != nil {
		return
	}
	ppid = sectionInt("PPid:(.*)", status)
	return
}
