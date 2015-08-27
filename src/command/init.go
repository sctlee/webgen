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
	FILES_TO_CHECK = []string{"info.yml", "papers.yml", "build",
		ASSETS_PATH, TEMPLATE_PATH}
	FILES_TO_CREATE = []string{"info.yml", "papers.yml", ".gitignore", "CNAME"}
)

func Init() {
	// help init
	fmt.Println("start init")

	for _, filename := range FILES_TO_CHECK {
		if utils.Exists(filename) {
			fmt.Println("has already inited. Please run 'reset' first")
			os.Exit(-1)
		}
	}

	getTemplate()
	createGHPages()
}

// get templete from DEFAULT_TMPL
func getTemplate() {
	var shell Shell

	if runtime.GOOS != "windows" {
		shell = &LinuxShell{sh.NewSession()}
	}
	shell.Gcl(DEFAULT_TMPL, "master", TEMPLATE_PATH, 1)

	// create resource files and .gitignore CNAME
	shell.Fmv(TEMPLATE_PATH, ".", "info.yml", "papers.yml", ASSETS_PATH, ".gitignore", "CNAME")

	// init gitignore file
	gitignore, _ := os.Create(".gitignore")
	gitignore.WriteString(fmt.Sprintf("%s\n%s", TARGET_PATH, PAPER_SRC_PATH))

	for _, filename := range FILES_TO_CREATE {
		if !utils.Exists(filename) {
			os.Create(filename)
		}
	}

	shell.Dmk(PAPER_SRC_PATH)
	shell.Frm(TEMPLATE_PATH, ".git")
	shell.Gmt("master", "git init", true)
}

// create initial gh-pages which contains only one index.html
func createGHPages() {
	var shell Shell
	originUrl := getOriginUrl()
	if runtime.GOOS != "windows" {
		shell = &LinuxShell{sh.NewSession()}
	}

	shell.Gck(GH_PAGES, true)
	shell.Gclear()
	shell.Fmk("index.html")
	shell.Gmt("gh-pages", "web init", true)

	shell.Gck("master", false)
	shell.Gcl(originUrl, GH_PAGES, "build", -1)
	shell.Fcp(TEMPLATE_PATH, TARGET_PATH, "css", "fonts", "img", "js")
}

func getOriginUrl() string {
	originUrl, err := gitconfig.OriginURL()
	if err == nil {
		return originUrl
	}
	return originUrl
}
