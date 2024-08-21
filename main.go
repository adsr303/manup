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
	sections := manpages.GetSections(m)
	browseManpages(sections)
}

func browseManpages(sections []manpages.Section) {
	var secOpts []huh.Option[string]
	for _, sec := range sections {
		secOpts = append(secOpts, huh.NewOption(
			formatSectionKey(sec),
			sec.Id,
		))
	}

	var section string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose section").
				Options(secOpts...).
				Value(&section),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(section)
}

func formatSectionKey(sec manpages.Section) string {
	if sec.Intro != "" {
		return fmt.Sprintf("%s - %s", sec.Id, sec.Intro)
	}
	return sec.Id
}
