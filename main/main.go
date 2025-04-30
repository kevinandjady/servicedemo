package main

import (
	"log"
	"os"
	"time"

	"github.com/kardianos/service"
)

var logger service.Logger

// MyService 实现了 service.Service 接口
type MyService struct{}

func (m *MyService) Start(s service.Service) error {
	go m.run()
	return nil
}

func (m *MyService) run() {
	// 在这里编写你的服务逻辑
	for {
		time.Sleep(1 * time.Second)
	}
}

func (m *MyService) Stop(s service.Service) error {
	// 停止服务的逻辑
	return nil
}

func main() {

	// 服务的名称、显示名称和描述
	svcConfig := &service.Config{
		Name:        "ChogoriAgentService",
		DisplayName: "Chogori Agent Service",
		Description: "This is a sample service by golang.",
	}

	prg := &MyService{}

	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	// 通过以下代码来控制服务的启动和停止
	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
