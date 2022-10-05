// Package minigrep provides configuration and execution of minigrep
package minigrep

import (
	"errors"
	"os"
	"strings"
)

type Config struct {
	Needle          string
	Filename        string
	CaseInsensitive bool
}

// NewConfig constructs a Config struct from os.Args
func NewConfig(args []string) (Config, error) {
	if len(args) == 0 {
		return Config{}, errors.New("<needle> is required!")
	}

	if len(args) == 1 {
		return Config{}, errors.New("<filename> is required!")
	}

	needle := args[0]
	filename := args[1]
	caseInsensitive := os.Getenv("CASE_INSENSITIVE") == "true"

	return Config{needle, filename, caseInsensitive}, nil
}

func Run(cfg Config) ([]string, error) {
	haystackBytes, err := os.ReadFile(cfg.Filename)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	haystack := string(haystackBytes[:])
	if cfg.CaseInsensitive {
		return SearchCaseInsensitive(&cfg.Needle, &haystack), nil
	} else {
		return Search(&cfg.Needle, &haystack), nil
	}
}

func Search(needle *string, haystack *string) (foundLines []string) {
	for _, s := range strings.Split(*haystack, "\n") {
		if strings.Contains(s, *needle) {
			foundLines = append(foundLines, s)
		}
	}

	return
}

func SearchCaseInsensitive(needle *string, haystack *string) (foundLines []string) {
	for _, line := range strings.Split(*haystack, "\n") {
		if strings.Contains(strings.ToLower(line), strings.ToLower(*needle)) {
			foundLines = append(foundLines, line)
		}
	}

	return
}
