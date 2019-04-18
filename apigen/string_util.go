package apigen

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//ReplaceFileContent - This finds and replace the content in a filee
func ReplaceFileContent(path string, find string, replaceWith string) {
	CreateFile(path)
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

// CreateFile - This func creates the file if not exist
func CreateFile(path string) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			var file, err = os.Create(path)
			if isError(err) {
				return
			}
			d1 := []byte("#Replace#")
			err = ioutil.WriteFile(path, d1, 0644)
			if isError(err) {
				return
			}
			defer file.Close()
		}
	}
	fmt.Println("==> done creating file", path)
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
