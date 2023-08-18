package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func vxRobot(mentionAll bool, IP string, Port string, Host string, Domain string, Country string) { //获取新流量通知到群
	mentionAllStr := ""
	if mentionAll {
		mentionAllStr = "@all"
	}
	dataJsonStr := fmt.Sprintf(`{"msgtype": "markdown", "markdown": {"content": "<font color=\"red\">[网站监控] 新增站点: </font>\nHost : %s\nDomain : %s\nIP : %s\nPort :%s\nCountry : %s\nTime : %s"", "mentioned_list": [%s]}}`, Host, Domain, IP, Port, Country, timenow(), mentionAllStr)
	// fmt.Println(dataJsonStr)

	resp, err := http.Post(
		"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+vxRobotKey,
		"application/json",
		bytes.NewBuffer([]byte(dataJsonStr)))
	if err != nil {
		// fmt.Println("weworkAlarm request error")
		return
	}
	defer resp.Body.Close()
}
func timenow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func dailyReportRobot(mentionAll bool, fakeWebResults []FakeWebResults, aliveUrl []string) {
	urlslen := len(aliveUrl)
	resultslen := len(fakeWebResults)
	mentionAllStr := ""

	if mentionAll {
		mentionAllStr = "@all"
	}
	dataJsonStr := fmt.Sprintf(`{"msgtype": "markdown", "markdown": {"content": "<font color=\"red\">[网站监控日报] </font>\n今日共计 ` + fmt.Sprint(resultslen) + ` 条匹配结果，存活站点共计 ` + fmt.Sprint(urlslen) + ` 个。\n时间 : ` + timenow() + `"", "mentioned_list": [` + mentionAllStr + `]}}`)
	fmt.Println(dataJsonStr)
	resp, err := http.Post(
		"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+vxRobotKey,
		"application/json",
		bytes.NewBuffer([]byte(dataJsonStr)))
	if err != nil {
		// fmt.Println("weworkAlarm request error")
		return
	}
	defer resp.Body.Close()
	// webAlive()
}

func downloadFileRobot(mediaid string) {
	webhookURL := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + vxRobotKey
	// 构建消息
	message := Message{
		MsgType: "file",
		File: struct {
			MediaID string `json:"media_id"`
		}{
			MediaID: mediaid,
		},
	}
	// 将消息转换为 JSON
	messageJSON, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 发送 POST 请求到企业微信机器人 webhook URL
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(messageJSON))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	fmt.Println("File uploaded successfully.")
}

func uploadFileRobot(filePath string) string {
	// 打开要上传的文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()
	// 创建一个缓冲区来构建请求体
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	// 创建一个表单字段并将文件内容写入
	fileField, err := writer.CreateFormFile("file", time.Now().Format("2006-01-02")+"网站.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	_, err = io.Copy(fileField, file)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 结束写入请求体并设置Content-Type
	writer.Close()
	// 创建一个POST请求
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media?key=" + vxRobotKey + "&type=file&debug=1"
	request, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 设置请求头，指定Content-Type为multipart/form-data
	request.Header.Set("Content-Type", writer.FormDataContentType())
	// 发送请求并获取响应
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer response.Body.Close()
	bodyBytes, _ := io.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	var mediaInfo MediaInfo
	err = json.Unmarshal([]byte(bodyString), &mediaInfo)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	fmt.Println(mediaInfo.MediaID)
	return mediaInfo.MediaID
}
