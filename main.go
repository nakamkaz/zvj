package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

var (
	Version string = ""
)

func dlog(msg ...interface{}) {
	if os.Getenv("ZVJ_DEBUG") == "1" {
		log.Print("ZVJ-debug:", msg)
	}
}
func main() {
	dlog("Version: ", Version)
	quotregex := ".*"
	invertmatch := false
	flag.BoolVar(&invertmatch, "v", false, "-v(--invert-match) in AWK")
	flag.Parse()
	if flag.NArg() != 0 {
		quotregex = flag.Args()[0]
		dlog("set regex:", quotregex)
	}
	dlog("-v flag is", invertmatch)
	dlog("process regex:", quotregex)
	rxop := regexp.MustCompile(quotregex)
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		line := stdin.Text()
		if invertmatch != rxop.MatchString(line) {
			fmt.Println(line)
		}
	}
}
