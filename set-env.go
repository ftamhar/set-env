package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"syscall"
)

func main() {
	var rootPath string

	flag.StringVar(&rootPath, "f", ".env", "file env")
	flag.Parse()

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
	if os.Getenv(key) == "" {
		os.Setenv(key, val)
	}
}
