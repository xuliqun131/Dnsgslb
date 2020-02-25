package main

import (
	"fmt"
	"net/http"

)

func main() {
	//生成client 参数为默认
	client := &http.Client{}

	//生成要访问的url
	url := "http://www.bai123123du.com"

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("err is :", err)
		//处理返回结果
		response, _ := client.Do(reqest)
		//返回的状态码
		status := response.StatusCode
		fmt.Println(status)
	}else {
		//处理返回结果
		response, _ := client.Do(reqest)
		//返回的状态码
		status := response.StatusCode
		fmt.Println(status)
	}
}