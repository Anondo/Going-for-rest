package main

import (
	"going_rest/cmd"
	"log"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
