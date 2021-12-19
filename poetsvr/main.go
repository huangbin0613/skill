package main

import (
	"fmt"
	"net/http"
	"poetsvr/data"
	"poetsvr/poet"
)

func main() {
	fmt.Println("hello world")

	http.HandleFunc("/intent", poet.Intent)

	//data.InsertPoet(&data.Poet{
	//	Title: "AAA",
	//	Author: "B",
	//})
	//
	//data.InsertPoet(&data.Poet{
	//	Title: "AAA124",
	//	Author: "B",
	//})
	//data.UpdatePoet(&data.Poet{
	//	Id: 1,
	//	Title: "A",
	//	Author: "B",
	//	Content: "AAAA",
	//})

	ret, count, err := data.GetPoetListByContent("", 1, 2)
	fmt.Println(poet.ToJson(ret), count, err)

	// 开启http服务
	err = http.ListenAndServe("127.0.0.1:12345", nil)
	fmt.Println(err)
}
