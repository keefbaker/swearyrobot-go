package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func sp() string {
	rand.Seed(time.Now().UTC().UnixNano())
	var spa string
	if rand.Intn(100) < 70 {
		spa = " "
	}
	return spa
}

func tweet(message string) {
	keys := twitconfig()
	config := oauth1.NewConfig(keys[0], keys[1])
	token := oauth1.NewToken(keys[2], keys[3])
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	tweety, resp, err := client.Statuses.Update(message, nil)
	fmt.Println(tweety, resp, err)
}

func getSwear() string {
	time.Sleep(5 * time.Millisecond)
	return strings.ToLower(Swear())
}

func generate() string {
	message := strings.Title(Swear())
	time.Sleep(5 * time.Millisecond)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < rand.Intn(12)+1; i++ {
		message += sp()
		message += getSwear()
	}
	message += Punc()
	message += Hashtag()
	return message
}

func twitconfig() []string {
	var lines []string
	file, err := os.Open("keys.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Hashtag() string {
	var message string
	rand.Seed(time.Now().UTC().UnixNano())
	if rand.Intn(100) > 40 {
		message = message + " #" + getSwear() + getSwear()
	}
	return message
}

func main() {
	tweet(generate())
}
