package main

import (
	"fmt"
	"needsync/internal/needs"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli sync")
		return
	}

	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	cmd := os.Args[1]

	switch cmd {
	case "sync":
		dir := "./example/needs"

		// Debug output
		absPath, _ := filepath.Abs(dir)
		fmt.Printf("[DEBUG] Working directory: %s\n", os.Getenv("PWD"))
		fmt.Printf("[DEBUG] Looking for needs in: %s\n", dir)
		fmt.Printf("[DEBUG] Absolute path: %s\n", absPath)

		// Check if directory exists
		info, err := os.Stat(dir)
		if err != nil {
			fmt.Printf("[DEBUG] Error accessing directory: %v\n", err)
		} else {
			fmt.Printf("[DEBUG] Directory exists: %v, IsDir: %v\n", dir, info.IsDir())
		}

		if err := needs.SyncCreate(dir); err != nil {
			panic(err)
		}
	default:
		fmt.Println("Unknown command")
	}
}
