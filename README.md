# Selenium Web Scraping Automation with Go

This project automates web scraping using the Selenium WebDriver with Go. The script launches a Chrome browser, navigates to a specified website, performs a search, and saves a screenshot of the result page.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Requirements

Before running this project, ensure you have the following:

- **Go** (version 1.16+)
- **ChromeDriver** (compatible with your Chrome version)
- **Google Chrome** browser
- **Selenium WebDriver for Go** installed
- **Git** (optional, to clone the repository)

## Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/devlov-LT/Selenium-Web-Automation.git
    cd Selenium-Web-Automation
    ```

2. **Install dependencies**:

    Ensure you have `chromedriver` installed and available on your machine. You can install it via Homebrew (Mac) or download it manually from [ChromeDriver official page](https://sites.google.com/chromium.org/driver/).

    For Mac users:

    ```bash
    brew install chromedriver
    ```

    Next, install the Selenium package for Go by running:

    ```bash
    go get github.com/tebeka/selenium
    ```

3. **Update the ChromeDriver path**:

    In the `main.go` file, ensure the path to `chromedriver` is correctly set:

    ```go
    path := "/opt/homebrew/bin/chromedriver"  // Update this if necessary
    ```

## Usage

1. **Modify search parameters**:

    Update the search terms and website URL in `utils.go` under the `to_search` and `website_to_check` variables:

    ```go
    var to_search = []string{
        "Selenium WebDriver",
        "deepika padukone",
    }

    var website_to_check = []string{
        "https://www.google.com",
    }
    ```

2. **Run the Go application**:

    After modifying the parameters, you can run the application with the following command:

    ```bash
    go run main.go utils.go
    ```

    The application will open a Chrome browser, search for the term, take a screenshot, and save it in the `screenshots` directory.

3. **Screenshots**:

    Screenshots will be saved with a unique name in the `screenshots/` directory, using the search string as the base filename.

## Project Structure

```bash
.
├── main.go          # Main script for Selenium interactions
├── utils.go         # Utility functions including search terms and screenshot saving
└── README.md        # Project documentation
```

### Key Functions

- **main.go**:
  - Launches ChromeDriver, navigates to the website, performs the search, and takes a screenshot.
  
- **utils.go**:
  - Contains helper functions such as `getUniqueFilename` to manage unique filenames for screenshots and search parameters.

## Contributing

Feel free to fork this repository and submit pull requests with any improvements or bug fixes!