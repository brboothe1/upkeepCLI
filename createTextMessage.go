package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	i        int
	textBody string
	// Daily tasks to send out in text messages
	morningTasks   string = "Morning Tasks: " + "\n" + "1. Make your bed when you wake up" + "\n" + "2. Drink a glass of water" + "\n" + "3. Record daily commitment" + "\n"
	afternoonTasks string = "Afternoon Tasks: " + "\n" + "1. Go for a walk" + "\n" + "2. Have a conversation with someone" + "\n" + "3. Drink water" + "\n"
	eveningTasks   string = "Evening Tasks: " + "\n" + "1. Work towards your goals" + "\n" + "2. Be thankful for the day" + "\n" + "3. Reflect on what you learned" + "\n"
)

func createTextMessage() {

	// OAuth2 login information
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Twilio Messaging Information
	msgData := url.Values{}
	msgData.Set("To", "+19548823282")
	msgData.Set("From", "+18783484989")

	// Switch the body of the text message to the correct task for that time of day.
	i++
	switch i {
	case 1:
		textBody = morningTasks
	case 2:
		textBody = afternoonTasks
	case 3:
		textBody = eveningTasks
		i = 0
	}
	msgData.Set("Body", textBody)
	msgDataReader := *strings.NewReader(msgData.Encode())

	// Client setup with auth information for Twilio
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Check status code. A status code beginning with a 2 means success.
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
}
