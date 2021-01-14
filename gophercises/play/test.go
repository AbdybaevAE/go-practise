package main

import (
	"flag"
	"fmt"
)

func main() {
	num := flag.Int("n", 42, "an int is ")
	flag.Parse()
	fmt.Println(*num)
}