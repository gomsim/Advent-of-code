package exit

import (
	"fmt"
	"os"
)

func If(err error) {
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func Error(message string) {
	fmt.Printf("Error: %s\n", message)
	os.Exit(1)
}

func Errorf(format string, v ...any) {
	fmt.Printf("Error: %s\n", fmt.Sprintf(format, v...))
	os.Exit(1)
}
