package main

import (
	"github.com/ahmedmansourxyz/EpicGamesWebScraper/scraper"
	"github.com/ahmedmansourxyz/EpicGamesWebScraper/slack"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main(){
	intervals := map[string]time.Duration{
		"nano":  time.Nanosecond,
		"micro": time.Microsecond,
		"milli": time.Millisecond,
		"sec":   time.Second,
		"min":   time.Minute,
		"h":     time.Hour,
		"d":     time.Hour * 24,
		"mon":   time.Hour * 24 * 30,
		"year":  time.Hour * 24 * 365,
	}

	var webhookUrl string

	fmt.Print("Enter your Webhook URL: ")
	fmt.Scanln(&webhookUrl)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter how often do you want to be notified (example: 15 sec, 1 minute..): ")
	scanner.Scan()
	constraints := strings.Split(scanner.Text(), " ")
	var interval time.Duration

	if len(constraints) == 2 {
		String, Value := constraints[0], constraints[1]
		count, err := strconv.ParseInt(String, 10, 64)
		if err != nil {
			panic(err.Error())
		}
		interval = time.Duration(count) * intervals[Value]
	} else {
		panic("Invalid input")
	}

	timer := time.NewTimer(interval)
	fmt.Println("Started Timer:")

	for {
		select {
		case <-timer.C:
			slack.SendSlackMessage(webhookUrl, scraper.Games())
			fmt.Println("Message sent.")
			timer = time.NewTimer(interval)
		}
	}
}
