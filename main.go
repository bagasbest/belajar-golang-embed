package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed version.txt
var version string

//go:embed logo_next.png
var logo []byte

//go:embed files/*.txt
var path embed.FS

// untuk build di terminal go build
// terus lanjutkan dengan .\belajar-golang-embed.exe
func main() {
	fmt.Println(version)

	err := os.WriteFile("logo_next_2.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
