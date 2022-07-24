package main

import (
	"fmt"
	"gitee.com/phper95/pkg/docker"
	"github.com/valyala/fasthttp"
	"net/http"
	"sync"
	"time"
)

const (
	StartTimeoutSecond = 180
	User               = "test"
	Pass               = "unit-test"
)

func main() {
	InitMiddleware()
}

func checkESSever() bool {
	url := "http://localhost:9200"
	for ticker := 1; ticker < StartTimeoutSecond; ticker++ {
		httpCode, _, _ := fasthttp.Get(nil, url)
		if httpCode == http.StatusOK {
			return true
		}
		time.Sleep(time.Second)
		fmt.Println("check ES", ticker)
	}
	return false
}

func startES(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	containerOption := docker.ContainerOptions{
		Name:      "elastic-unittest",
		ImageName: "phper95/es8.1.0",
		Options: map[string]string{
			"xpack.security.enabled": "false",
			"discovery.seed_hosts":   "127.0.0.1:9300",
			"discovery.type":         "single-node",
		},
		MountPath:  "/usr/share/elasticsearch/data",
		PortExpose: "9200",
	}
	ESDocker := &docker.Docker{}
	if !ESDocker.IsInstalled() {
		panic("docker has`t install")
	}
	err := ESDocker.RemoveIfExists()
	if err != nil {
		panic(err)
	}
	res, err := ESDocker.Start(containerOption)
	if err != nil {
		fmt.Println(res)
		panic(err)
	}
	fmt.Println("docker", containerOption.ImageName, "has started")
	if checkESSever() {
		fmt.Println("es sever has started")
	} else {
		fmt.Println("es sever started timeout")
		ESDocker.RemoveIfExists()
	}
}

func StartMysql(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	mysqlOptions := map[string]string{
		"MYSQL_ROOT_PASSWORD": Pass,
		"MYSQL_USER":          User,
		"MYSQL_PASSWORD":      Pass,
		"MYSQL_DATABASE":      "shop",
	}

	containerOption := docker.ContainerOptions{
		Name:       "mysql-unittest",
		ImageName:  "mysql:5.7",
		Options:    mysqlOptions,
		MountPath:  "/var/lib/mysql",
		PortExpose: "3306",
	}
	mysqlDocker := docker.Docker{}
	if !mysqlDocker.IsInstalled() {
		panic("docker has`t install")
	}
	err := mysqlDocker.RemoveIfExists()
	if err != nil {
		panic(err)
	}
	res, err := mysqlDocker.Start(containerOption)
	if err != nil {
		fmt.Println(res)
		panic(err)
	}
	mysqlDocker.WaitForStartOrKill(StartTimeoutSecond)
	fmt.Println("mysql has started")
}

func InitMiddleware() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go startES(&wg)
	go StartMysql(&wg)
	wg.Wait()

}
