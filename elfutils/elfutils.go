package elfutils

import (
	"os"
	"path"
	"strings"
)

func SplitByLine(t string) []string {
	tmp := []string{}
	for _, v := range strings.Split(t, "\n") {
		if len(v) > 1 {
			tmp = append(tmp, v)
		}
	}
	return tmp
}

func GetInputByDay(day string) (f []byte) {
	d, _ := os.Getwd()
	p := path.Join(d, day, "input.txt")
	f, _ = os.ReadFile(p)
	return
}
