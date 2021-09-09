package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

func main() {
	cli := http.Client{}
	now := time.Now().String()
	urlmap := url.Values{}
	urlmap.Add("Content-Type", "application/json")
	urlmap.Add("X-Request-Id", "123456")
	urlmap.Add("X-Request-Date", now)
	parms := ioutil.NopCloser(strings.NewReader(urlmap.Encode()))
	req, err := http.NewRequest("GET", "http://127.0.0.1:56666", parms)
	if err != nil {
		log.Fatal(err)
	}

	headerPart := "application/json|123456|" + now

	queryPart := ""
	Headerkeys := []string{"Content-Type", "X-Request-Id", "X-Request-Date"}
	sort.Strings(Headerkeys)
	for _, key := range Headerkeys {
		if len(urlmap[key]) > 0 {
			if queryPart != ""{
				queryPart += "&"
			}
			queryPart += key +"="+strings.Join(urlmap[key], ",")
		}
	}

	buf, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	bodyPart := strings.ToLower(NewMd5ByByte(buf))

	//生成加密
	Authorization := "HAMC-SHA1"
	headerrequest := EncondeSignature(headerPart, queryPart, bodyPart)
	signature := EncondeSignature(headerrequest, "root2", "root1")
	Authorization += " " + "root1:" + signature

	req.Header.Set("Authorization", Authorization)
	cli.Do(req)
}

func NewMd5ByByte(data []byte) string {
	m := md5.New()
	m.Write(data)
	sss := m.Sum(nil)
	return fmt.Sprintf("%x", sss)
}

func EncondeHeaderRequest(headerpart, querypart, bodypart string) string {
	part := strings.ToLower(headerpart + querypart + bodypart)
	s := sha1.New()
	s.Write([]byte(part))
	return hex.EncodeToString(s.Sum(nil))
}

func EncondeSignature(headerRequest, appSecret, AccessKey string) string {
	h := hmac.New(sha1.New, []byte(AccessKey))
	h.Write([]byte(headerRequest + appSecret))
	return hex.EncodeToString(h.Sum(nil))
}
