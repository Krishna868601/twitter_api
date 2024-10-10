package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dghubble/oauth1"
)

// Tweet structure to hold the tweet content
type Tweet struct {
	Text string `json:"text"`
}

// Function to post a tweet
func postTweet(httpClient *http.Client, tweetText string) string {
	tweet := Tweet{
		Text: tweetText,
	}

	// Convert tweet to JSON
	tweetJSON, err := json.Marshal(tweet)
	if err != nil {
		fmt.Println("Error marshaling tweet:", err)
		return ""
	}

	// Post the tweet
	tweetURL := "https://api.twitter.com/2/tweets"
	req, err := http.NewRequest("POST", tweetURL, bytes.NewBuffer(tweetJSON))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Error while posting tweet:", err)
		return ""
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusCreated {
		var responseBody map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
			fmt.Println("Error decoding response body:", err)
			return ""
		}
		fmt.Println("Error while posting tweet:", resp.Status, responseBody)
		return ""
	}

	// Print success response
	var successResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&successResponse); err != nil {
		fmt.Println("Error decoding success response:", err)
		return ""
	}
	tweetID := successResponse["data"].(map[string]interface{})["id"].(string)
	fmt.Println("Tweet posted successfully! Tweet ID:", tweetID)

	return tweetID
}

// Function to delete a tweet using the provided tweet ID
func deleteTweet(httpClient *http.Client, tweetID string) {
	// Delete the tweet
	deleteURL := fmt.Sprintf("https://api.twitter.com/2/tweets/%s", tweetID)
	req, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Send the request
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Error while deleting tweet:", err)
		return
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode == http.StatusOK {
		var responseBody map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
			fmt.Println("Error decoding response body:", err)
			return
		}
		deleted := responseBody["data"].(map[string]interface{})["deleted"].(bool)
		if deleted {
			fmt.Println("Tweet deleted successfully!")
		} else {
			fmt.Println("Tweet deletion failed.")
		}
	} else {
		var responseBody map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
			fmt.Println("Error decoding response body:", err)
			return
		}
		fmt.Println("Error while deleting tweet:", resp.Status, responseBody)
	}
}

func main() {
	// Load environment variables
	apiKey := os.Getenv("TWITTER_API_KEY")
	apiSecretKey := os.Getenv("TWITTER_API_SECRET_KEY")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

	// Set up OAuth 1.0a authentication
	config := oauth1.NewConfig(apiKey, apiSecretKey)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	// Create an HTTP client
	httpClient := config.Client(oauth1.NoContext, token)

	// Command-line flag to specify if the user wants to delete a tweet
	deleteFlag := flag.Bool("delete", false, "Use this flag to delete a tweet by providing the tweet ID")
	tweetIDFlag := flag.String("tweetID", "", "Tweet ID to delete")

	// Parse the flags
	flag.Parse()

	// Check if the delete flag is set
	if *deleteFlag {
		if *tweetIDFlag == "" {
			fmt.Println("Please provide the tweet ID to delete using the -tweetID flag.")
			return
		}
		// Call the delete function with the provided tweet ID
		deleteTweet(httpClient, *tweetIDFlag)
	} else {
		// Generate unique tweet text with timestamp
		tweetText := fmt.Sprintf("Hello from Twitter API! %v", time.Now().Format("2006-01-02 15:04:05"))
		tweetID := postTweet(httpClient, tweetText)
		if tweetID == "" {
			fmt.Println("Failed to post tweet.")
		}
	}
}
