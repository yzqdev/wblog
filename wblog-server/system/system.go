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

type Mysql struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}
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
	Reg                bool   `mapstructure:"reg"`              // signup enabled or not
	QiniuAccessKey     string `mapstructure:"qiniu_accesskey" ` // qiniu
	QiniuSecretKey     string `mapstructure:"qiniu_secretkey"`
	QiniuFileServer    string `mapstructure:"qiniu_fileserver"`
	QiniuBucket        string `mapstructure:"qiniu_bucket"`
	GithubClientId     string `mapstructure:"github_clientid"` // github
	GithubClientSecret string `mapstructure:"github_clientsecret"`
	GithubAuthUrl      string `mapstructure:"github_authurl"`
	GithubRedirectURL  string `mapstructure:"github_redirecturl"`
	GithubTokenUrl     string `mapstructure:"github_token_url" `
	GithubScope        string `mapstructure:"github_scope"`
	SmtpUsername       string `mapstructure:"smtp_username"`  // username
	SmtpPassword       string `mapstructure:"smtp_password"`  //password
	SmtpHost           string `mapstructure:"smtp_host"`      //host
	SessionSecret      string `mapstructure:"session_secret"` //session_secret
	Domain             string `mapstructure:"domain"`         //domain
	Public             string `mapstructure:"public"`         //public
	Addr               string `mapstructure:"addr"`           //addr
	BackupKey          string `mapstructure:"backup_key"`     //backup_key
	DSN                string `mapstructure:"dsn"`            //database dsn
	NotifyEmails       string `mapstructure:"notify_emails"`  //notify_emails
	PageSize           int    `mapstructure:"page_size"`      //page_size
	SmmsFileServer     string `mapstructure:"smms_fileserver"`
	Zap                *Zap   `mapstructure:"zap"`
	Mysql              *Mysql `mapstructure:"mysql"`
}

const (
	DEFAULT_PAGESIZE = 10
)

var configuration *Configuration

func GetConfiguration() *Configuration {

	conf := "./conf/config.yml"
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
