package goengine

import (
	"fmt"
	"os"
	"strings"
)

func OpenTextFile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func Find(str, find string, index int) int {
	if index < 0 || index >= len(str) {
		return -1
	}
	substr := str[index:]
	return strings.Index(substr, find)
}
