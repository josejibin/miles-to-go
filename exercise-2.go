package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error enter a number", num)
                            os.Exit(0)
	}
              if num%2 == 0 {
                    fmt.Println("even", num)
              } else  {
                fmt.Println("odd")
              }
}
