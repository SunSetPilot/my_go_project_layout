package conf

//import (
//	"github.com/spf13/viper"
//
//)
//
//type Config struct {
//	CommonConf CommonConfig `mapstructure:"common"`
//	DataConf   DataConfig   `mapstructure:"data"`
//}
//
//type CommonConfig struct {
//	GinPort      string // gin端口 :8080
//	DebugAddress string // debug端口
//	LogLevel     string // 日志级别
//}
//
//type DataConfig struct {
//	MysqlDSN string // mysql连接字符串
//	KconfKey string // kconf的键名
//}
//
//func Init() error {
//	var (
//		viperConf *viper.Viper
//		err       error
//		config    Config
//	)
//	viperConf = viper.New()
//	viperConf.SetConfigFile(global.ConfigFile)
//
//	if err = viperConf.ReadInConfig(); err != nil {
//		return err
//	}
//	if err = viperConf.Unmarshal(&config); err != nil {
//		return err
//	}
//
//	// 提高环境变量优先级
//	config.CommonConf.GinPort = utils.ConvertPortToAddr(global.APIPort)
//
//	global.ProjectConfig.Store(config)
//
//	return nil
//}
