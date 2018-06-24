/*
 * Copyright (c) 2018
 * time:   6/24/18 3:22 PM
 * author: linhuanchao
 * e-mail: 873085747@qq.com
 */

package main
import (
	"github.com/gin-gonic/gin"
	"SensitiveWords/controller"
	"sync"
	"flag"
	"SensitiveWords/config"
	"os"
	"path/filepath"
	"fmt"
	"io/ioutil"
	"os/exec"
)

var configure *config.Config
var appPath string
var pidPath string
func init()  {
	appPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	configure = config.GetConfig()
	pidPath = appPath + configure.PidFilePath
	gin.SetMode(gin.ReleaseMode)
}

var wg sync.WaitGroup
func main(){
	var start bool
	var stop bool
	var daemon bool
	var restart bool
	flag.BoolVar(&start,"start",false,"up your http server, just like this: -start or -start=true|false.")
	flag.BoolVar(&stop,"stop",false,"down your http server, just like this: -stop or -stop=true|false.")
	flag.BoolVar(&daemon,"d",false,"daemon, just like this: -start -d or -d=true|false.")
	flag.BoolVar(&restart,"restart",false,"restart your http server, just like this: -restart or -restart=true|false.")
	flag.Parse()

	if start {
		if daemon {
			cmd := exec.Command("./SensitiveWords", "-start")
			cmd.Start()
			os.Exit(0)
		}
		wg.Add(1)
		fmt.Println("http server start.")
		Start()
		wg.Wait()
	}

	if stop {
		Stop()
	}

	if restart {
		Restart()
	}
}

func Start()  {
	defer wg.Done()

	ioutil.WriteFile(pidPath, []byte(fmt.Sprintf("%d", os.Getpid())), 0666)//记录pid
	router := gin.Default()
	router.GET("/all", controller.All)
	router.GET("/check", controller.Check)
	router.Run(":"+configure.Port)
}

func Stop()  {
	pid, _ := ioutil.ReadFile(pidPath)
	cmd := exec.Command("kill","-9", string(pid))
	cmd.Start()
	ioutil.WriteFile(pidPath, nil, 0666)//清除pid
	fmt.Println("bye~")
}

func Restart()  {
	fmt.Println("restarting...")
	pid, _ := ioutil.ReadFile(pidPath)
	stop := exec.Command("kill","-9", string(pid))
	stop.Start()
	start := exec.Command("./SensitiveWords", "-start", "-d")
	start.Start()
}