package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type cont struct {
}

func (cont) Write(bs []byte) (int, error) {
	s := string(bs)
	fmt.Println(s)

	return len(bs), nil
}
func main() {
	file, err := os.Open(os.Args[1]) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	content := cont{}

	io.Copy(content, file)

}
