package pluginstrava

var (
	flagJob string = ""
)

// func ImportActivitySummariesWorkerFn(ctx context.Context, job *worker.Job) error {

// }

// func ImportActivitySummaries() {

// }

// func main() {
// 	if len(os.Args) < 2 {
// 		log.Fatal("usage: strava [auth|run]")
// 	}

// 	// https://developers.strava.com/docs/getting-started/
// 	if os.Args[1] == "auth" {
// 		clientID := os.Getenv("STRAVA_CLIENT_ID")
// 		clientSecret := os.Getenv("STRAVA_CLIENT_SECRET")
// 		redirectURI := "http://localhost:9923/oauth/redirect"
// 		state := getRandomString(20)

// 		authorizeURL := fmt.Sprintf(
// 			"https://www.strava.com/api/v3/oauth/authorize?client_id=%s&response_type=code&redirect_uri=%s&approval_prompt=force&scope=activity:read_all&state=%s",
// 			clientID,
// 			redirectURI,
// 			state,
// 		)

// 		fmt.Printf("OPEN THIS URL IN A WEB BROWSER:\n\n%s\n\n", authorizeURL)

// 		// Create a new redirect route route
// 		http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
// 			q := r.URL.Query()
// 			code := q["code"][0]
// 			redirectState := q["state"][0]

// 			if redirectState != state {
// 				panic("redirect state does not match origin state")
// 			}

// 			data := url.Values{}
// 			data.Set("client_id", clientID)
// 			data.Set("client_secret", clientSecret)
// 			data.Set("code", code)
// 			data.Set("grant_type", "authorization_code")

// 			client := &http.Client{}
// 			tokenRequest, err := http.NewRequest("POST", "https://www.strava.com/api/v3/oauth/token", strings.NewReader(data.Encode()))
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			tokenRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
// 			tokenRequest.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

// 			res, err := client.Do(tokenRequest)
// 			if err != nil {
// 				panic(err)
// 			}

// 			if err != nil {
// 				panic(err)
// 			}
// 			defer res.Body.Close()

// 			b, err := ioutil.ReadAll(res.Body)
// 			if err != nil {
// 				panic(err)
// 			}

// 			type Athlete struct {
// 				FirstName string `json:"firstname"`
// 				LastName  string `json:"lastname"`
// 				Profile   string `json:"profile"`
// 			}
// 			type TokenResponse struct {
// 				AccessToken  string  `json:"access_token"`
// 				RefreshToken string  `json:"refresh_token"`
// 				Athlete      Athlete `json:"athlete"`
// 			}

// 			tokenResponse := &TokenResponse{}
// 			json.Unmarshal(b, tokenResponse)

// 			w.WriteHeader(200)
// 			w.Header().Add("Content-Type", "text/html; charset=UTF-8")
// 			w.Write([]byte(fmt.Sprintf(`<!DOCTYPE html><html>Hello %s %s!<br /><img src="%s"></html>`, tokenResponse.Athlete.FirstName, tokenResponse.Athlete.LastName, tokenResponse.Athlete.Profile)))

// 			fmt.Printf("ACCESS_TOKEN=%s\n", tokenResponse.AccessToken)
// 			fmt.Printf("REFRESH_TOKEN=%s\n", tokenResponse.RefreshToken)
// 		})

// 		http.ListenAndServe(":9923", nil)
// 	}

// 	if os.Args[1] == "run" {
// 		client := &http.Client{}
// 		accessToken := os.Getenv("STRAVA_ACCESS_TOKEN")
// 		fmt.Println(accessToken)

// 		page := 1
// 		perPage := 30

// 		for {

// 			url := fmt.Sprintf("https://www.strava.com/api/v3/athlete/activities?page=%d&per_page=%d", page, perPage)

// 			req, err := http.NewRequest("GET", url, nil)
// 			if err != nil {
// 				panic(err)
// 			}
// 			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
// 			res, _ := client.Do(req)

// 			b, err := ioutil.ReadAll(res.Body)
// 			if err != nil {
// 				panic(err)
// 			}

// 			type Activity struct {
// 				ID int64 `json:"id"`
// 			}

// 			activities := []*Activity{}
// 			activitiesIFC := []interface{}{}
// 			json.Unmarshal(b, &activities)
// 			json.Unmarshal(b, &activitiesIFC)

// 			for i, a := range activities {
// 				header := map[string]string{
// 					"source":    "strava",
// 					"type":      "summary-activity",
// 					"strava_ur": fmt.Sprintf("https://www.strava.com/activities/%d", a.ID),
// 				}
// 				sqbtPost(fmt.Sprintf("strava.summary-activity.%d", a.ID), activitiesIFC[i], header)
// 			}

// 			if len(activities) == 0 {
// 				break
// 			}

// 			page++
// 		}
// 	}
// }

// const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// func getRandomString(n int) string {
// 	b := make([]byte, n)
// 	for i := range b {
// 		b[i] = letters[rand.Intn(len(letters))]
// 	}
// 	return string(b)
// }

// func ifcToBuf(ifc interface{}) (*bytes.Buffer, error) {
// 	bs, err := json.Marshal(ifc)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bytes.NewBuffer(bs), nil
// }

// func sqbtPost(id string, body interface{}, header interface{}) error {
// 	buf, err := ifcToBuf(
// 		map[string]interface{}{
// 			"id":     id,
// 			"header": header,
// 			"body":   body,
// 		})
// 	if err != nil {
// 		return err
// 	}

// 	res, err := http.Post("http://localhost:9922/api/documents", "application/json", buf)
// 	if err != nil {
// 		return err
// 	}

// 	b, _ := ioutil.ReadAll(res.Body)
// 	fmt.Println(string(b))

// 	return nil
// }
