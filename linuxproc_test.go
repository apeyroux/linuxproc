package linuxproc

import (
	"testing"
)

func TestFindProcess(t *testing.T) {
	if p, _ := FindProcess(1); p.Name != "systemd" {
		t.Error("p.Name != de systemd en pid 1 ?")
	}
}

func TestUid(t *testing.T) {
	p, _ := FindProcess(1)
	if uid, _ := p.Uid(); len(uid) == 0 {
		t.Error("Oops len uid == 0 ?")
	}
}

func TestGid(t *testing.T) {
	p, _ := FindProcess(1)
	if gid, _ := p.Gid(); len(gid) == 0 {
		t.Error("Oops len gid == 0 ?")
	}
}

func TestPPid(t *testing.T) {
	p, _ := FindProcess(1)
	if ppid, _ := p.PPid(); ppid != 0 {
		t.Error("Oops PPID de systemd != 0 ?")
	}
}

func TestState(t *testing.T) {
	p, _ := FindProcess(1)
	if state, _ := p.State(); state != "S (sleeping)" {
		t.Error("Oops state systemd != S (sleeping) ?")
	}
}

func TestVmData(t *testing.T) {
	p, _ := FindProcess(1)
	if vmdata, _ := p.VmData(); vmdata == 0 {
		t.Error("Oops vmdata == 0 ?")
	}
}

func TestVmPeak(t *testing.T) {
	p, _ := FindProcess(1)
	if vmpeak, _ := p.VmPeak(); vmpeak == 0 {
		t.Error("Oops vmpeak == 0 ?")
	}
}

func TestVmSize(t *testing.T) {
	p, _ := FindProcess(1)
	if vmsize, _ := p.VmSize(); vmsize == 0 {
		t.Error("Oops vmsize == 0 ?")
	}
}

// Test memory

func TestMemTotal(t *testing.T) {
	var m = Memory{}
	if mt, _ := m.MemTotal(); mt == 0 {
		t.Error("Oops mémoire dispo. 0 ?")
	}
}

func TestMemFree(t *testing.T) {
	var m = Memory{}
	if mf, _ := m.MemFree(); mf == 0 {
		t.Error("Oops mémoire free 0 ?")
	}
}
