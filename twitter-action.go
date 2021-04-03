package main

import (
	"flag"
	"log"
	"os"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "", "Twitter Access Token")
	accessSecret := flags.String("access-secret", "", "Twitter Access Secret")
	tweetMessage := flags.String("message", "", "Tweet Message")
	dryRun := flags.Bool("dry", false, "Test mode, nothing will be sent to twitter")
	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")

	// Validating the credentials are available (unless dryRun)
	if (*consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "") && !*dryRun {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	tweetContent := string(*tweetMessage)

	// Validation a content is available and does not exeed 280 char
	if len(tweetContent) == 0 {
		log.Fatal("Your tweet is empty!")
	} else if len(tweetContent) > 280 {
		log.Fatal("Tweet must be less than 280 char")
	}

	// Posting tweet
	if *dryRun {
		log.Print("Logging in, creating client and updating status.")
	} else {
		var err error
		// Setup auth
		config := oauth1.NewConfig(*consumerKey, *consumerSecret)
		token := oauth1.NewToken(*accessToken, *accessSecret)

		// http.Client will automatically authorize Requests
		httpClient := config.Client(oauth1.NoContext, token)

		// Twitter client
		client := twitter.NewClient(httpClient)
		_, _, err = client.Statuses.Update(tweetContent, nil)
		// Handling Error
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("Status updated with: " + tweetContent)
}
