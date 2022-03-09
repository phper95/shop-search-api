package docker

import "fmt"

const mysqlStartTimeoutSecond = 10

type MysqlDocker struct {
	Docker Docker
}

func (m *MysqlDocker) StartMysqlDocker() {
	mysqlOptions := map[string]string{
		"MYSQL_ROOT_PASSWORD": "root",
		"MYSQL_USER":          "go",
		"MYSQL_PASSWORD":      "root",
		"MYSQL_DATABASE":      "godb",
	}
	containerOption := ContainerOption{
		Name:              "mysql-for-unittest",
		Options:           mysqlOptions,
		MountVolumePath:   "/var/lib/mysql",
		PortExpose:        "3306",
		ContainerFileName: "mysql:5.7",
	}
	m.Docker = Docker{}
	res, err := m.Docker.Start(containerOption)
	if err != nil {
		fmt.Println("Docker.Start error", err, res)
	}
	err = m.Docker.WaitForStartOrKill(mysqlStartTimeoutSecond)
	if err != nil {
		fmt.Println("WaitForStartOrKill", err)
	}
}

func (m *MysqlDocker) Stop() {
	m.Docker.Stop()
}
