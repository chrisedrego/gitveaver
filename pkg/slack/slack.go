package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func constructSlackMessage(prList, listSlackUser []string) (slackMessage string) {
	// Construct Slack Message
	var SlackUserId string
	SlackUserId = strings.Join(listSlackUser, " ")
	fmt.Println(SlackUserId)
	SlackMessageText := "\nWe need your attention, could you quickly reviews these Pull Request assigned to you."
	var prTextPayload string
	for _, val := range prList {
		prTextPayload = val + "\n"
		fmt.Println(prTextPayload)
	}
	slackMessage = "Hey! :wave:" + SlackUserId + SlackMessageText + "\n" + prTextPayload
	fmt.Println("Slack Message:", slackMessage)
	return slackMessage
}

func SlackNotifier(channel, username, icon_emoji string, prList, slackUserList []string) {
	// Send Slack Notification
	slackUrl := os.Getenv("SLACK_URI")
	fmt.Println("Slack-URI Details", slackUrl)
	textPayload := constructSlackMessage(prList, slackUserList)
	payload := map[string]interface{}{"channel": channel, "username": username, "icon_emoji": icon_emoji, "text": textPayload}
	// fmt.Println(payload)
	postBody, _ := json.Marshal(payload)
	// fmt.Println(postBody)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(slackUrl, "application/json", responseBody)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(json.NewDecoder(resp.Body))
	fmt.Println(resp.StatusCode)
}
