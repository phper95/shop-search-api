package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var Cfg = &Config{}

type Config struct {
	App           App           `yaml:"app"`
	Mysql         Mysql         `yaml:"mysql"`
	MongoDB       MongoDB       `yaml:"mongodb"`
	Elasticsearch Elasticsearch `yaml:"elasticsearch"`
	Redis         Redis         `yaml:"redis"`
}

type App struct {
	SK            string `yaml:"sk"`
	AK            string `yaml:"ak"`
	AppSignExpire int64  `yaml:"app_sign_expire"`
	PageSize      int    `yaml:"app_secret"`
	RunMode       string `yaml:"app_secret"`
	HttpPort      int64  `yaml:"app_secret"`
	ReadTimeout   int64
	WriteTimeout  time.Duration
	PrefixUrl     string

	RuntimeRootPath string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

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

type Elasticsearch struct {
	URL            string `yaml: "url"`
	User           string `yaml: "user"`
	Password       string `yaml: "password"`
	BulkActionNum  int    `yaml: "bulk_action_num"`
	BulkActionSize int    `yaml: "bulk_action_size"` //kb
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

// Setup initialize the configuration instance
func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("../../config/config.yml")
	//viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("Config file changed:", e.Name)
	//})
	if err := viper.Unmarshal(&Cfg); err != nil {
		logrus.Error(err)
		panic(err)
	}
}
