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

type Channel struct {
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
	Title       string
	Description string
	Link        string
}

type WebTemplate struct {
	Info    Channel
	Home    string
	Current Paper
}

func Build() {
	fmt.Println("start build")

	check := func(err error) {
		if err != nil {
			fmt.Println(err)
		}
	}

	channel := getChannel("channel.yml")
	items := getItems("items.yml")

	content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", TEMPLATE_PATH, "myindex.tmpl"))
	check(err)

	f, err := os.Create(fmt.Sprintf("%s/%s", TARGET_PATH, "index.html"))
	check(err)

	funcs := template.FuncMap{"alt": alt, "trunc": truncate}
	t := template.Must(template.New("website").Funcs(funcs).Parse(string(content[:])))
	err = t.Execute(f, WebTemplate{
		Info:    channel,
		Home:    "#current",
		Current: items[0],
	})
	check(err)

	if runtime.GOOS != "windows" {
		ls := &LinuxShell{sh.NewSession()}
		ls.Fcp(".", TARGET_PATH, "assets")
		ls.Fcp(TEMPLATE_PATH, TARGET_PATH, "css", "font-awesome", "fonts", "img", "js")
		// cpFiles(session, ".", TARGET_PATH, "assets")
		// cpFiles(session, TEMPLATE_PATH, TARGET_PATH, "css", "font-awesome", "fonts", "img", "js")
	}
}

func getChannel(path string) (channel Channel) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(data, &channel)
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
