// Tool gokr-rsync is an rsync receiver Go implementation.
package main

import (
	"log"
	"os"

	"github.com/superrxan/rsync/internal/receivermaincmd"
)

func main() {
	if _, err := receivermaincmd.Main(os.Args, os.Stdin, os.Stdout, os.Stderr); err != nil {
		log.Fatal(err)
	}
}
