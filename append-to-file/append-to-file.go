package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	filename := "foo.txt"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for i := 0; i < 10; i++ {
		if _, err = f.WriteString(
			fmt.Sprintf("%s\n",
				time.Now(),
			)); err != nil {
			panic(err)
		}
	}
}
