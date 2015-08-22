package command

import (
	"fmt"
	"runtime"

	"github.com/codeskyblue/go-sh"
)

func Push() {
	fmt.Println("start push")
	if runtime.GOOS != "windows" {
		ls := &LinuxShell{sh.NewSession()}
		ls.Gmt("master", "new feature", true)
		ls.session.SetDir("./build")
		ls.Gmt("gh-pages", "new feature", true)
	}
}
