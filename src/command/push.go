package command

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
)

func Push() {
	fmt.Println("start push")
	session := sh.NewSession()
	gitCommit(session, "master", "new feature", true)
	session.SetDir("./build")
	gitCommit(session, "gh-pages", "new feature", true)
}
