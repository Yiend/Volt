package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

var (
	global *Config
)


// 加载全局配置
// @Param fpath "文件路径"
func LoadConfig(fpath string) error {
	c, err := ParseConfig(fpath)
	if err != nil {
		return err
	}
	global = c
	return nil
}

// 获取全局配置
func GetConfig() *Config {
	if global == nil {
		return &Config{}
	}
	return global
}

// 解析配置文件
func ParseConfig(fpath string) (*Config, error) {
	var c Config
	_, err := toml.DecodeFile(fpath, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}


// Config 配置参数
type Config struct {
	HTTP            HTTP        `toml:"http"`
	Xorm            Xorm        `toml:"xorm"`
	MySQL           MySQL       `toml:"mysql"`
}


// HTTP http配置参数
type HTTP struct {
	Host            string `toml:"host"`
	Port            int    `toml:"port"`
	ShutdownTimeout int    `toml:"shutdown_timeout"`
}

// Gorm gorm配置参数
type Xorm struct {
	ShowSql      bool   `toml:"show_sql"`
	ShowExecTime bool   `toml:"show_exec_time"`
	MaxLifetime  int    `toml:"max_lifetime"`
	MaxOpenConns int    `toml:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	TablePrefix  string `toml:"table_prefix"`
}

// MySQL mysql配置参数
type MySQL struct {
	Host       string `toml:"host"`
	Port       int    `toml:"port"`
	User       string `toml:"user"`
	Password   string `toml:"password"`
	DBName     string `toml:"db_name"`
	Parameters string `toml:"parameters"`
}

// DSN 数据库连接串
func (a MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}