package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Port      int    `mapstructure:"port"`
	Mode      string `mapstructure:"mode"`
	StartTime string `mapstructure:"start_time"`
	MachineId int64  `mapstructure:"machine_id"`

	*LogConfig   `mapstructure:"log"`
	*MySqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackUps int    `mapstructure:"max_backups"`
}

type MySqlConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password""`
	DbName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_connect_count"`
	MaxIdleConns int    `mapstructure:"max_idle"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_max_size"`
}

func Init() (err error) {

	viper.SetConfigFile("config.yaml") //获取本地配置文件
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml") //指定配置文件类型，从远程服务获取配置信息的时候使用到
	viper.AddConfigPath(".") //配置文件所在路径

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Println("反序列化配置文件失败 err=", err)
		return err
	}

	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件有更新")
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Println("更新配置文件错误 err=", err)
		}
	})
	return
}
