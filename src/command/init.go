package command

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
	"github.com/tcnksm/go-gitconfig"
	"os"
	"utils"
)

var (
	FILES_TO_CHECK = []string{"channel.yml", "item.yml", "build",
		ASSETS_PATH, TEMPLATE_PATH}
)

func Init() {
	// help init
	fmt.Println("start init")

	for _, filename := range FILES_TO_CHECK {
		if utils.Exists(filename) {
			fmt.Println("has already inited")
			os.Exit(-1)
		}
	}
	getTemplate()
	createGHPages()
}

func getTemplate() {
	session := sh.NewSession()
	session.Command("git", "clone", "--depth=1", DEFAULT_TMPL, TEMPLATE_PATH).Run()
	rmFiles(session, ".git")
	mvFiles(session, "channel.yml", "items.yml", ASSETS_PATH, ".gitignore", "CNAME")
	gitCommit(session, "master", "git init", true)
}

func createGHPages() {
	originUrl := getOriginUrl()
	session := sh.NewSession()
	session.Command("git", "branch", "-D", GH_PAGES).Run()
	session.Command("git", "checkout", "--orphan", GH_PAGES).Run()
	session.Command("git", "rm", "-rf", ".").Run()
	session.Command("touch", "index.html").Run()

	gitCommit(session, "gh-pages", "web init", true)

	session.Command("git", "checkout", "master").Run()

	session.Command("git", "clone", "-b", GH_PAGES, originUrl, "build").Run()

	cpFiles(session, TEMPLATE_PATH, TARGET_PATH, "css", "font-awesome", "fonts", "img", "js")

}

func rmFiles(session *sh.Session, files ...string) {
	for _, filename := range files {
		session.Command("rm", "-rf", fmt.Sprintf("%s/%s", TEMPLATE_PATH, filename)).Run()
	}
}

func mvFiles(session *sh.Session, files ...string) {
	for _, filename := range files {
		session.Command("mv", fmt.Sprintf("%s/%s", TEMPLATE_PATH, filename), ".").Run()
	}
}

func cpFiles(session *sh.Session, src string, dest string, files ...string) {
	for _, filename := range files {
		session.Command("cp", "-r", fmt.Sprintf("%s/%s", src, filename), dest).Run()
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

func getOriginUrl() string {
	originUrl, err := gitconfig.OriginURL()
	if err == nil {
		return originUrl
	}
	return originUrl
}
