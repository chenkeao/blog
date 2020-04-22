package system

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

type Configuration struct {
	Mode           bool   `yaml:"debug_mode"` //是否为debug模式
	DbUser         string `yaml:"db_user"`
	DbPassword     string `yaml:"db_password"`
	DbName         string `yaml:"db_name"`
	DbHost         string `yaml:"db_host"`
	DbPort         string `yaml:"db_port"`
	SignupEnabled  bool   `yaml:"signup_enabled"` // signup enabled or not
	SmtpUsername   string `yaml:"smtp_username"`  // username
	SmtpPassword   string `yaml:"smtp_password"`  //password
	SmtpHost       string `yaml:"smtp_host"`      //host
	SessionSecret  string `yaml:"session_secret"` //session_secret
	Domain         string `yaml:"domain"`         //domain
	Public         string `yaml:"public"`         //public
	Addr           string `yaml:"addr"`           //addr
	BackupKey      string `yaml:"backup_key"`     //backup_key
	DSN            string `yaml:"dsn"`            //database dsn
	NotifyEmails   string `yaml:"notify_emails"`  //notify_emails
	PageSize       int    `yaml:"page_size"`      //page_size
	SmmsFileServer string `yaml:"smms_fileserver"`
}

const (
	DefaultPageSize = 10
)

var configuration *Configuration

func LoadConfiguration(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	if config.PageSize <= 0 {
		config.PageSize = DefaultPageSize
	}
	configuration = &config
	return err
}

func GetConfiguration() *Configuration {
	return configuration
}
