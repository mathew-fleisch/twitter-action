# twitter-action

This repo was forked so I could release just the go-binary to use this tweet functionality in other contexts than just a github-action. I have also removed the file input methods for simplicity.

## Auth
This process to get twitter api keys took a few days of back and forth emails explaining what this app would do.
About the authentication see: https://developer.twitter.com/en/apps
create an account, create an app
@see https://apps.twitter.com/

### retrieve the access tokens
@see https://developer.twitter.com/en/apps

# Build
```
go get .
go build .
```

## Usage
```
export TWITTER_CONSUMER_KEY=xxx
export TWITTER_CONSUMER_SECRET=xxx
export TWITTER_ACCESS_TOKEN=xxx
export TWITTER_ACCESS_SECRET=xxx
./twitter-action -message "Hello Twitter :)"
```

# Docker
```
# If building locally
docker build -t mathew-fleisch/twitter-action .

# else:
docker run --rm -e TWITTER_CONSUMER_KEY=${TWITTER_CONSUMER_KEY} \
       -e TWITTER_CONSUMER_SECRET=${TWITTER_CONSUMER_SECRET} \
       -e TWITTER_ACCESS_TOKEN=${TWITTER_ACCESS_TOKEN} \
       -e TWITTER_ACCESS_SECRET=${TWITTER_ACCESS_SECRET} \
       mathew-fleisch/twitter-action -message "Hello Twitter :)"
```

# Test
```
./twitter-action -message "Here is the big news: I removed the file functionality" -dry
2021/03/30 16:04:47 Logging in, creating client and updating status.
2021/03/30 16:04:47 Status updated with: Here is the big news: I removed the file functionality
```
