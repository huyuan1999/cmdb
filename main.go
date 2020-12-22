package main

import (
	"cmdb/client/Linux"
	"encoding/json"
	"log"
)

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func SpecialJson(v interface{}, err error) string {
	FatalError(err)
	data, err := json.Marshal(v)
	FatalError(err)
	return string(data)
}

func main() {
	log.Println("处理器信息: ", SpecialJson(Linux.NewCPU()))
	log.Println("硬盘信息: ", SpecialJson(Linux.NewDisk()))
	log.Println("内存信息: ", SpecialJson(Linux.NewMemory()))
	log.Println("主板信息: ", SpecialJson(Linux.NewMainBoard()))
	log.Println("网卡信息: ", SpecialJson(Linux.NewNIC()))
	log.Println("系统信息: ", SpecialJson(Linux.NewSystem()))
}
