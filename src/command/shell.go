package command

import (
	"fmt"

	"github.com/codeskyblue/go-sh"
)

type Shell interface {
	Frm(files ...string)
	Fmv(src string, dest string, files ...string)
	Fcp(src string, dest string, files ...string)
	Gmt(branch string, message string, setUpstream bool)
}

type LinuxShell struct {
	session *sh.Session
}

type WindowsShell struct{}

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

func (ls *LinuxShell) Gmt(branch string, message string, setUpstream bool) {
	ls.session.Command("git", "add", ".").Run()
	ls.session.Command("git", "commit", "-a", "-m", message).Run()
	if setUpstream {
		ls.session.Command("git", "push", "-u", "origin", branch).Run()
	} else {
		ls.session.Command("git", "push", "origin", branch).Run()
	}
}
