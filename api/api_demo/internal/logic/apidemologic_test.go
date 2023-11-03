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
	var result map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}

	fmt.Println(result)
}
