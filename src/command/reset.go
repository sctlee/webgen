package command

import (
	"fmt"

	"github.com/codeskyblue/go-sh"
)

func Reset() {
	fmt.Println("start reset")

	shell := &LinuxShell{sh.NewSession()}

	shell.session.Command("git", "branch", "-D", GH_PAGES).Run()
	shell.session.Command("git", "push", "origin", "--delete", GH_PAGES).Run()
	shell.Frm(".", TARGET_PATH, TEMPLATE_PATH, ASSETS_PATH, PSRC_PATH, "info.yml", "papers.yml")
	shell.Gmt("master", "reset", true)

}
