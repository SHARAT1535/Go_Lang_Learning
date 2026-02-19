package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type notxterror struct {
	message string
}

// Implement Error() method so it satisfies error interface
func (e *notxterror) Error() string {
	return e.message
}

func loadmsg(filename string) (string, error) {

	if !strings.HasSuffix(filename, ".txt") {
		return "", &notxterror{"opening non txt file not allowed"}
	}

	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return filename, nil
}

func main() {

	s, err := loadmsg("go.mod")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist")
	}
	if err != nil {
		if _, ok := err.(*notxterror); ok {
			fmt.Println("Custom Error:", err)
		} else {
			fmt.Println("An unexpected error occurred:", err)
		}
	} else {
		fmt.Println("File opened successfully:", s)
	}
}
