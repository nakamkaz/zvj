package main

import (
	"bufio"
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
		log.Print("ZVJ-dlog: ", msg)
	}
}
func main() {
	dlog("Version: ", Version)
	quotregex := os.Args[1]
	dlog("regex: ", quotregex)
	rxop := regexp.MustCompile(quotregex)
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		line := stdin.Text()
		if rxop.MatchString(line) {
			fmt.Println(line)
		}
	}
}
