package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/kardianos/service"
)

var Main = &gcmd.Command{
	Name:  "chogori-agent",
	Brief: "chogori-agent",
	Arguments: []gcmd.Argument{
		{Name: "kubeconfig", Brief: "kubernetes config file, parse containerId from v1.Pod"},
		{Name: "controller", Brief: "chogori-controller for alloc/free GPU resources"},
		{Name: "chogori-resource-prefix", Brief: "chogori device resource prefix"},
		{Name: "chogori-resource-compat", Brief: "chogori device resource compat"},
		{Name: "chogori-root", Brief: "chogori root directory"},
		{Name: "config", Brief: "chogori configuration file"},
		{Name: "version", Short: "v", Orphan: true, Brief: "show version"},
		{Name: "sock", Brief: "path of chogori.sock (uds)"},
		{Name: "listen", Brief: "listen address of chogori-agent (http)"},
		{Name: "enable-dynamic-free", Brief: "dynamic free of device resources"},
		{Name: "free-ticker", Brief: "check the interval for free app (second)"},
		{Name: "heartbeat-timeout", Brief: "heartbeat timeout (second)"},
	},
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		glog := g.Log("logger_windows")
		if cfgFile := parser.GetOpt("kubeconfig"); cfgFile != nil {
			logger.Infof("set config file: %s", cfgFile.String())
			glog.Infof(ctx, "set config file: %s", cfgFile.String())
		}
		glog.Infof(ctx, "121212")
		return nil
	},
}

var logger service.Logger

type program struct {
	serviceArgs []string
}

func (p *program) Start(s service.Service) error {
	// 只有服务真正启动时才会调用
	go p.run()
	return nil
}

func (p *program) run() {
	Main.Run(gctx.New())
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "AAAAAAAAAAAAA",
		DisplayName: "AAAAAAAAAAAAA",
		Description: "带参数的服务示例",
	}

	prg := &program{}

	if len(os.Args) > 1 {
		cmd := os.Args[1]
		switch cmd {
		case "install":
			// 获取要传递给服务的参数（跳过install命令本身）
			serviceArgs := os.Args[2:]

			// 将参数转换为服务启动参数
			var args []string
			for _, arg := range serviceArgs {
				// 处理可能包含空格或特殊字符的参数
				if strings.Contains(arg, " ") {
					arg = `"` + arg + `"`
				}
				args = append(args, arg)
			}

			// 设置服务启动参数
			svcConfig.Arguments = args

			s, err := service.New(prg, svcConfig)
			if err != nil {
				log.Fatal(err)
			}
			if err := s.Install(); err != nil {
				log.Fatal(err)
			}
			log.Printf("服务安装成功，参数: %v", args)
			return
		case "uninstall":
			s, err := service.New(prg, svcConfig)
			if err != nil {
				log.Fatal(err)
			}
			if err := s.Uninstall(); err != nil {
				log.Fatal(err)
			}
			log.Println("服务卸载成功")
			return
		case "start":
			s, err := service.New(prg, svcConfig)
			if err != nil {
				log.Fatal(err)
			}
			if err := s.Start(); err != nil {
				log.Fatal(err)
			}
			log.Println("服务启动成功")
			return
		case "stop":
			s, err := service.New(prg, svcConfig)
			if err != nil {
				log.Fatal(err)
			}
			if err := s.Stop(); err != nil {
				log.Fatal(err)
			}
			log.Println("服务停止成功")
			return
		}
	}

	// 参数
	prg.serviceArgs = os.Args[1:]
	//logger.Infof("os.Args完整内容: %v", os.Args)

	// 创建服务
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	// 运行服务
	if err := s.Run(); err != nil {
		logger.Error(err)
	}
}
