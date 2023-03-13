package main

import (
	"fmt"
	"net/http"

	"github.com/achal1304/go-httpclient/gohttp"
)

var (
	githubHttpClient = getGithubHttpClient()
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	getUrls()
	getUrls()
}

func getGithubHttpClient() gohttp.HttpClient {
	client := gohttp.New()
	commonHeader := make(http.Header)
	commonHeader.Set("Authorization", "Bearer 123_abc")

	client.SetHeaders(commonHeader)
	return client
}

func getUrls() {
	headers := make(http.Header)

	response, err := githubHttpClient.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
}

func createUser(user User) {
	headers := make(http.Header)

	response, err := githubHttpClient.Post("https://api.github.com", headers, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
}
