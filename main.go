package main

import (
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func getResults() [][]string {
	url := "https://fofa.info/api/v1/search/all?email=" + fofaEmail + "&key=" + fofaKey + "&qbase64=" + qbase64s[i] + "&size=" + fmt.Sprint(size) + "&fields=" + fields + "&page=" + fmt.Sprint(page) + ""
	// 发送 GET 请求
	log.Println(url)
	response, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil
	}
	var fakeWebInfo FakeWebInfo
	err = json.Unmarshal([]byte(string(body)), &fakeWebInfo)
	if err != nil {
		return nil
	}
	if len(fakeWebInfo.Results) > 0 {
		results := fakeWebInfo.Results
		if page <= fakeWebInfo.Size/size+1 {
			page = page + 1
			results = append(results, getResults()...)
		}
		return results
	} else {
		fmt.Println(string(body))
		page = 1
		fmt.Println()
		// if len(fakeWebInfo.Results) > 0 {
		results := fakeWebInfo.Results
		if i < len(qbase64s)-1 { //3210 /1000 =3
			i = i + 1
			results = append(results, getResults()...)
		}
		return results
	}
}

func getFakeWebResults() []FakeWebResults {
	var fakeWebResults []FakeWebResults
	for _, info := range getResults() {
		hashValue := md5.Sum([]byte(info[0] + info[1] + info[2] + info[13]))
		hashString := hex.EncodeToString(hashValue[:])
		fakeWebResult := FakeWebResults{
			IP:              info[0],
			Port:            info[1],
			Domain:          info[2],
			Country:         info[3],
			Title:           info[4],
			Protocol:        info[5],
			Country_name:    info[6],
			Region:          info[7],
			City:            info[8],
			Longitude:       info[9],
			Latitude:        info[10],
			As_number:       info[11],
			As_organization: info[12],
			Host:            info[13],
			Os:              info[14],
			Server:          info[15],
			Icp:             info[16],
			Jarm:            info[17],
			Header:          info[18],
			Banner:          info[19],
			Cert:            info[20],
			Base_protocol:   info[21],
			Cname:           info[22],
			HashValue:       hashString,
		}
		fakeWebResults = append(fakeWebResults, fakeWebResult)
	}
	return fakeWebResults
}

func checkURL(url string) string {
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"
	// 创建自定义的 HTTP 客户端，忽略证书告警
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	// 创建带超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return ""
	}
	req.Header.Add("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ""
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	bodyString := string(bodyBytes)
	if strings.Contains(bodyString, "【国家反诈中心、工信部反诈中心、中国电信联合提醒】") {
		return "【已被反诈中心屏蔽】"
	}
	return "【仍可访问】"
}

func writeFile(stringSlice []string) {
	file, err := os.Create(time.Now().Format("2006-01-02") + ".txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	// 将每个字符串写入文件中，每个字符串一行
	for _, str := range stringSlice {
		_, err := fmt.Fprintln(file, str)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

}

// 定时任务
func scheduleTask(task func(), hour, minute int) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
	if now.After(today) {
		today = today.Add(24 * time.Hour) // 将时间设置为明天
	}
	duration := today.Sub(now)
	time.AfterFunc(duration, task)
}

func main() {
	for {
		scheduleTask(func() {
			log.Println("任务开始！")
			page = 1
			i = 0
			results := getFakeWebResults()
			// 记录
			writeSql(results)
			urls := selectSqlAliveUrls()
			writeFile(urls)
			// 日报
			dailyReportRobot(false, results, urls)
			downloadFileRobot(uploadFileRobot(time.Now().Format("2006-01-02") + ".txt"))
		}, 12, 29)
		time.Sleep(2 * time.Minute)
	}

}

