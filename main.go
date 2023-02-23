package main

import (
	"checkVerification/model"
	"checkVerification/utils"
	"log"
)

func main() {

	err := utils.ParsConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = model.InitDB()
	if err != nil {
		log.Fatalf("数据库连接错误 %s \n", err.Error())
	}
}
