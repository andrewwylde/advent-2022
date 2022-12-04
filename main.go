package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func main() {
	args := os.Args[1:]
	dir, _ := os.Getwd()
	if len(args) == 0 {
		fl, _ := os.ReadDir(dir)
		for _, de := range fl {
			if de.IsDir() && len(de.Name()) == 1 {
				runDay(path.Join(dir, de.Name()))
			}
		}
	}
	if len(args) == 1 {
		runDay(path.Join(dir, args[0]))
	}
}

func runDay(day string) (err error) {
	cmd := exec.Command("go", "run", day)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return
}
