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
	runDay(path.Join(dir, args[0]))

}

func runDay(day string) (err error) {
	args := []string{"run", day}
	args = append(args, os.Args...)
	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return
}
