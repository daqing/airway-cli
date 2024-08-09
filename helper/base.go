package helper

import (
	"fmt"
	"os"
	"text/template"
)

func pathExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}

// Execute the data with template from path and write to out
func ExecTemplate(txt string, out string, data any) error {
	if pathExists(out) {
		return os.ErrExist
	}

	return execTemplate(txt, out, data)
}

func ExecTemplateForce(path string, out string, data any) error {
	if !pathExists(out) {
		return os.ErrExist
	}

	if !pathExists(path) {
		return fmt.Errorf("template %s does not exist", path)
	}

	return execTemplate(path, out, data)
}

func execTemplate(txt, out string, data any) error {
	tmpl, err := template.New("test").Parse(string(txt))
	if err != nil {
		return err
	}

	outputFile, err := os.Create(out)
	if err != nil {
		panic(err)
	}

	defer outputFile.Close()

	return tmpl.Execute(outputFile, data)
}

func Help(msg string) {
	fmt.Printf("usage: %s\n", msg)
	os.Exit(1)
}
