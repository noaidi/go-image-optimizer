package main

import (
	"fmt"
	"os"

	"github.com/noaidi/imgopt"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go <from> <to>")
		os.Exit(1)
	}

	if err := optimize(os.Args[1], os.Args[2]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func optimize(from string, to string) error {
	src, err := os.Open(from)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.OpenFile(to, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer dst.Close()

	return imgopt.Optimize(src, dst, 3600, 60)
}
