package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

func clientFromFile(filename string) (*http.Client, error) {
	creds, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Unable to open %q: %w", filename, err)
	}
	// drive.DriveScope is required for team drive access
	config, err := google.ConfigFromJSON(creds, drive.DriveScope)
	if err != nil {
		return nil, fmt.Errorf("Unable to initialize auth config: %w", err)
	}

	token := &oauth2.Token{}

	// Attempt to read the cached token file.
	bytes, err := ioutil.ReadFile(".token")
	if err == nil {
		// File found, try to parse as JSON.
		err = json.Unmarshal(bytes, token)
	}
	// Either file not found, not not containing correct JSON.
	if err != nil {
		token = getTokenFromWeb(config)
		cache := []byte{}
		if cache, err = json.Marshal(token); err == nil {
			err = ioutil.WriteFile(".token", cache, 0600)
		}
	}
	if err != nil {
		return nil, fmt.Errorf("Unable to initialize client: %w", err)
	}

	return config.Client(context.Background(), token), nil
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}
