package main

import (
	"ex2.1/tempconv"
	"fmt"
)

func main() {
	k := tempconv.Kelvin(19)
	fmt.Println(tempconv.KToC(k))

}
