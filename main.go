package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync"

	"github.com/zserge/lorca"
)

var (
	SOURCE_DIR string
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

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(http.Dir(getHtmlDir())))
	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

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

func getHtmlDir() string {
	if SOURCE_DIR == "" {
		return "./dist"
	}
	return SOURCE_DIR
}
