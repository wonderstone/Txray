package cmd

import (
	"Txray/core/manage"
	"Txray/xray"
	"github.com/abiosoft/ishell"
	"strconv"
)

func InitServiceShell(shell *ishell.Shell) {
	// 启动或重启服务
	shell.AddCmd(&ishell.Cmd{
		Name: "run",
		Help: "启动或重启服务",
		Func: func(c *ishell.Context) {
			if len(c.Args) == 1 {
				xray.Start(c.Args[0])
			} else {
				xray.Start(strconv.Itoa(manage.Manager.SelectedIndex()))
			}

		},
	})
	// 停止服务
	shell.AddCmd(&ishell.Cmd{
		Name: "stop",
		Help: "停止服务",
		Func: func(c *ishell.Context) {
			xray.Stop()
		},
	})
		// 设置proxy
		shell.AddCmd(&ishell.Cmd{
			Name: "setproxy",
			Help: "设置代理",
			Func: func(c *ishell.Context) {
				xray.SetProxy()
			},
		})
	
		// 反置proxy
		shell.AddCmd(&ishell.Cmd{
			Name: "unsetproxy",
			Help: "取消代理",
			Func: func(c *ishell.Context) {
				xray.UnsetProxy()
			},
		})
	
	
}
