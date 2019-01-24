package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cn100800/watermark/cmd"
)

var (
	version string
)

func main() {
	v := flag.Bool("v", false, "version")
	flag.Parse()
	if *v {
		fmt.Printf("version is %s \n", version)
		os.Exit(1)
	}

	cmd.CreateTextImage()
	cmd.MergeImage()
}
