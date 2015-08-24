package command

import (
	"fmt"
	"os"
	"runtime"
	"utils"

	"github.com/codeskyblue/go-sh"
	"github.com/tcnksm/go-gitconfig"
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
	var shell Shell

	if runtime.GOOS != "windows" {
		shell = &LinuxShell{sh.NewSession()}
	}
	shell.Gcl(DEFAULT_TMPL, TEMPLATE_PATH)
	shell.Fmv(TEMPLATE_PATH, ".", "channel.yml", "items.yml", ASSETS_PATH, ".gitignore", "CNAME")
	shell.Dmk(PSRC_PATH)
	shell.Frm(TEMPLATE_PATH, ".git")
	shell.Gmt("master", "git init", true)
	// rmFiles(session, ".git")
	// mvFiles(session, "channel.yml", "items.yml", ASSETS_PATH, ".gitignore", "CNAME")
	// gitCommit(session, "master", "git init", true)
}

func createGHPages() {
	originUrl := getOriginUrl()
	if runtime.GOOS != "windows" {
		ls := &LinuxShell{sh.NewSession()}
		ls.session.Command("git", "branch", "-D", GH_PAGES).Run()
		ls.session.Command("git", "checkout", "--orphan", GH_PAGES).Run()
		ls.session.Command("git", "rm", "-rf", ".").Run()
		ls.session.Command("touch", "index.html").Run()

		ls.Gmt("gh-pages", "web init", true)
		// gitCommit(session, "gh-pages", "web init", true)

		ls.session.Command("git", "checkout", "master").Run()

		ls.session.Command("git", "clone", "-b", GH_PAGES, originUrl, "build").Run()

		ls.Fcp(TEMPLATE_PATH, TARGET_PATH, "css", "fonts", "img", "js")
		// cpFiles(session, TEMPLATE_PATH, TARGET_PATH, "css", "font-awesome", "fonts", "img", "js")
	}
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
