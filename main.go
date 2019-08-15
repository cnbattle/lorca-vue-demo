package main

import (
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/signal"
	"runtime"
	"sync"

	"github.com/zserge/lorca"
)

type counter struct {
	sync.Mutex
	count int
}

func (c *counter) Add(n int) {
	c.Lock()
	defer c.Unlock()
	c.count = c.count + n
}

func (c *counter) Value() int {
	c.Lock()
	defer c.Unlock()
	return c.count
}

func main() {
	args := getArgs()
	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("", "", 480, 320, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	c := &counter{}
	ui.Bind("counterAdd", c.Add)
	ui.Bind("counterValue", c.Value)
	html, err := getHtmlString("./index.html")
	if err != nil {
		log.Fatal(err)
	}
	ui.Load("data:text/html," + url.PathEscape(html))

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}
	log.Println("exiting...")
}

// getArgs
func getArgs() (args []string) {
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	return args
}

// getHtml
func getHtmlString(path string) (html string, err error) {
	bt, err := ioutil.ReadFile(path)
	return string(bt), err
}
