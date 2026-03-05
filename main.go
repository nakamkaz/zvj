package main

import (
	"bufio"
	"compress/bzip2"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

var (
	Version     string = ""
	VersionInfo bool   = false
	DebugFlag   bool   = false
	Bzip2On     bool   = false
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
	skipchar := 0
	flag.BoolVar(&invertmatch, "v", false, "-v(--invert-match) in AWK")
	flag.BoolVar(&Bzip2On, "bz", false, "decode bz2 stream")
	flag.BoolVar(&printmatchcount, "pc", false, "prints count to STDERR")
	flag.BoolVar(&DebugFlag, "debug", false, "default false; --debug when enable")
	flag.IntVar(&skipchar, "sk", 0, "Skip charactors from head")
	flag.BoolVar(&VersionInfo, "V", false, "show version")
	flag.Parse()
	if VersionInfo {
		fmt.Println("Version", Version)
		os.Exit(0)
	}
	if flag.NArg() != 0 {
		quotregex = flag.Args()[0]
		dlog("Args0 set regex:", quotregex)
	}
	dlog("Version: ", Version)
	dlog("-v flag is ", invertmatch)
	dlog("-bz flag is ", Bzip2On)
	dlog("process regex: ", quotregex)
	dlog("Skip num: ", skipchar)
	reader := bufio.NewReader(os.Stdin)
	dofilter(reader, quotregex, invertmatch, printmatchcount, skipchar)
}

func dofilter(reader io.Reader, quotregex string, invertmatch bool, printmatchcount bool, skip int) {
	var brrdr io.Reader
	if !Bzip2On {
		brrdr = reader
	} else {
		brrdr = bzip2.NewReader(reader)
	}
	scanner := bufio.NewScanner(brrdr)
	rxop := regexp.MustCompile(quotregex)
	matchcount := 0
	line := ""
	dlog("Start matching")
	for scanner.Scan() {
		line_all := scanner.Text()
		maxlen := len(line_all)
		if skip == 0 {
			line = line_all
		} else if (skip >= 1) && (skip < maxlen) {
			line = line_all[skip:]
		} else {
			line = line_all[maxlen:]
		}
		if invertmatch != rxop.MatchString(line) {
			if printmatchcount {
				matchcount++
			}
			fmt.Println(line)
		}
	}
	dlog("End matching")
	if printmatchcount {
		log.Println("Matched: ", !invertmatch,
			" for ", matchcount,
			" times on [", quotregex, "]")
	}
}
