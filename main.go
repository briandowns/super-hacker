package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/briandowns/super-hacker/templates"
	term "github.com/nsf/termbox-go"
)

var (
	name    string
	version string
	gitSHA  string
)

// reset resets the terminal back to defaults.
func reset() {
	term.Sync()
}

const usage = `version: %s - git: %s
Usage: %s [-bvh]
Options:
	-h            help menu
	-v            show version
	-b            output buffer size
	-l            language
Examples:
	%[3]s -b 24
`

func main() {
	flag.Usage = func() {
		w := os.Stderr
		for _, arg := range os.Args {
			if arg == "-h" {
				w = os.Stdout
				break
			}
		}
		fmt.Fprintf(w, usage, version, gitSHA, name)

	}
	var vers bool
	var bufSize int
	var lang string
	flag.BoolVar(&vers, "v", false, "")
	flag.IntVar(&bufSize, "b", 3, "")
	flag.StringVar(&lang, "l", "", "")
	flag.Parse()

	if vers {
		fmt.Fprintf(os.Stdout, "version: %s - %s\n", version, gitSHA)
		return
	}

	if lang == "" {
		fmt.Println("error: -l flag required")
		return
	}

	rand.Seed(time.Now().Unix())

	if err := term.Init(); err != nil {
		fmt.Println(err)
		return
	}
	defer term.Close()

	code, err := templates.Random(lang)
	if err != nil {
		fmt.Println(err)
		return
	}
	codeBuf := strings.NewReader(code)
	buffer := make([]byte, bufSize)

keyPressListenerLoop:
	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc, term.KeyCtrlC:
				reset()
				break keyPressListenerLoop
			default:
				read, err := codeBuf.Read(buffer)
				if err != nil {
					if err != io.EOF {
						fmt.Println(err)
					}
					break
				}
				fmt.Print(string(buffer[:read]))
			}
		case term.EventError:
			fmt.Println(ev.Err)
			return
		}
	}
}
