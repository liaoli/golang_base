package myhttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
*@author: 廖理
*@date:2022/11/9
**/

func StartServer() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/register", register)
	http.Handle("/demo", DemoHandler{})
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}

func register(writer http.ResponseWriter, request *http.Request) {

	defer request.Body.Close()
	switch request.Method {
	case "GET":
		data := request.URL.Query()
		fmt.Println(data.Get("name"))
		fmt.Println(data.Get("age"))
		break
	case "POST":

		contentType := request.Header.Get("Content-Type")

		if contentType == "application/json" {
			b, err := ioutil.ReadAll(request.Body)
			if err != nil {
				fmt.Printf("read request.Body failed, err:%v\n", err)
				return
			}
			fmt.Println(string(b))
		} else if contentType == "application/x-www-form-urlencoded" {
			request.ParseForm()
			fmt.Println(request.PostForm) // 打印form数据
			fmt.Println(request.PostForm.Get("name"), request.PostForm.Get("age"))
			// 2. 请求类型是application/json时从r.Body读取数据
		}

	}

	answer := `{"status": 200,"msg":"ok"}`
	writer.Write([]byte(answer))

}

//默认的Server
//ListenAndServe使用指定的监听地址和处理器启动一个HTTP服务端。处理器参数通常是nil，
//这表示采用包变量DefaultServeMux作为处理器。
//Handle和HandleFunc函数可以向DefaultServeMux添加处理器。
func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "hello ,细狗")
}

type DemoHandler struct {
}

func (receiver DemoHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	answer := `{"status": "ok"}`
	writer.Write([]byte(answer))
}
