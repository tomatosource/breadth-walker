package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var ignores = []string{
	"node_modules",
	".git",
	"vendor",
}

func main() {
	args := os.Args[1:]
	limit := 50
	if len(args) > 0 {
		limit, _ = strconv.Atoi(args[0])
		if limit == 0 {
			limit = 50
		}
	}

	if err := do(limit); err != nil {
		panic(err)
	}

}

func do(depthLimit int) error {
	q := []string{"."}
	currDepth := 0
	for len(q) > 0 && currDepth < depthLimit {
		path := q[0]
		q = q[1:]

		infos, err := ioutil.ReadDir(path)
		if err != nil {
			return err
		}

		for _, info := range infos {
			if info.IsDir() && !ignore(info.Name()) {
				newPath := path + "/" + info.Name()
				fmt.Println(newPath)
				q = append(q, newPath)
			}
		}
		if len(q) > 0 && depth(path) != depth(q[0]) {
			currDepth++
		}
	}
	return nil
}

func depth(path string) int {
	return len(strings.Split(path, "/"))
}

func ignore(dir string) bool {
	for _, i := range ignores {
		if dir == i {
			return true
		}
	}
	return false
}
