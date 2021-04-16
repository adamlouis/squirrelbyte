package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var (
	flagJob string = ""
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: strava [auth|run]")
	}

	// https://developers.strava.com/docs/getting-started/
	if os.Args[1] == "auth" {
		clientID := os.Getenv("STRAVA_CLIENT_ID")
		clientSecret := os.Getenv("STRAVA_CLIENT_SECRET")
		redirectURI := "http://localhost:9923/oauth/redirect"
		state := getRandomString(20)

		authorizeURL := fmt.Sprintf(
			"https://www.strava.com/api/v3/oauth/authorize?client_id=%s&response_type=code&redirect_uri=%s&approval_prompt=force&scope=read_all&state=%s",
			clientID,
			redirectURI,
			state,
		)

		fmt.Printf("OPEN THIS URL IN A WEB BROWSER:\n\n%s\n\n", authorizeURL)

		// Create a new redirect route route
		http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query()
			code := q["code"][0]
			redirectState := q["state"][0]

			if redirectState != state {
				panic("redirect state does not match origin state")
			}

			data := url.Values{}
			data.Set("client_id", clientID)
			data.Set("client_secret", clientSecret)
			data.Set("code", code)
			data.Set("grant_type", "authorization_code")

			client := &http.Client{}
			tokenRequest, err := http.NewRequest("POST", "https://www.strava.com/api/v3/oauth/token", strings.NewReader(data.Encode()))
			if err != nil {
				log.Fatal(err)
			}
			tokenRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			tokenRequest.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

			res, err := client.Do(tokenRequest)
			if err != nil {
				panic(err)
			}

			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			type Athlete struct {
				FirstName string `json:"firstname"`
				LastName  string `json:"lastname"`
				Profile   string `json:"profile"`
			}
			type TokenResponse struct {
				AccessToken  string  `json:"access_token"`
				RefreshToken string  `json:"refresh_token"`
				Athlete      Athlete `json:"athlete"`
			}

			tokenResponse := &TokenResponse{}
			json.Unmarshal(b, tokenResponse)

			w.WriteHeader(200)
			w.Header().Add("Content-Type", "text/html; charset=UTF-8")
			w.Write([]byte(fmt.Sprintf(`<!DOCTYPE html><html>Hello %s %s!<br /><img src="%s"></html>`, tokenResponse.Athlete.FirstName, tokenResponse.Athlete.LastName, tokenResponse.Athlete.Profile)))

			fmt.Printf("ACCESS_TOKEN=%s\n", tokenResponse.AccessToken)
			fmt.Printf("REFRESH_TOKEN=%s\n", tokenResponse.RefreshToken)
		})

		http.ListenAndServe(":9923", nil)
	}

	if os.Args[1] == "run" {
		fmt.Println("run!")
	}
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
