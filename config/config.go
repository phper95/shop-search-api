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
	App           App           `yaml:"app"`
	Mysql         Mysql         `mapstructure:"mysql"`
	MongoDB       MongoDB       `mapstructure:"mongodb"`
	Elasticsearch Elasticsearch `mapstructure:"elasticsearch"`
	Redis         Redis         `mapstructure:"redis"`
}

/**
这里推荐使用mapstructure作为序列化标签
yaml不支持 AppSignExpire int64  `yaml:"app_sign_expire"` 这种下划线的标签
使用mapstructure值得注意的地方是，只要标签中使用了下划线等连接符，":"后就
不能有空格。
比如： AppSignExpire int64  `yaml:"app_sign_expire"`是可以被解析的
          AppSignExpire int64  `yaml: "app_sign_expire"` 不能被解析
*/

type App struct {
	AppSignExpire   int64         `mapstructure:"app_sign_expire"`
	RunMode         string        `mapstructure:"run_mode"`
	HttpPort        int           `mapstructure:"http_port"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
	RuntimeRootPath string        `mapstructure:"runtime_root_path"`
	AppLogPath      string        `mapstructure:"app_log_path"`
}

type Mysql struct {
	DBName            string        `mapstructure:"dbname"`
	User              string        `mapstructure:"user"`
	Password          string        `mapstructure:"password"`
	Host              string        `mapstructure:"host"`
	MaxOpenConn       int           `mapstructure:"max_open_conn"`
	MaxIdleConn       int           `mapstructure:"max_idle_conn"`
	ConnMaxLifeSecond time.Duration `mapstructure:"conn_max_life_second"`
	TablePrefix       string        `mapstructure:"table_prefix"`
}

type MongoDB struct {
	DBname   string `mapstructure:"dbname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
}

type Elasticsearch struct {
	URL            string `mapstructure:"url"`
	User           string `mapstructure:"user"`
	Password       string `mapstructure:"password"`
	BulkActionNum  int    `mapstructure:"bulk_action_num"`
	BulkActionSize int    `mapstructure:"bulk_action_size"` //kb
	BulkWorkersNum int    `mapstructure:"bulk_workers_num"`
}

type Redis struct {
	Host        string `mapstructure:"host"`
	DB          int    `mapstructure:"db"`
	Password    string `mapstructure:"password"`
	MinIdleConn int    `mapstructure:"min_idle_conn"`
	PoolSize    int    `mapstructure:"pool_size"`
	MaxRetries  int    `mapstructure:"max_retries"`
}

// 加载配置，失败直接panic
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
