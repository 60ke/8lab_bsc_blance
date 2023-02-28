package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

/*
通知钉钉bsc启动,此通知只有在docker启动时会
触发,故可以用来表示bsc节点重启(启动)信息
*/

type DingText struct {
	Content string `json:"content"`
}

type DingMsg struct {
	MsgType string   `json:"msgtype"`
	Text    DingText `json:"text"`
}

func genMsg(content string) string {
	var msg DingMsg
	msg.MsgType = "text"
	msg.Text.Content = content
	data, err := json.Marshal(msg)
	if err != nil {
		Logger.Sugar().Error(err.Error())
	}
	return string(data)
}

// 发送钉钉通知
func sendMsg(content string) {
	Logger.Sugar().Info("send ", "Msg:", content)
	url := viper.GetString("ding-url") + viper.GetString("ding-token")
	ip := GetOutboundIP().String()
	msgData := fmt.Sprintf("BSCBALANCE-%s:%s %30s", viper.GetString("ding-prefix"), ip, content)
	msg := genMsg(msgData)

	payload := strings.NewReader(msg)
	post(url, payload)
}

func post(url string, payload *strings.Reader) ([]byte, error) {

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		Logger.Sugar().Error(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		Logger.Sugar().Error(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		Logger.Sugar().Error(err)
		return nil, err
	}
	Logger.Sugar().Info(string(body))
	return body, nil
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
