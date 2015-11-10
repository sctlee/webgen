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
	template_repo  string
	has_parser     bool
	FILES_TO_CHECK = []string{"info.yml", "papers.yml", "build",
		ASSETS_PATH, TEMPLATE_PATH}
	FILES_TO_CREATE = []string{"info.yml", "papers.yml", ".gitignore", "CNAME"}
)

func Init(t string, has_p bool) {
	// help init
	fmt.Println("start init")

	for _, filename := range FILES_TO_CHECK {
		if utils.Exists(filename) {
			fmt.Println("has already inited. Please run 'reset' first")
			os.Exit(-1)
		}
	}
	template_repo = t
	has_parser = has_p

	getTemplate()
	createGHPages()

	if has_parser {
		getParser()
	}
}

// get templete from DEFAULT_TMPL
func getTemplate() {
	var shell Shell

	if runtime.GOOS != "windows" {
		shell = &LinuxShell{sh.NewSession()}
	}
	shell.Gcl(template_repo, "master", TEMPLATE_PATH, 1)

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

// get parser to parse *.md to *.html
func getParser() {
	var shell Shell

	if runtime.GOOS != "windows" {
		shell = &LinuxShell{sh.NewSession()}
	}
	shell.Gcl(DEFAULT_PARSE, "master", PARSER_PATH, -1)
	gitignore, _ := os.OpenFile(".gitignore", os.O_APPEND, 0666)
	gitignore.WriteString(fmt.Sprintf("\n%s", PARSER_PATH))

	shell.Frm(PARSER_PATH, ".git")
}

func getOriginUrl() string {
	originUrl, err := gitconfig.OriginURL()
	if err == nil {
		return originUrl
	}
	return originUrl
}
