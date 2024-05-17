package qiniu

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JobBody struct {
	Job string `json:"job"`
	Id  int64  `json:"id"`
	Url string `json:"url"`
	Uid int64  `json:"uid"`
}

type SafeResponse struct {
	Status  string `json:"status"`
	Request struct {
		Data struct {
			Id string `json:"id"`
		} `json:"data"`
	} `json:"request"`
	Result struct {
		Result struct {
			Suggestion string `json:"suggestion"`
		} `json:"result"`
	} `json:"result"`
}

type UploadFile struct {
	Id  int64  `json:"id"`
	Url string `json:"url"`
	Uid int64  `json:"uid"`
}

func GetJobBack(job, secretKey, accessKey string) (status, vid, suggestion string, err error) {
	method := "GET"
	path := "/v3/jobs/video/" + job
	host := "ai.qiniuapi.com"
	requestStr := fmt.Sprintf("%s %s\nHost: %s\n\n", method, path, host)
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write([]byte(requestStr))
	sign := h.Sum(nil)

	// Base64 编码签名
	key := base64.URLEncoding.EncodeToString(sign)
	// 添加 Authorization 头部
	url := "http://" + host + path
	token := "Qiniu " + accessKey + ":" + key
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	var safeInfo SafeResponse
	err = json.Unmarshal(body, &safeInfo)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	return safeInfo.Status, safeInfo.Request.Data.Id, safeInfo.Result.Result.Suggestion, nil
}
