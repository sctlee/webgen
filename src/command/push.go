package command

import (
	"fmt"
	"runtime"

	"github.com/codeskyblue/go-sh"
)

func Push() {
	fmt.Println("start push")

	var shell Shell
	if runtime.GOOS != "windows" {
		shell = &LinuxShell{sh.NewSession()}
	}

	shell.Gmt("master", "new feature", true)
	shell.Dcd("./build")
	shell.Gmt("gh-pages", "new feature", true)
}
