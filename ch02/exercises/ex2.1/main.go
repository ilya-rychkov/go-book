package main

import (
	"fmt"

	"ex2.1/tempconv"
)

func main() {
	k := tempconv.Kelvin(19)
	fmt.Println(tempconv.KToC(k))

}
