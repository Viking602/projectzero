package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func LoadConfig(file string, c Conf) *Conf {
	// 读取配置文件
	v := viper.New()
	v.SetConfigFile(file)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置解析失败:%s\n", err))
	}
	// 动态监测配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生改变")
		if err := v.Unmarshal(&c); err != nil {
			panic(fmt.Errorf("配置重载失败:%s\n", err))
		}
	})
	if err := v.Unmarshal(&c); err != nil {
		panic(fmt.Errorf("配置重载失败:%s\n", err))
	}

	return &c
}
