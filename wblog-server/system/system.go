package system

import (
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Zap struct {
	Level         string
	Format        string
	Prefix        string
	Director      string
	ShowLine      bool
	EncodeLevel   string
	StacktraceKey string
	LogInConsole  bool
}
type Configuration struct {
	SignupEnabled      bool   `yaml:"signup_enabled"`  // signup enabled or not
	QiniuAccessKey     string `yaml:"qiniu_accesskey"` // qiniu
	QiniuSecretKey     string `yaml:"qiniu_secretkey"`
	QiniuFileServer    string `yaml:"qiniu_fileserver"`
	QiniuBucket        string `yaml:"qiniu_bucket"`
	GithubClientId     string `yaml:"github_clientid"` // github
	GithubClientSecret string `yaml:"github_clientsecret"`
	GithubAuthUrl      string `yaml:"github_authurl"`
	GithubRedirectURL  string `yaml:"github_redirecturl"`
	GithubTokenUrl     string `yaml:"github_tokenurl"`
	GithubScope        string `yaml:"github_scope"`
	SmtpUsername       string `yaml:"smtp_username"`  // username
	SmtpPassword       string `yaml:"smtp_password"`  //password
	SmtpHost           string `yaml:"smtp_host"`      //host
	SessionSecret      string `yaml:"session_secret"` //session_secret
	Domain             string `yaml:"domain"`         //domain
	Public             string `yaml:"public"`         //public
	Addr               string `yaml:"addr"`           //addr
	BackupKey          string `yaml:"backup_key"`     //backup_key
	DSN                string `yaml:"dsn"`            //database dsn
	NotifyEmails       string `yaml:"notify_emails"`  //notify_emails
	PageSize           int    `yaml:"page_size"`      //page_size
	SmmsFileServer     string `yaml:"smms_fileserver"`
	Zap                *Zap
}

const (
	DEFAULT_PAGESIZE = 10
)

var configuration *Configuration

//func LoadConfiguration(path string) error {
//	data, err := ioutil.ReadFile(path)
//	if err != nil {
//		return err
//	}
//	var config Configuration
//	err = yaml.Unmarshal(data, &config)
//	if err != nil {
//		return err
//	}
//	if config.PageSize <= 0 {
//		config.PageSize = DEFAULT_PAGESIZE
//	}
//	configuration = &config
//	return err
//}

func GetConfiguration() *Configuration {

	conf := "./conf/conf.yml"
	viper.SetConfigFile(conf)
	content, err := ioutil.ReadFile(conf)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read conf file fail: %s", err.Error()))
	}
	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse conf file fail: %s", err.Error()))
	}

	if err = viper.Unmarshal(&configuration); err != nil {
		color.Red.Println("unable to decode into struct, %v", err)
	}
	if configuration.PageSize <= 0 {
		configuration.PageSize = DEFAULT_PAGESIZE
	}
	return configuration
}
