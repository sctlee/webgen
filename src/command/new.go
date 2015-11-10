package command

import (
	"log"
	"net/http"
	"time"
	// "runtime"

	// "github.com/codeskyblue/go-sh"
)

func New() {
	// var shell Shell
	//
	// if runtime.GOOS != "windows" {
	// 	shell = &LinuxShell{sh.NewSession()}
	// }
	url := "http://127.0.0.1:8080/index.html"

	go func() {
		time.Sleep(1)
		log.Printf("Starting your default browser with %s\n", url)
	}()

	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("./parser/jison"))))

}
