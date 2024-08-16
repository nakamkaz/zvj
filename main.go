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
	Version     string = ""
	VersionInfo bool   = false
	DebugFlag   bool   = false
)

func dlog(msg ...interface{}) {
	if os.Getenv("ZVJ_DEBUG") == "1" || DebugFlag {
		log.Print("ZVJ_DEBUG:", msg)
	}
}
func main() {
	quotregex := ".*"
	invertmatch := false
	printmatchcount := false
	flag.BoolVar(&invertmatch, "v", false, "-v(--invert-match) in AWK")
	flag.BoolVar(&printmatchcount, "pc", false, "prints count to STDERR")
	flag.BoolVar(&DebugFlag, "debug", false, "default false; --debug when enable")
	flag.BoolVar(&VersionInfo, "V", false, "show version")
	flag.Parse()
	if VersionInfo {
		fmt.Println("Version", Version)
		os.Exit(0)
	}
	if flag.NArg() != 0 {
		quotregex = flag.Args()[0]
		dlog("set regex:", quotregex)
	}
	dlog("Version:", Version)
	dlog("-v flag is", invertmatch)
	dlog("process regex:", quotregex)
	rxop := regexp.MustCompile(quotregex)
	stdin := bufio.NewScanner(os.Stdin)
	matchcount := 0
	for stdin.Scan() {
		line := stdin.Text()
		if invertmatch != rxop.MatchString(line) {
			if printmatchcount {
				matchcount++
			}
			fmt.Println(line)
		}
	}
	if printmatchcount {
		log.Println("Matched: ", matchcount, " times")
	}
}
