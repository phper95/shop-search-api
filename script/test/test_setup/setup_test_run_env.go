package main

import (
	"fmt"
	"gitee.com/phper95/pkg/shutdown"
	"shop-search-api/script/test/docker"
)

const StartTimeoutSecond = 180

func main() {

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
	shutdown.NewHook().Close(func() {
		fmt.Println("exited")
	})

}
