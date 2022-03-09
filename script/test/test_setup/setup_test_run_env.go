package main

import (
	"fmt"
	"gitee.com/phper95/pkg/shutdown"
	"shop-search-api/script/test/docker"
)

const StartTimeoutSecond = 180

func main() {
	//startMysql()
	startES()
}

func startMysql() {
	mysqlOptions := map[string]string{
		"MYSQL_ROOT_PASSWORD": "root",
		"MYSQL_USER":          "test",
		"MYSQL_PASSWORD":      "test",
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

	//[run -d --name mysql-unittest
	//	-p 3306:3306 -e MYSQL_USER = test -e MYSQL_PASSWORD = test
	//	-e MYSQL_DATABASE = shop
	//	-e MYSQL_ROOT_PASSWORD = root
	//	--tmpfs /var /lib/mysql mysql:5.7
	//	]
	//	docker network create elastic
	//	docker run --name elastic-unittest --net elastic -p 9200:9200 -it docker.elastic.co/elasticsearch/elasticsearch:8.1.0

	//EsOptions := map[string]string{
	//	"MYSQL_ROOT_PASSWORD": "root",
	//	"MYSQL_USER":          "test",
	//	"MYSQL_PASSWORD":      "test",
	//	"MYSQL_DATABASE":      "shop",
	//}
	--net = host

	containerOption := docker.ContainerOption{
		Name:              "elastic-unittest",
		Options:           nil,
		MountVolumePath:   "/usr/share/elasticsearch/data",
		PortExpose:        "9200",
		ContainerFileName: "docker.elastic.co/elasticsearch/elasticsearch:8.1.0",
	}

	ESDocker := docker.Docker{}
	if !ESDocker.IsInstalled() {
		panic("docker has`t installed")
	}
	err := ESDocker.RemoveIfExists(containerOption)
	if err != nil {
		panic(err)
	}
	res, err := ESDocker.Start(containerOption)
	if err != nil {
		fmt.Println(res)
		panic(err)
	}
	fmt.Println("docker", containerOption.ContainerFileName, "has started")
	ESDocker.WaitForStartOrKill(StartTimeoutSecond)

	//退出时清理掉
	shutdown.NewHook().Close(func() {
		fmt.Println("exited  signal...")
		ESDocker.RemoveIfExists(containerOption)
	})
}
