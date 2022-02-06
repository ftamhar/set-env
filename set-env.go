package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
)

var boolReplace bool

func main() {
	boolReplace = true
	var rootPath string

	var replace string
	flag.StringVar(&rootPath, "f", ".env", "file env")
	flag.StringVar(&replace, "r", "true", "replace environment variables")
	flag.Parse()

	if replace != "true" && replace != "false" {
		log.Fatalf("option -r must be %q or %q", "true", "false")
	}

	if replace == "false" {
		boolReplace = false
	}

	file, err := os.ReadFile(rootPath)
	if err != nil {
		log.Fatalf("\nfile %s not found\nplease set file with -f\nfor help: set-env -h", rootPath)
	}

	arrEnv := strings.Split(string(file), "\n")

	for _, v := range arrEnv {
		idx := strings.Index(v, "=")

		if idx < 1 {
			continue
		}

		idxComment := strings.Index(v, "#")

		// no comment on end env value
		if idxComment == -1 {
			setEnv(v[:idx], strings.Trim(v[idx+1:], " "))
		} else if idxComment > idx { // have comment on end env value
			setEnv(v[:idx], strings.Trim(v[idx+1:idxComment], " "))
		}
	}
	syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())
}

func setEnv(key, val string) {
	before := os.Getenv(key)
	if before == "" || boolReplace {
		if before != val {
			fmt.Printf("%s : Before = %q, After = %q\n", key, before, val)
			os.Setenv(key, val)
		}
	}
}
