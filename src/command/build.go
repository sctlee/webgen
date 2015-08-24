package command

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/codeskyblue/go-sh"
)

type Info struct {
	Title       string
	Link        string
	Description string
	Image       string
	Copyright   string
	Language    string
	Author      string
	Categories  []string
	Page        int
	Twitter     string
	Github      string
	Linkedin    string
}

type Paper struct {
	Image       string
	Author      string
	Title       string
	Description string
	Link        string
	PubDate     string
	Tag         string
}

type WebTemplate struct {
	Info    Info
	Home    string
	Current Paper
	Papers  []Paper
}

func Build() {
	fmt.Println("start build")

	check := func(err error) {
		if err != nil {
			fmt.Println(err)
		}
	}

	info := getInfo("info.yml")
	items := getItems("papers.yml")

	content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", TEMPLATE_PATH, "index.tmpl"))
	check(err)

	f, err := os.Create(fmt.Sprintf("%s/%s", TARGET_PATH, "index.html"))
	check(err)

	funcs := template.FuncMap{"alt": alt, "trunc": truncate}
	t := template.Must(template.New("website").Funcs(funcs).Parse(string(content[:])))
	err = t.Execute(f, WebTemplate{
		Info:    info,
		Home:    "#current",
		Current: items[0],
		Papers:  items[1:],
	})
	check(err)

	if runtime.GOOS != "windows" {
		ls := &LinuxShell{sh.NewSession()}
		ls.Dmk(fmt.Sprintf("%s/%s", TARGET_PATH, "papers"))
		for i, item := range items {
			content_paper, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", TEMPLATE_PATH, "paper.tmpl"))
			check(err)

			f_paper, err := os.Create(fmt.Sprintf("%s/%s/%s", TARGET_PATH, "papers", fmt.Sprintf("%d.html", i+1)))
			check(err)

			t_paper := template.Must(template.New("paper").Parse(string(content_paper[:])))
			err = t_paper.Execute(f_paper, item)
		}
		ls.Fcp(".", TARGET_PATH, "assets")
		// ls.Fcp(".", TARGET_PATH, PSRC_PATH)
		ls.Fcp(TEMPLATE_PATH, TARGET_PATH, "css", "fonts", "img", "js")
		// cpFiles(session, ".", TARGET_PATH, "assets")
		// cpFiles(session, TEMPLATE_PATH, TARGET_PATH, "css", "font-awesome", "fonts", "img", "js")
	}
}

func getInfo(path string) (info Info) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(data, &info)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func getItems(path string) (items []Paper) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(data, &items)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func alt(x int) string {
	if x%2 == 0 {
		return "a"
	} else {
		return "b"
	}
}

func truncate(str string) string {
	data := []rune(str)
	if len(data) <= MAX_DESCRIPTION {
		return str
	} else {
		return string(data[:MAX_DESCRIPTION-1]) + "..."
	}
}
