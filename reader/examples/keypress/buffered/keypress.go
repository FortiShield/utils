package main

import (
	"log"
	"sync"
	"time"

	"github.com/khulnasoft-labs/utils/reader"
	stringsutil "github.com/khulnasoft-labs/utils/strings"
)

func main() {
	stdr := reader.KeyPressReader{
		Timeout: time.Duration(5 * time.Second),
		Once:    &sync.Once{},
	}

	stdr.Start()
	defer stdr.Stop()

	for {
		data := make([]byte, stdr.BufferSize)
		n, err := stdr.Read(data)
		log.Println(n, err)

		if stringsutil.IsCTRLC(string(data)) {
			break
		}
	}
}
