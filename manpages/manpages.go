package manpages

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"slices"
	"strings"
)

type Manpage struct {
	Name        string
	Section     string
	Description string
}

func GetManpages() ([]Manpage, error) {
	cmd := exec.Command("whatis", "--wildcard", "*", "--long")
	b, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("whatis: %w", err)
	}
	var result []Manpage
	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		m := whatisLine.FindStringSubmatch(scanner.Text())
		if len(m) != 4 {
			return nil, fmt.Errorf("unexpected content: %s", scanner.Text())
		}
		result = append(result, Manpage{
			Name:        m[1],
			Section:     m[2],
			Description: m[3],
		})
	}
	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("reading whatis output: %w", err)
	}
	return result, nil
}

var whatisLine = regexp.MustCompile(`^(\S+) \(([0-9a-z][a-z]*)\) +- (.+)$`)

type Section struct {
	Id    string
	Intro string
}

func GetSections(pages []Manpage) []Section {
	sections := make(map[string]bool)
	intros := make(map[string]string)
	for _, m := range pages {
		sections[m.Section] = true
		if m.Name == "intro" {
			intros[m.Section] = m.Description
		}
	}
	var result []Section
	for k := range sections {
		intro := intros[k]
		result = append(result, Section{Id: k, Intro: intro})
	}
	slices.SortFunc(result, func(a, b Section) int {
		return strings.Compare(a.Id, b.Id)
	})
	return result
}
