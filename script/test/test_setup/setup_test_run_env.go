package main

import (
	"fmt"
	"gitee.com/phper95/pkg/shutdown"
	"github.com/valyala/fasthttp"
	"shop-search-api/script/test/docker"
	"strings"
	"time"
)

const (
	StartTimeoutSecond = 180
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
		Name:              "elastic-unittest",
		Options:           nil,
		MountVolumePath:   "/usr/share/elasticsearch/data",
		PortExpose:        "9200",
		ContainerFileName: "phper95/es8.1",
	}

	//cmdArgs := []string{"exec", "-it", containerOption.Name, "/usr/share/elasticsearch/bin/elasticsearch-users", "useradd", User, "-p", Pass, "-r", "superuser"}
	cmd := fmt.Sprintf("docker exec -it %s /usr/share/elasticsearch/bin/elasticsearch-users useradd %s -p %s -r superuser", containerOption.Name, User, Pass)

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
	ESDocker.WaitForStartOrKill(StartTimeoutSecond)
	//检测服务是否就绪
	if checkESServer() {
		fmt.Println("es server started")
		res, err = docker.RunCommand(cmd)
		fmt.Println("res:", res, "err:", err)
	} else {
		fmt.Println("es server start timeout")
		ESDocker.RemoveIfExists(containerOption)
	}

	//退出时清理掉
	shutdown.NewHook().Close(func() {
		ESDocker.RemoveIfExists(containerOption)
	})
}

func checkESServer() bool {
	for tick := 0; tick < StartTimeoutSecond; tick++ {
		_, _, err := fasthttp.Get(nil, "https://localhost:9200")
		if strings.Contains(err.Error(), "authority") {
			return true
		}
		time.Sleep(time.Second)
	}
	return false
}
