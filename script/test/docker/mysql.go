package docker

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
	m.Docker.Start(containerOption)
	m.Docker.WaitForStartOrKill(mysqlStartTimeoutSecond)
}

func (m *MysqlDocker) Stop() {
	m.Docker.Stop()
}
