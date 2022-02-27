package config

import (
	"log"
	"time"
)

type Config struct {
	App    App    `yaml:"app"`
	Server Server `yaml:"server"`
	Mysql  Mysql  `yaml:"mysql"`
	Server Server `yaml:"server"`
	Server Server `yaml:"server"`
	Server Server `yaml:"server"`
}

type App struct {
	AppSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Mysql struct {
	DBName      string `yaml: "dbname"`
	User        string `yaml: "user"`
	Password    string `yaml: "password"`
	Host        string `yaml: "host"`
	Port        string `yaml: "port"`
	TablePrefix string
}

type MongoDB struct {
	DBname   string `yaml: "dbname"`
	User     string `yaml: "user"`
	Password string `yaml: "password"`
	Host     string `yaml: "host"`
	Port     string `yaml: "port"`
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
