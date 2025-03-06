package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteToFile(args ...string) {
	if len(args) < 2 {
		return
	}
	// filename := args[0]
	filename := filepath.Clean("sklog.txt")
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Problem Opening the file", err.Error())
	}
	defer file.Close()
	for _, v := range args {
		if _, err := file.WriteString(CleanNewLines(v)); err != nil {
			fmt.Println("Error:", err.Error())
			fmt.Println("string was:", v)
			fmt.Println()
		}
		file.WriteString(string("\n"))
	}
	file.WriteString(string("------------------------------------\n"))
}
