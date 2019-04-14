package apigen

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//ReplaceFileContent - This finds and replace the content in a filee
func ReplaceFileContent(path string, find string, replaceWith string) {
	read, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(path)

	newContents := strings.Replace(string(read), find, replaceWith, -1)

	fmt.Println(newContents)

	err = ioutil.WriteFile(path, []byte(newContents), 0)
	if err != nil {
		panic(err)
	}
}
