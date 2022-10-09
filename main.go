package main

import (
	"fmt"
	"log"

	"github.com/hsmtkk/animated-carnival/unixepoch"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Run:  run,
		Args: cobra.ExactArgs(1),
	}
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func run(cmd *cobra.Command, args []string) {
	t, err := unixepoch.StringToTime(args[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", t)
}
