package main

import (
	"fmt"
	"os"
)

var (
	// mention what to act
	to_search = []string{
		"Selenium WebDriver",
		"deepika padukone",
	}

	// mention which site to open for testing
	website_to_check = []string{
		"https://www.google.com",
	}
)

func getUniqueFilename(base, ext string) string {
	filename := base + ext
	counter := 1

	for {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			break
		}
		filename = fmt.Sprintf("%s-%d%s", base, counter, ext)
		counter++
	}

	return filename
}
