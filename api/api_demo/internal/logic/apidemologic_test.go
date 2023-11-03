package logic

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpc"
	"io/ioutil"
	"net/http"
	"testing"
)

type Request struct {
	Type string `form:"type"`
	//ID     int    `form:"id"`
	Header string `header:"X-Header"`
}

var domain = flag.String("domain", "https://api.vvhan.com/api/hotlist", "the domain to request")

func TestApi_demoLogic_Api_demo(t *testing.T) {
	flag.Parse()

	req := Request{
		Type: "zhihuHot",
	}
	resp, err := httpc.Do(context.Background(), http.MethodGet, *domain, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	//得到返回结果
	var result Gen
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}

	fmt.Println(result)
}

type Gen struct {
	UpdateTime string `json:"update_time" gorm:"column:update_time"`
	Data       []struct {
		MobilUrl string `json:"mobilUrl" gorm:"column:mobilUrl"`
		Index    int    `json:"index" gorm:"column:index"`
		Pic      string `json:"pic" gorm:"column:pic"`
		Title    string `json:"title" gorm:"column:title"`
		Hot      string `json:"hot" gorm:"column:hot"`
		Url      string `json:"url" gorm:"column:url"`
		Desc     string `json:"desc" gorm:"column:desc"`
	} `json:"data" gorm:"column:data"`
	Success  bool   `json:"success" gorm:"column:success"`
	Subtitle string `json:"subtitle" gorm:"column:subtitle"`
	Title    string `json:"title" gorm:"column:title"`
}
