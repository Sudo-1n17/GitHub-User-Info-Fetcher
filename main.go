package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type GitHubUser struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	Company     string `json:"company"`
	Blog        string `json:"blog"`
	Location    string `json:"location"`
	Creation    string `json:"created_at"`
	PrivetRepos int    `json:"total_private_repos"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
}

func main() {
	// check user arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage : go run main.go <userName>")
	}

	// Use Argument for url (github_api)
	UserName := os.Args[1]
	url := fmt.Sprintf("https://api.github.com/users/%s", UserName)

	// send request
	response, err := http.Get(url)

	// Handle errors
	if err != nil {
		fmt.Println("Error Fetching Data : ", err)
		return
	}
	defer response.Body.Close()

	// check http status code
	if response.StatusCode != 200 {
		fmt.Println("Faild to get user info : ", response.StatusCode)
		return
	}

	// Convert Json to Struct
	var user GitHubUser
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		fmt.Println("Error Decoding response : ", err)
		return
	}

	// print datas
	fmt.Printf("👤 GitHub User: %s\n", user.Login)
	fmt.Printf("📅 Account Created At: %s\n", user.Creation)
	fmt.Printf("🧑 Name: %s\n", user.Name)
	fmt.Printf("🏢 Company: %s\n", user.Company)
	fmt.Printf("🌍 Location: %s\n", user.Location)
	fmt.Printf("📝 Blog: %s\n", user.Blog)
	fmt.Printf("🔒 Private Repos: %d\n", user.PrivetRepos)
	fmt.Printf("📦 Public Repos: %d\n", user.PublicRepos)
	fmt.Printf("👥 Followers: %d | Following: %d\n", user.Followers, user.Following)
}
