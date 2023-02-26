package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"test/config"
	"test/model"

	"golang.org/x/oauth2"
)

func init() {
	// load config from env file
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
}

func main() {
	conf := ConfigData()
	oauthURL := conf.AuthCodeURL(config.Cfg.State)
	fmt.Println(oauthURL)
	http.HandleFunc("/login/oauth2/code/goto", OAuth)
	fmt.Println("Server started at successfully!")
	http.ListenAndServe(":5000", nil)
}

func OAuth(w http.ResponseWriter, req *http.Request) {
	oauthData := ConfigData()
	state := req.URL.Query().Get("state")
	if state != config.Cfg.State {
		log.Fatalln("Ignoring authorization code with unexpected state")
		return
	}
	authorizationCode := req.URL.Query().Get("code")
	token, err := oauthData.Exchange(context.TODO(), authorizationCode)
	if err != nil {
		fmt.Println("Access Token Error:", err)
		return
	}
	data, err := ReceiveWebhookData()
	if err != nil {
		log.Fatalln(err)
	}
	SendMessage(data, strings.TrimSpace(token.AccessToken))

}

func ReceiveWebhookData() (*model.ContentData, error) {
	resp, err := http.Get("https://webhook.site/token/" + config.Cfg.TokenID + "/requests?sorting=newest")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	newBody := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(string(body), "\"{", "{"), "}\"", "}"), "\\\"", "\"")
	var data model.ContentData
	err = json.Unmarshal([]byte(newBody), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func SendMessage(webhookData *model.ContentData, accessToken string) {
	for _, v := range webhookData.Data {

		msg := fmt.Sprintf("Hi %s,\nI show your email address is %s\n", v.Content.User.FirstName, v.Content.User.Email)
		fmt.Printf("Sending SMS: \n%v", msg)

		url := "https://api.jive.com/messaging/v1/messages"
		method := "POST"
		value := map[string]string{
			"ownerPhoneNumber":    config.Cfg.OwnerPhoneNumber,
			"contactPhoneNumbers": v.Content.User.Phone,
			"body":                msg,
		}
		payload, err := json.Marshal(value)
		if err != nil {
			fmt.Println(err)
			return
		}

		client := &http.Client{}
		req, err := http.NewRequest(method, url, strings.NewReader(string(payload)))
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("content-type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()
	}
}

func ConfigData() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.Cfg.ClientID,
		ClientSecret: config.Cfg.ClientSecret,
		RedirectURL:  config.Cfg.RedirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://authentication.logmeininc.com/oauth/authorize",
			TokenURL: "https://authentication.logmeininc.com/oauth/token",
		},
		Scopes: []string{"messaging.v1.send"},
	}
}
