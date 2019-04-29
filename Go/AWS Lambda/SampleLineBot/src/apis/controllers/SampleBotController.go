package controllers

import (
	"core/templates"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/linebot/httphandler"
	"log"
	"os"
	"runtime"
	"html/template"
)

func DoSomething(c *gin.Context) {
	log.Printf("Start: %s", funcName())
	t := template.Must(template.New("SampleBot").Parse(tmpl.SmapleBotLifftmpl))
	t.Execute(c.Writer, nil)
	log.Printf("End: %s", funcName())
}

func GetSomething(c *gin.Context) {
	log.Printf("Start: %s", funcName())
	lineId := c.Query("lineId")
	msg := c.Query("msg")

	bot := newClient()
	if bot != nil {
		if _, err := bot.PushMessage(lineId, linebot.NewTextMessage(msg)).Do(); err != nil {
			log.Fatalf("push msg error: %+v", err)
		}
	}
	log.Printf("End: %s", funcName())
}

func newClient() *linebot.Client {
	handler, err := httphandler.New(
		os.Getenv("Secret"),
		os.Getenv("Token"))
	if err != nil {
		log.Fatalf("error %s", err.Error())
		return nil
	}

	bot, err := handler.NewClient()
	if err != nil {
		log.Fatalf("error %s", err.Error())
		return nil
	}
	return bot
}

func funcName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}