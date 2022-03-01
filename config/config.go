package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var Cfg = &Config{}

type Config struct {
	App App `yaml:"app"`
	//Mysql         Mysql         `yaml:"mysql"`
	//MongoDB       MongoDB       `yaml:"mongodb"`
	//Elasticsearch Elasticsearch `yaml:"elasticsearch"`
	//Redis         Redis         `yaml:"redis"`
}

/**
这里推荐使用mapstructure作为序列化标签
yaml不支持 AppSignExpire int64  `yaml:"app_sign_expire"` 这种下划线的标签
*/

type App struct {
	SK              string        `mapstructure:"sk"`
	AK              string        `mapstructure:"ak"`
	AppSignExpire   int64         `mapstructure:"app_sign_expire"`
	RunMode         string        `mapstructure:"run_mode"`
	HttpPort        int64         `mapstructure:"http_port"`
	ReadTimeout     int64         `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
	RuntimeRootPath string        `mapstructure:"runtime_root_path"`
	LogPath         string        `mapstructure:"log_path"`
}

type Mysql struct {
	DBName      string `mapstructure: "dbname"`
	User        string `mapstructure: "user"`
	Password    string `mapstructure: "password"`
	Host        string `mapstructure: "host"`
	Port        string `mapstructure: "port"`
	TablePrefix string `mapstructure: "table_prefix"`
}

type MongoDB struct {
	DBname   string `mapstructure: "dbname"`
	User     string `mapstructure: "user"`
	Password string `mapstructure: "password"`
	Host     string `mapstructure: "host"`
	Port     string `mapstructure: "port"`
}

type Elasticsearch struct {
	URL            string `mapstructure: "url"`
	User           string `mapstructure: "user"`
	Password       string `mapstructure: "password"`
	BulkActionNum  int    `mapstructure: "bulk_action_num"`
	BulkActionSize int    `mapstructure: "bulk_action_size"` //kb
}

type Redis struct {
	Host        string        `mapstructure: "host"`
	Password    string        `mapstructure: "password"`
	MaxIdle     int           `mapstructure: "max_idle"`
	MaxActive   int           `mapstructure: "max_active"`
	IdleTimeout time.Duration `mapstructure: "idle_timeout"`
}

// Setup initialize the configuration instance
func LoadConfig() {
	viper := viper.New()
	//1.设置配置文件路径
	viper.SetConfigFile("config/config.yml")

	//2.配置读取
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	//3.将配置映射成结构体
	if err := viper.Unmarshal(&Cfg); err != nil {
		logrus.Error(err)
		panic(err)
	}

	//4. 监听配置文件变动,重新解析配置
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := viper.Unmarshal(&Cfg); err != nil {
			logrus.Error(err)
			panic(err)
		}
	})
}
