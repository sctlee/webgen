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
	Gmt(branch string, message string, setUpstream bool)
	Gcl(url string, path string)
}

type LinuxShell struct {
	session *sh.Session
}

type WindowsShell struct{}

func (ls *LinuxShell) Fmk(files ...string) {
	for _, file := range files {
		ls.session.Command("touch", file)
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

func (ls *LinuxShell) Gmt(branch string, message string, setUpstream bool) {
	ls.session.Command("git", "add", ".").Run()
	ls.session.Command("git", "commit", "-a", "-m", message).Run()
	if setUpstream {
		ls.session.Command("git", "push", "-u", "origin", branch).Run()
	} else {
		ls.session.Command("git", "push", "origin", branch).Run()
	}
}

func (ls *LinuxShell) Gcl(url string, path string) {
	ls.session.Command("git", "clone", "--depth=1", url, path).Run()
}
