package main

import (
	"net/http"
	"net/url"
	"slipe/freq"
	"slipe/read"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func tianyanInfo(key string) []string {

	client := &http.Client{}

	//生成要访问的url
	url := "https://www.tianyancha.com/search/t400?key=" + key

	content := []string{}

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	// //处理返回结果
	response, _ := client.Do(reqest)

	doc, _ := goquery.NewDocumentFromResponse(response)

	doc.Find(".search_result_type").Each(func(i int, contentSelection *goquery.Selection) {
		title := contentSelection.Find(".row_3 .item-right").Eq(1).Find("span").Text()
		// fmt.Println(title)
		mark := title[2:3]
		// fmt.Println(mark)
		switch mark {
		case "1":
			title = "发明专利"
		case "2":
			title = "实用新型专利"
		case "3":
			title = "外观专利"

		}
		// fmt.Println(title)
		content = append(content, title)

	})

	if doc.Find(".result-footer ul").HasClass("pagination") {
		len := doc.Find(".result-footer li").Length()
		for i := 2; i < len; i++ {

			newURL := "https://www.tianyancha.com/search/t402/p" + strconv.Itoa(i) + "?key=" + key

			time.Sleep(1 * time.Second)

			//提交请求
			reqest, err := http.NewRequest("GET", newURL, nil)

			if err != nil {
				panic(err)
			}

			// //处理返回结果
			response, _ := client.Do(reqest)

			newdoc, _ := goquery.NewDocumentFromResponse(response)

			newdoc.Find(".search_result_type").Each(func(i int, contentSelection *goquery.Selection) {
				title := contentSelection.Find(".row_3 .item-right").Eq(1).Find("span").Text()
				mark := title[2:3]
				switch mark {
				case "1":
					title = "发明专利"
				case "2":
					title = "实用新型专利"
				case "3":
					title = "外观专利"

				}
				// fmt.Println(title)
				content = append(content, title)

			})

		}

	}

	return content
}

func main() {

	countryMap := read.Read()
	// fmt.Println(countryMap)
	for k, v := range countryMap {
		key := url.QueryEscape(v)
		content := tianyanInfo(key)
		// fmt.Println(content)
		countentMap := freq.SplicFreq(content, k)
		read.Write(countentMap)
		time.Sleep(3 * time.Second)
	}

}
