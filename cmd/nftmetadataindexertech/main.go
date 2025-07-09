// cmd/nftmetadataindexertech/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"nftmetadataindexertech/internal/nftmetadataindexertech"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := nftmetadataindexertech.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
