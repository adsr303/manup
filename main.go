package main

import (
	"fmt"

	"github.com/adsr303/manup/manpages"
)

func main() {
	m, err := manpages.GetManpages()
	fmt.Println("manup", len(m), err)
}
