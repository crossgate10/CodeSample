package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// aws lambda mode
	lambda.Start(Handler)

	// local debug mode
	//NotifySlack()
}

// Handler handle done event after calling NotifySlack in aws lambda.
func Handler(ctx context.Context) {
	NotifySlack()
	ctx.Done()
}

// NotifySlack do thing as it's name.
func NotifySlack() {
	channel := os.Getenv("SLACK_CHANNEL")
	hookURL := os.Getenv("HOOK_URL")
	message := os.Getenv("MESSAGE")
	log.Printf("[Getenv] SLACK_CHANNEL=%s, HOOK_URL=%s, MESSAGE=%s", channel, hookURL, message)
	jsonMessage := map[string]interface{}{
		"channel": channel,
		"text":    message,
	}

	// 取時間 & 時區設定
	// n := time.Now()
	// location, _ := time.LoadLocation("") // UTC
	// n.In(location).Format("2006/01/02 15:04:05")

	bytesRepresentation, err := json.Marshal(jsonMessage)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(
		hookURL,
		"application/json",
		bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	// 處理 Response
	// var result map[string]interface{}
	// json.NewDecoder(resp.Body).Decode(&result)
	// log.Println(result)
	// log.Println(result["data"])
	log.Println("Send to Slack api response StatusCode: " + strconv.Itoa(resp.StatusCode))
}
