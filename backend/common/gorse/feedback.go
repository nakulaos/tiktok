package gorse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func Feedback(recommendUrl, feedType string, vid int, uid int) (err error) {
	baseUrl := recommendUrl + "/api/feedback"
	req := []map[string]interface{}{
		{
			"FeedbackType": feedType,
			"ItemId":       strconv.Itoa(vid),
			"Timestamp":    time.Now().Format("2006-01-02T15:04:05.000Z"),
			"UserId":       strconv.Itoa(uid),
		},
	}
	jsonData, err := json.Marshal(req)
	if err != nil {
		fmt.Println("JSON编码失败:", err)
		return
	}
	println(string(jsonData))
	resp, err := http.Post(baseUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("POST请求发送失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容失败:", err)
		return
	}
	fmt.Println("上传反馈成功", string(s))

	return err
}
