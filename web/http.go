package web

import (
	"checkVerification/middleware"
	"checkVerification/utils"
	"fmt"
	"net/http"
)

var website Website

func InitHttp() (err error) {
	// website
	http.Handle("/getVerify", middleware.HttpHandler(http.HandlerFunc(website.GetVerify)))

	fmt.Printf("Server is running port: %s \n", utils.GetConfig().Port)

	err = http.ListenAndServe(":"+utils.GetConfig().Port, nil)

	if err != nil {
		return err
	}

	return nil
}
