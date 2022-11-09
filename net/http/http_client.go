package myhttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/**
*@author: 廖理
*@date:2022/11/9
**/

func HttpRequest() {
	resp, err := http.Get("http://127.0.0.1:9090")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}

	fmt.Println(string(body))
}

func GetRequestWithParam() {
	apiUri := "http://127.0.0.1:9090/register"

	data := url.Values{}

	data.Set("name", "细狗")
	data.Set("age", "44")

	u, err := url.ParseRequestURI(apiUri)

	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}

	u.RawQuery = data.Encode()

	fmt.Println(u.String())

	resp, err := http.Get(u.String())

	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}

	fmt.Println(string(body))
}

func PostRequestWithParam() {
	apiUri := "http://127.0.0.1:9090/register"

	contentType := "application/x-www-form-urlencoded"
	data := "name=小王子&age=18"
	// json
	//data:=`{"name":"细狗","age":"44"}`
	//contentType := "application/json"

	resp, err := http.Post(apiUri, contentType, strings.NewReader(data))

	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}

	fmt.Println(string(body))
}
