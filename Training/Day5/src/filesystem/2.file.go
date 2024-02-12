package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Directory Operations")

	dirName := "TestDirectory"

	dir := "./1/2/3/4/testDir"

	os.Mkdir(dirName, 0777)

	os.MkdirAll(dir, 0777)

	os.Remove(dir)
	os.RemoveAll("./1")

}
