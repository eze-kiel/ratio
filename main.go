package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var inc bool
	flag.BoolVar(&inc, "i", false, "Increase the size")
	flag.Parse()

	if len(flag.Args()) != 2 {
		fmt.Printf("Usage: %s [width] [height]\n", os.Args[0])
		os.Exit(1)
	}

	width, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	height, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%%\twidth\theight\n")
	if inc {
		for i := 1; i < 11; i++ {
			newW := width / 10 * i
			newH := int((float64(height) / float64(width)) * float64(newW))
			fmt.Printf("%d\t%d\t%d\n", i*10+100, newW+width, newH+height)
		}
	} else {
		for i := 1; i < 11; i++ {
			newW := width / 10 * i
			newH := int((float64(height) / float64(width)) * float64(newW))
			fmt.Printf("%d\t%d\t%d\n", i*10, newW, newH)
		}
	}
}
