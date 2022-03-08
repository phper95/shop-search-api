package main

import (
	"shop-search-api/script/test/docker"
	"testing"
)

func TestInit(m *testing.M) {

	mysqlDocker := docker.MysqlDocker{docker.Docker{
		ContainerID:   "",
		ContainerName: "",
	}}
	mysqlDocker.StartMysqlDocker()
}
