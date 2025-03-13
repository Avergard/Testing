package main

import (
	"errors"
	"fmt"
	"strconv"
)

func err(a, c int) (int, error) {
	d := (a + c) % 2
	if d == 0 {
		return 0, nil
	} else {
		return d, errors.New(strconv.Itoa(a) + " " + strconv.Itoa(c))
	}
}

func main() {
	a := 12
	c := 11
	fmt.Println(a + c)
	fmt.Println(err(a, c))

}
