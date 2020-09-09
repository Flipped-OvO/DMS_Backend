package main

import (
	"backend/pkg/lixinger"
	"backend/util"
	"backend/util/router"
	"backend/util/setting"
	"fmt"
	"log"
	"net/http"
	"time"
)

// TODO 数据库加密
func main() {
	startServer()
}

func startServer() {
	_, err := time.LoadLocation("Local")
	if err != nil {
		log.Println(err)
	}
	util.LogInitialization()
	go func() {
		lixinger.FetchIncomStatement()
		//lixinger.FetchBalanceSheet()
		//lixinger.FetchCashFlow()
	}()

	r := router.Initialization()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.Config.Runtime.Port),
		Handler:        r,
		TLSConfig:      nil,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = s.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
