package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"

	"github.com/paulhammond/msg/internal/msg"
)

func main() {
	cmd := pflag.NewFlagSet("msg", pflag.ExitOnError)
	var help *bool = cmd.BoolP("help", "h", false, "display help")
	var version *bool = cmd.BoolP("version", "v", false, "display version")

	err := cmd.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		usage(2)
	}
	if *help || cmd.Arg(0) == "help" {
		usage(0)
	}

	if *version {
		fmt.Printf("msg version %s\n", msg.Version())
		os.Exit(0)
	}

	args := cmd.Args()
	if len(args) != 2 {
		usage(2)
	}

	err = msg.Run(args[0], args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
	}
}

func usage(code int) {
	fmt.Fprintln(os.Stderr, "usage: msg <config> <output>")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Arguments:")
	fmt.Fprintln(os.Stderr, "<config>    Config File")
	fmt.Fprintln(os.Stderr, "<output>    Output directory")
	os.Exit(code)
}
