package main

import (
	"fmt"
	"needsync/internal/needs"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli sync")
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "sync":
		if err := needs.SyncCreate("./needs"); err != nil {
			panic(err)
		}
	default:
		fmt.Println("Unknown command")
	}
}
