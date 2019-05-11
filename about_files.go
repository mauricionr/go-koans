package go_koans

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func aboutFiles() {
	filename := "about_files.go"
	contents, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(contents), "\n")
	fmt.Println(lines[5])
	assert(lines[5] == "	\"strings\"") // handling files is too trivial
}
