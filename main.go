package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a number\n", max)
	}

	if len(os.Args) < 2 {
		fmt.Println("Give me the size of the table.")
	}

	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}

}
