package main

import (
	"fmt"
	"gitee.com/phper95/pkg/shutdown"
	"github.com/valyala/fasthttp"
	"net/http"
	"shop-search-api/internal/pkg/docker"
	"time"
)

const (
	StartTimeoutSecond = 180
	ESUser             = "elastic"
	User               = "test"
	//注意，ES密码需要大于6位
	Pass = "unit-test"
)

func main() {
	//startMysql()
	startES()
}

func startMysql() {
	mysqlOptions := map[string]string{
		"MYSQL_ROOT_PASSWORD": Pass,
		"MYSQL_USER":          User,
		"MYSQL_PASSWORD":      Pass,
		"MYSQL_DATABASE":      "shop",
	}
	containerOption := docker.ContainerOption{
		Name:              "mysql-for-unittest",
		Options:           mysqlOptions,
		MountVolumePath:   "/var/lib/mysql",
		PortExpose:        "3306",
		ContainerFileName: "mysql:5.7",
	}

	mysqlDocker := docker.Docker{}
	if !mysqlDocker.IsInstalled() {
		panic("docker has`t installed")
	}
	err := mysqlDocker.RemoveIfExists(containerOption)
	if err != nil {
		panic(err)
	}
	res, err := mysqlDocker.Start(containerOption)
	if err != nil {
		fmt.Println(res)
		panic(err)
	}
	fmt.Println("docker", containerOption.ContainerFileName, "has started")
	mysqlDocker.WaitForStartOrKill(StartTimeoutSecond)

	//退出时清理掉
	shutdown.NewHook().Close(func() {
		fmt.Println("exited  signal...")
		mysqlDocker.RemoveIfExists(containerOption)
	})
}

func startES() {
	containerOption := docker.ContainerOption{
		Name: "elastic-unittest",
		Options: map[string]string{
			"xpack.security.enabled": "false",
			"discovery.seed_hosts":   "127.0.0.1:9300",
			"discovery.type":         "single-node"},
		MountVolumePath:   "/usr/share/elasticsearch/data",
		PortExpose:        "9200",
		ContainerFileName: "phper95/es8.1.0",
		//ContainerFileName: "docker.elastic.co/elasticsearch/elasticsearch:8.1.0",
	}

	//cmd := fmt.Sprintf("docker exec -it %s /usr/share/elasticsearch/bin/elasticsearch-users useradd %s -p %s -r superuser", containerOption.Name, User, Pass)
	//network := "elastic"
	ESDocker := docker.Docker{}
	if !ESDocker.IsInstalled() {
		panic("docker has`t installed")
	}
	err := ESDocker.RemoveIfExists(containerOption)
	if err != nil {
		panic(err)
	}
	//ESDocker.CreateNetwork(network)
	res, err := ESDocker.Start(containerOption)
	if err != nil {
		fmt.Println(res)
		panic(err)
	}
	fmt.Println("docker", containerOption.ContainerFileName, "has started")
	//检测服务是否就绪
	if checkESServer() {
		fmt.Println("es server started")
	} else {
		fmt.Println("es server start timeout")
		//ESDocker.RemoveIfExists(containerOption)
	}

	//退出时清理掉
	shutdown.NewHook().Close(func() {
		ESDocker.RemoveIfExists(containerOption)
	})
}

func checkESServer() bool {
	url := "http://localhost:9200"
	fmt.Println(url)
	for tick := 1; tick <= StartTimeoutSecond; tick++ {
		httpCode, _, _ := fasthttp.Get(nil, url)
		if httpCode == http.StatusOK {
			return true
		}
		time.Sleep(time.Second)
		fmt.Println("check cost ", tick)
	}
	return false
}
