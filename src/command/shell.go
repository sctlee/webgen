package command

import (
	"fmt"

	"github.com/codeskyblue/go-sh"
)

type Shell interface {
	Fmk(files ...string)
	Frm(src string, files ...string)
	Fmv(src string, dest string, files ...string)
	Fcp(src string, dest string, files ...string)
	Dmk(path string)
	Dcd(path string)
	Gmt(branch string, message string, setUpstream bool)
	Gcl(url string, branch string, path string, depth int)
	Gck(branch string, isOrphan bool)
	Gclear()
}

type LinuxShell struct {
	session *sh.Session
}

type WindowsShell struct{}

func (ls *LinuxShell) Fmk(files ...string) {
	for _, file := range files {
		ls.session.Command("touch", file).Run()
	}
}

func (ls *LinuxShell) Frm(src string, files ...string) {
	for _, filename := range files {
		ls.session.Command("rm", "-rf", fmt.Sprintf("%s/%s", src, filename)).Run()
	}
}

func (ls *LinuxShell) Fmv(src string, dest string, files ...string) {
	for _, filename := range files {
		ls.session.Command("mv", fmt.Sprintf("%s/%s", src, filename), dest).Run()
	}
}

func (ls *LinuxShell) Fcp(src string, dest string, files ...string) {
	for _, filename := range files {
		ls.session.Command("cp", "-r", fmt.Sprintf("%s/%s", src, filename), dest).Run()
	}
}

func (ls *LinuxShell) Dmk(path string) {
	ls.session.Command("mkdir", "-p", path).Run()
}

func (ls *LinuxShell) Dcd(path string) {
	ls.session.SetDir(path)
}

func (ls *LinuxShell) Gmt(branch string, message string, setUpstream bool) {
	ls.session.Command("git", "add", ".").Run()
	ls.session.Command("git", "commit", "-a", "-m", message).Run()
	if setUpstream {
		ls.session.Command("git", "push", "-u", "origin", branch).Run()
	} else {
		ls.session.Command("git", "push", "origin", branch).Run()
	}
}

func (ls *LinuxShell) Gcl(url string, branch string, path string, depth int) {
	if depth > 0 {
		ls.session.Command("git", "clone", fmt.Sprintf("--depth=%d", depth), "-b", branch, url, path).Run()
	} else {
		ls.session.Command("git", "clone", "-b", branch, url, path).Run()
	}
}

func (ls *LinuxShell) Gck(branch string, isOrphan bool) {
	if isOrphan {
		ls.session.Command("git", "checkout", "--orphan", branch).Run()
	} else {
		ls.session.Command("git", "checkout", branch).Run()
	}
}

func (ls *LinuxShell) Gclear() {
	ls.session.Command("git", "rm", "-rf", ".").Run()
}
