package manpages

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

type Manpage struct {
	Name        string `json:"name"`
	Section     string `json:"section"`
	Description string `json:"description"`
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
