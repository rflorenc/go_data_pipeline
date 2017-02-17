package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rflorenc/go_data_pipeline/streamer"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " <config_file>")
		os.Exit(2)
	}
	var cfg, err = streamer.LoadProperties(os.Args[1])

	if err != nil {
		fmt.Printf("Error occurred reading the properties file %s [err: %s]\n", os.Args[1], err)
		os.Exit(2)
	}
	log.Printf("Loaded config: %s\n", cfg.ToString())
	fmt.Println(cfg)
}
