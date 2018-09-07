package main

import (
	"net/http"
	"net/url"
	"slipe/freq"
	"slipe/read"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func tianyanID(key string) string {

	client := &http.Client{}

	//生成要访问的url
	url := "https://www.tianyancha.com/search?key=" + key

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	// //处理返回结果
	response, _ := client.Do(reqest)

	doc, _ := goquery.NewDocumentFromResponse(response)

	str, _ := doc.Find(".search-result-single").Attr("data-id")

	return str
}

func tianyanInfo(id string) []string {

	client := &http.Client{}

	//生成要访问的url
	url := "https://www.tianyancha.com/company/" + id
	// url := "https://www.tianyancha.com/company/1265306443"

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	// //处理返回结果
	response, _ := client.Do(reqest)

	doc, _ := goquery.NewDocumentFromResponse(response)
	content := []string{}

	doc.Find("#_container_patent tbody > tr").Each(func(i int, contentSelection *goquery.Selection) {
		title := contentSelection.Find("td:nth-child(6) > span").Text()
		content = append(content, title)

	})

	return content
}

func main() {

	countryMap := read.Read()
	// fmt.Println(countryMap)

	for k, v := range countryMap {

		key := url.QueryEscape(v)
		id := tianyanID(key)
		time.Sleep(1 * time.Second)
		content := tianyanInfo(id)
		countentMap := freq.SplicFreq(content, k)
		read.Write(countentMap)
		time.Sleep(1 * time.Second)
	}

}
