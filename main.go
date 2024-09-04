package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

func main() {

	searching_string := to_search[2]
	checking_website := website_to_check[0]

	hub := "http://127.0.0.1:9515/wd/hub"

	path := "/opt/homebrew/bin/chromedriver"
	opts := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(path, 9515, opts...)
	if err != nil {
		log.Fatalf("Error in starting ChromeDriver: %v", err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	wd, err := selenium.NewRemote(caps, hub)
	if err != nil {
		log.Fatalf("Error in starting the ChromeDrvier: %v", err)
	}
	defer wd.Quit()

	if err := wd.Get(checking_website); err != nil {
		log.Fatalf("Failed to load page: %v", err)
	}

	searchBox, err := wd.FindElement(selenium.ByCSSSelector, "textarea[name='q']")
	if err != nil {
		// log.Fatalf("here - 1")
		log.Fatalf("Failed to find search box: %v", err)
	}

	// log.Fatalf("here - 2")

	// err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
	// 	elem, err := wd.FindElement(selenium.ByCSSSelector, "textarea[name='q']")
	// 	if err != nil {
	// 		return false, nil
	// 	}
	// 	return elem.IsDisplayed()
	// }, 10*time.Second)

	// log.Fatalf("here - 3")

	if err != nil {
		// log.Fatalf("here - 4")
		log.Fatalf("Failed to wait: %v", err)
	}

	if err := searchBox.SendKeys(searching_string); err != nil {
		log.Fatalf("Failed to enter search query: %v", err)
	}

	if err := searchBox.SendKeys(selenium.EnterKey); err != nil {
		log.Fatalf("Failed to submit: %v", err)
	}

	time.Sleep(2 * time.Second)

	title, err := wd.Title()
	if err != nil {
		log.Fatalf("Failed to get the page tile: %v", err)
	}

	fmt.Printf("Page title is: %s\n", title)

	SaveScreenshot(wd, searching_string)
}

func SaveScreenshot(wd selenium.WebDriver, searching_string string) {

	screenshot, err := wd.Screenshot()
	if err != nil {
		log.Fatalf("Failed to take ss: %v", err)
	}

	if err := os.MkdirAll("screenshots", os.ModePerm); err != nil {
		log.Fatalf("Failed to create screenshots directory: %v", err)
	}

	filename := getUniqueFilename("screenshots/"+searching_string, ".png")

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create a screenshot: %v", err)
	}
	defer file.Close()

	if _, err := file.Write(screenshot); err != nil {
		log.Fatalf("Failed to save ss: %v", err)
	}

	fmt.Printf("Screenshot saved as %s\n", filename)
}
