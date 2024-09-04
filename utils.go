package main

import (
	"fmt"
	"os"
)

var (
	to_search = []string{
		"Selenium WebDriver",
		"deepika padukone",
		"sanchit agarwal",
	}

	website_to_check = []string{
		"https://www.google.com",
		"https://www.cashfree.com/customers",
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
