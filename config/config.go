package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Name string //配置文件名
}

// 初始化配置文件
func (c *Config) initConfig() error {
	// 如果指定了配置文件，解析指定的配置文件。否则解析默认配置文件
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	// 读取匹配的环境变量
	viper.AutomaticEnv()
	viper.SetEnvPrefix("RESTDEMO")
	replacer := strings.NewReplacer(".", "_") // 例: RESTDEMO_URL
	viper.SetEnvKeyReplacer(replacer)
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// 监控配置文件，热加载
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Infof("Config file changed: %s", in.Name)
	})
}

// 初始化日志
func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)
}

func Init(cfg string) error {

	c := Config{
		cfg,
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 初始化日志
	c.initLog()

	// 热加载
	c.watchConfig()

	return nil
}
