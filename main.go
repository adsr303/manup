package main

import (
	"fmt"
	"log"

	"github.com/adsr303/manup/manpages"
	"github.com/charmbracelet/huh"
)

func main() {
	m, err := manpages.GetManpages()
	if err != nil {
		log.Fatalf("loading manpages list: %v", err)
	}
	browseManpages(m)
}

func browseManpages(pages []manpages.Manpage) {
	var manOpts []huh.Option[string]
	for _, page := range pages {
		manOpts = append(manOpts, huh.NewOption(
			page.Description,
			page.Name+"."+page.Section,
		))
	}

	var page string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose manpage").
				Options(manOpts...).
				Value(&page),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(page)
}
