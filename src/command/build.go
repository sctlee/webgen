package command

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"utils"

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
}

type WebTemplate struct {
	Info    Info
	Home    string
	Current Paper
	Content template.HTML
	Prev    string
	Next    string
	Papers  []Paper
}

func Build() {
	fmt.Println("start build")

	// Init environment shell
	var shell Shell
	if runtime.GOOS != "windows" {
		shell = &LinuxShell{sh.NewSession()}
	}

	// update resource files
	shell.Fcp(".", TARGET_PATH, ASSETS_PATH)
	shell.Fcp(TEMPLATE_PATH, TARGET_PATH, "css", "fonts", "img", "js")

	// generate real content
	// get user's data
	info := getInfo("info.yml")
	items := getItems("papers.yml")

	// get index template
	content_index, err := ioutil.ReadFile(TEMPLATE_INDEX_FILE)
	utils.Check(err)
	funcs := template.FuncMap{"alt": alt, "trunc": truncate, "judge": judge, "draw": draw}
	t_index := template.Must(template.New("website").Funcs(funcs).Parse(string(content_index[:])))

	// get single papar template
	content_paper, err := ioutil.ReadFile(TEMPLATE_PAPER_FILE)
	utils.Check(err)
	t_paper := template.Must(template.New("paper").Funcs(funcs).Parse(string(content_paper[:])))

	// generate paper single html
	shell.Dmk(PAPER_TARGET_PATH)
	l := len(items)
	for i, item := range items {
		f_paper, err := os.Create(fmt.Sprintf("%s/%s", PAPER_TARGET_PATH, fmt.Sprintf("%d.html", i+1)))
		utils.Check(err)

		content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", PAPER_SRC_PATH, item.Link))
		utils.Check(err)
		fn, _ := f_paper.Stat()

		// build paper link current, prev, next
		items[i].Link = fn.Name()
		var prev, next string
		if i-1 >= 0 {
			prev = items[i-1].Title
			prev += fmt.Sprintf("|%d.html", i)
		}
		if i+1 < l {
			next = items[i+1].Title
			next += fmt.Sprintf("|%d.html", i+2)
		}
		err = t_paper.Execute(f_paper, WebTemplate{
			Info:    info,
			Current: items[0],
			Prev:    prev,
			Next:    next,
			Content: template.HTML(content[:]),
		})
		utils.Check(err)
	}

	// generate index html
	f_index, err := os.Create(fmt.Sprintf("%s/%s", TARGET_PATH, "index.html"))
	utils.Check(err)
	err = t_index.Execute(f_index, WebTemplate{
		Info:    info,
		Current: items[0],
		Papers:  items[1:],
	})
	utils.Check(err)

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

func judge(str string) bool {
	if str == "" {
		return false
	} else {
		return true
	}
}

func draw(str string, column int) string {
	return strings.Split(str, "|")[column]
}
