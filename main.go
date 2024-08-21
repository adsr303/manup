package main

import (
	"fmt"
	"log"

	"github.com/adsr303/manup/manpages"
	"github.com/charmbracelet/huh"
)

func main() {
	m, err := manpages.GetManpages()
	fmt.Println("manup", len(m), err)
	browseManpages()
}

func browseManpages() {
	var section string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose section").
				Options(
					huh.NewOption("1", "1"),
					huh.NewOption("1posix", "1posix"),
					huh.NewOption("3", "3"),
				).
				Value(&section),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(section)
}
