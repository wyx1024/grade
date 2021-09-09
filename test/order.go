package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	orderUrl := fmt.Sprintf("%s/page/index/abnormal_order_detail?order_id=%s&commit_id=%s",
		"http://servicev2.1plustore.com:56666/", "163090998170228660", "3718")
	reqURL := "http://order.gz-cube.com/server/order.php"
	values := url.Values{}
	values.Set("deviceid", "866262048840564")
	values.Set("orderid", "163090998170228660")
	values.Set("msg", "正常订单")
	values.Set("url", url.QueryEscape(orderUrl))
	cli := http.Client{Timeout: time.Second * 30}
	req, _ := http.NewRequest("POST", reqURL, strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := cli.Do(req)
	if err != nil {
		fmt.Println("Order Exception Notify WeChat,OrderID:%s, MID, %s Err %s", "163063645370285934", "122", err.Error())
	}

	reader := simplifiedchinese.GB18030.NewDecoder().Reader(resp.Body)
	all, _ := ioutil.ReadAll(reader)
	fmt.Println(string(all))

}
