package main

import (
	"checkVerification/model"
	"checkVerification/utils"
	"checkVerification/web"
	"log"
)

func main() {

	err := utils.ParsConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	if utils.GetConfig().Mysql != "" {
		err = model.InitDB()
		if err != nil {
			log.Fatalf("Mysql连接错误 %s \n", err.Error())
		}
	} else {
		err = model.InitRedis()
		if err != nil {
			log.Fatalf("Redis连接错误 %s \n", err.Error())
		}
	}

	err = web.InitHttp()
	if err != nil {
		log.Fatalf("http server err %s \n", err.Error())
	}
}
