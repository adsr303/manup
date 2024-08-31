package main

import (
	"log"

	"github.com/adsr303/manup/manpages"
	"github.com/charmbracelet/huh"
)

func main() {
	m, err := manpages.GetManpages()
	if err != nil {
		log.Fatalf("loading manpages list: %v", err)
	}
	selected := browseManpages(m)
	err = manpages.ShowManpage(selected)
	if err != nil {
		log.Fatalf("showing man %s: %v", selected, err)
	}
}

func browseManpages(pages []manpages.Manpage) string {
	var manOpts []huh.Option[string]
	for _, page := range pages {
		manOpts = append(manOpts, huh.NewOption(
			trimToEllipsis(page.Description),
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

	return page
}
