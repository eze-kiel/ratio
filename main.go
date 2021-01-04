package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		usage(os.Args[0])
		os.Exit(1)
	}
	width, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	height, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	// ratio := width / height
	fmt.Printf("%%\twidth\theight\n")
	for i := 1; i < 11; i++ {
		newW := width / 10 * i
		newH := int((float64(height) / float64(width)) * float64(newW))
		fmt.Printf("%d\t%d\t%d\n", i*10, newW, newH)
	}
}

func usage(exec string) {
	fmt.Printf("Usage: %s [width] [height]\n", exec)
}
