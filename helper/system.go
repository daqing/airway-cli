package helper

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func System(command string, debug bool) error {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	if debug {
		fmt.Println(string(output))
	}

	return nil
}

// returns: ["/foo/bar/a.txt", "/foo/bar/b.txt"]
func Glob(dir, patten string) ([]string, error) {
	files, err := filepath.Glob(fmt.Sprintf("%s/%s", dir, patten))
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return files, nil
}

func Basename(path string) string {
	return filepath.Base(path)
}
