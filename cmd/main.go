package main

import (
	"fmt"
	"os"
	"strings"
	"sylva_parser/transformer"
)

func main() {
	err := execute()
	if err != nil {
		fmt.Println("Runtime error:", err)
	}
}

func execute() error {
	fileLocation := "testing.sylva"
	data, err := os.ReadFile(fileLocation)
	if err != nil {
		return fmt.Errorf("error while reading file: %v", err)
	}
	contents := string(data)
	contents = strings.ReplaceAll(contents, "\r\n", "\n")
	contents = strings.ReplaceAll(contents, "\r", "\n")

	fmt.Println("code:", contents)

	json, err := transformer.TransformJSON(contents, fileLocation, false)
	if err != nil {
		return err
	}

	fmt.Println(string(json))

	f, err := os.OpenFile("testing.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	if _, err := f.Write(json); err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}
