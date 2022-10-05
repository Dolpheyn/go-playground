package main

import (
	"fmt"
	"log"
	"minigrep/minigrep"
	"os"
	"strings"
)

func main() {
	cfg, err := minigrep.NewConfig(os.Args[1:])

	if err != nil {
		log.Fatal("Error:", err)
	}

	if results, err := minigrep.Run(cfg); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} else {
		fmt.Println(strings.Join(results, "\n"))
	}
}
