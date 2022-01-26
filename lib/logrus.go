package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/david8998/lib/config"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var Logger *logrus.Logger

func NewLog(file string) (log *logrus.Logger) {
	log = logrus.New()
	file = strings.TrimSpace(file)
	if file == "" {
		file = "default"
	}
	log.SetOutput(&lumberjack.Logger{
		Filename: "/home/work/logs/data_sdk/" + file + ".log",
		MaxSize:  1024, // 1g
		MaxAge:   3,    // 7days
		Compress: true,
	})
	//log.SetReportCaller(true)
	tg := &TgNotify{token: "5007642930:AAEcmks_aoCudACQWM5H0I_i9pRhm3BsOms", chatId: "-509509547"}
	log.AddHook(tg)
	return
}

type TgNotify struct {
	token  string
	chatId string
}

func (t *TgNotify) Levels() []logrus.Level {
	return []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel}
}
func (t *TgNotify) Fire(entry *logrus.Entry) error {
	msg := fmt.Sprintf("\n==================================\n "+
		"env :%s\n ts: %s \n msg: %s\n"+
		"==================================\n",
		config.GetEnv(), time.Now().String(), entry.Message)
	go func() {
		body := struct {
			ChatId string `json:"chat_id"`
			Text   string `json:"text"`
		}{
			ChatId: t.chatId,
			Text:   msg,
		}
		url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.token)
		Post(url, body, "application/json")
	}()
	return nil
}

func Post(url string, data interface{}, contentType string) string {
	// 超时时间：5秒
	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
