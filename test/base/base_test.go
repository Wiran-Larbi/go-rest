package base

import (
	"fmt"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
	} else {
		fmt.Println("Current working directory:", cwd)
	}
}
