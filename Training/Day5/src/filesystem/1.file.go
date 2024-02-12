package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {

	fmt.Println("Directory walk through")
	path := "/Users/sadanandd/Documents/"
	fssytem := os.DirFS(path)

	fs.WalkDir(fssytem, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path, d, err)
		fInfo, err := d.Info()
		fmt.Println(fInfo.ModTime())
		return nil
	})
}
