package main

import (
	"fmt"
	"github.com/yutify/gengo-api/config"
	"github.com/yutify/gengo-api/controller/handler"
	"github.com/yutify/gengo-api/controller/router"
	"github.com/yutify/gengo-api/infrastructure/datastore/mysql"
	"github.com/yutify/gengo-api/usecase/service"
	"net/http"
)

func main() {
	conf := &config.Config{}
	if err := config.New(conf, "./.config/config.toml"); err != nil {
		panic(err)
	}
	cm, err := mysql.New(&conf.DBMaster)
	if err != nil {
		panic(err)
	}
	defer cm.Close()

	cs, err := mysql.New(&conf.DBSlave)
	if err != nil {
		panic(err)
	}
	defer cs.Close()

	// create mysql client
	cr := mysql.NewChargeRepository(cm, cs)
	ur := mysql.NewUserRepository(cm, cs)
	ua := service.NewUserService(ur)
	ca := service.NewChargeService(cr, ur)
	handler := handler.NewHandler(ua, ca)
	r := router.Route(handler)
	s := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	defer s.Close()
	fmt.Println("=== start server ===")
	if err := s.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("Failed ListenAndServe. err: %v", err))
	}
}
