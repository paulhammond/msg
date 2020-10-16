package main

import (
	"fmt"
	"os"

	"github.com/paulhammond/msg/internal/msg"
)

func main() {
	err := msg.Run(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
	}
}
