package main

import (
	"fmt"
	"log"

	"test/quad"
)

func main() {
	r1, r2, ok, err := quad.CalculateRoots(1, 2, 0)
	if err != nil {
		log.Fatal(err)
	}

	if ok {
		fmt.Println("корни:", r1, r2)
	} else {
		fmt.Println("нет корней")
	}
}
