package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
)

var (
	boolReplace bool = true
	verbose     bool = false
)

func main() {
	var rootPath string
	var replace string

	flag.StringVar(&rootPath, "f", ".env", "file env")
	flag.StringVar(&replace, "r", "true", "replace environment variables")
	flag.BoolVar(&verbose, "v", false, "verbose")
	flag.Parse()

	if replace != "true" && replace != "false" {
		log.Fatalf("option -r must be %q or %q", "true", "false")
	}

	boolReplace = replace == "true"

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

		setEnv(v[:idx], strings.Trim(v[idx+1:], " "))
	}
	syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())
}

func setEnv(key, val string) {
	before := os.Getenv(key)
	if val[len(val)-1] == '"' && val[0] == '"' {
		val = strings.Trim(val, `"`)
	}
	if before == "" || boolReplace {
		if before != val {
			if verbose {
				fmt.Printf("%s : Before = %q, After = %q\n", key, before, val)
			}
			os.Setenv(key, val)
		}
	}
}
