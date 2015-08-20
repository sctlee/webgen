package command

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
)

const (
	DEFAULT_TMPL    = "https://github.com/tyrchen/podgen-basic" // SDF
	TEMPLATE_PATH   = "template"
	GH_PAGES        = "gh-pages"
	TARGET_PATH     = "build"
	ASSETS_PATH     = "assets"
	MAX_DESCRIPTION = 96
	DEFAULT_PORT    = 6060
)

func Init() {
	// help init
	fmt.Println("start init")
	session := sh.NewSession()
	session.Command("git", "clone", "--depth=1", DEFAULT_TMPL, TEMPLATE_PATH).Run()
}

func getTemplate() {
	session := sh.NewSession()
	session.Command("git", "clone", "--depth=1", DEFAULT_TMPL, TEMPLATE_PATH).Run()
	rmFiles(session, ".git")
	mvFiles(session, "channel.yml", "items.yml", ASSETS_PATH, ".gitignore", "CNAME")
	gitCommit(session, "master", "git init", true)
}

func rmFiles(session *sh.Session, files ...string) {
	for _, filename := range files {
		session.Command("rm", "-rf", fmt.Sprintf("%s/%s", TEMPLATE_PATH, filename))
	}
}

func mvFiles(session *sh.Session, files ...string) {
	for _, filename := range files {
		session.Command("mv", fmt.Sprintf("%s/%s", TEMPLATE_PATH, filename, "."))
	}
}

func gitCommit(session *sh.Session, branch string, message string, setUpstream bool) {
	session.Command("git", "add", ".").Run()
	session.Command("git", "commit", "-a", "-m", message).Run()
	if setUpstream {
		session.Command("git", "push", "-u", "origin", branch).Run()
	} else {
		session.Command("git", "push", "origin", branch).Run()
	}
}
