<img width="622" height="229" alt="image" src="https://github.com/user-attachments/assets/0b403fd0-9c5c-4c56-80b6-0a34f1bf8839" />

# GitHub User Info Fetcher

A simple Go program that fetches and displays GitHub user information using the GitHub API.

## Features

- Retrieves and displays user’s name, company, blog, location, number of public repos, number of private repos, followers, and following.
- Shows the account creation date.
- Nicely formatted output with emojis for better readability.

## Usage

1. Clone the repository or download the code.
2. Install Go on your system.
3. Run the program in the terminal by providing a GitHub username as an argument:

    go run main.go <github-username>

Example:

    go run main.go sudo-1n17

## Requirements

- Go 1.15 or higher
- Internet connection to access the GitHub API

## Program Structure

- `main.go`: The main file that sends HTTP requests to the GitHub API and processes the data.
- Uses `net/http` package to send requests.
- Uses `encoding/json` package to decode JSON data into Go structs.
  
---

