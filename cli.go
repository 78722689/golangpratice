package main

import (
	"flag"
	"fmt"
)

func usage() {
	fmt.Println("")
	fmt.Println("USAGE: sample.exe  \n")
	fmt.Println("Samples set.\n")
	fmt.Println("ARGUMENTS:\n")
	fmt.Println("    -kafka, --kafka: Kafka client calls. -help to get more help\n")
	fmt.Println("    -help, --help: Usage\n")

	flag.PrintDefaults()
}
