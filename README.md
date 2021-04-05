# twitter-action

This repository was [forked](https://github.com/xorilog/twitter-action) so I could release just the go-binary to use this tweet functionality in other contexts than just a github-action. I have also removed the file input methods for simplicity, and created a release pipeline to generate and host the go-binary and container.

## Auth
This process to get twitter api keys took a few days of back and forth emails explaining what this app would do.
About the authentication see: https://developer.twitter.com/en/apps
create an account, create an app
@see https://apps.twitter.com/

### retrieve the access tokens
@see https://developer.twitter.com/en/apps

## Build
```bash
go get .
go build .
```

## Usage
```bash
export TWITTER_CONSUMER_KEY=xxx
export TWITTER_CONSUMER_SECRET=xxx
export TWITTER_ACCESS_TOKEN=xxx
export TWITTER_ACCESS_SECRET=xxx
./twitter-action -message "Hello Twitter :)"
```

## Docker
```bash
# If building locally
docker build -t mathew-fleisch/twitter-action .

# else:
docker run --rm -e TWITTER_CONSUMER_KEY=${TWITTER_CONSUMER_KEY} \
       -e TWITTER_CONSUMER_SECRET=${TWITTER_CONSUMER_SECRET} \
       -e TWITTER_ACCESS_TOKEN=${TWITTER_ACCESS_TOKEN} \
       -e TWITTER_ACCESS_SECRET=${TWITTER_ACCESS_SECRET} \
       mathew-fleisch/twitter-action -message "Hello Twitter :)"
```

## Test
```bash
./twitter-action -message "Here is the big news: I removed the file functionality" -dry
2021/03/30 16:04:47 Logging in, creating client and updating status.
2021/03/30 16:04:47 Status updated with: Here is the big news: I removed the file functionality
```

## Releases

There is a [github-action](.github/workflows/tag-release.yaml) to automate the release process. This action will build a docker container that generates a go-binary as a build artifact. The go-binary is uploaded as a release asset, and the docker container is pushed to docker hub, when triggered by pushing a new git tag (that starts with the letter 'v'). The [releases](https://github.com/mathew-fleisch/twitter-action/releases) page contains a go-binary to update a personal twitter status (tweet), by setting a few api keys as environment variables, and executing the go-binary.

**Example Usage**

```bash
twitter_action_tag=v1.0.1
echo "Check environment variables are set..."
expected="TWITTER_CONSUMER_KEY TWITTER_CONSUMER_SECRET TWITTER_ACCESS_TOKEN TWITTER_ACCESS_SECRET GIT_TOKEN"
for expect in $expected; do
  if [[ -z "${!expect}" ]]; then
    echo "Missing secret: $expect"
    exit 1
  fi
done
echo "Pull twitter-action: $twitter_action_tag"
curl -sL -H "Authorization: token $GIT_TOKEN" \
  "https://api.github.com/repos/mathew-fleisch/twitter-action/releases/tags/$twitter_action_tag" \
  | jq -r '.assets[] | select(.name == "twitter-action").browser_download_url' \
  | xargs -I {} curl -sL -H "Authorization: token $GIT_TOKEN" -H "Accept:application/octet-stream" -O {}
chmod +x twitter-action
./twitter-action -message "This is neat https://github.com/mathew-fleisch/blog"
```