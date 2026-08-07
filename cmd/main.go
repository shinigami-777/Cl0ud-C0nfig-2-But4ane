package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"Cl0ud-C0nfig-2-But4ane/internal/converter"
	"gopkg.in/yaml.v3"
)

func main() {
	inputFile := flag.String("f", "", "Input cloud-config file (default: stdin)")
	outputFile := flag.String("o", "", "Output Butane file (default: stdout)")
	flag.Parse()

	var input []byte
	var err error

	if *inputFile == "" || *inputFile == "-" {
		input, err = io.ReadAll(os.Stdin)
	} else {
		input, err = os.ReadFile(*inputFile)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	var cc converter.CloudConfig
	if err := yaml.Unmarshal(input, &cc); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing cloud-config YAML: %v\n", err)
		os.Exit(1)
	}

	bc := converter.Convert(&cc)

	outBytes, err := yaml.Marshal(bc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling Butane YAML: %v\n", err)
		os.Exit(1)
	}

	if *outputFile == "" || *outputFile == "-" {
		os.Stdout.Write(outBytes)
	} else {
		if err := os.WriteFile(*outputFile, outBytes, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
			os.Exit(1)
		}
	}
}
