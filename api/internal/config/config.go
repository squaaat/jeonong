package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/spf13/viper"

	_const "github.com/squaaat/nearsfeed/api/internal/const"
	"github.com/squaaat/nearsfeed/api/internal/er"
)

func New(e string) (*Config, error) {
	op := er.CallerOp()
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	s := ssm.New(sess)
	fmt.Println(op, fmt.Sprintf("/%s/%s/%s/application.yml", _const.Project, _const.App, e))
	param, err := s.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(fmt.Sprintf("/%s/%s/%s/application.yml", _const.Project, _const.App, e)),
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	value := *(param.Parameter.Value)
	viper.SetConfigType("yaml")
	viper.ReadConfig(strings.NewReader(value))

	return &Config{
		Version: viper.GetString("version"),
		App: &AppConfig{
			Env:     viper.GetString("env.app.env"),
			Debug:   viper.GetBool("env.app.debug"),
			Project: viper.GetString("env.app.project"),
			AppName: viper.GetString("env.app.app_name"),
		},
		ServerHTTP: &ServerHTTPConfig{
			Port:    viper.GetString("env.server_http.port"),
			Timeout: viper.GetDuration("env.server_http.timeout"),
		},
		ServiceDB: &ServiceDBConfig{
			Host:     viper.GetString("env.service_db.host"),
			Port:     viper.GetString("env.service_db.port"),
			Dialect:  viper.GetString("env.service_db.dialect"),
			Schema:   viper.GetString("env.service_db.schema"),
			Username: viper.GetString("env.service_db.username"),
			Password: viper.GetString("env.service_db.password"),
		},
	}, nil
}

type ServerHTTPConfig struct {
	Port    string
	Timeout time.Duration
}

type ServiceDBConfig struct {
	Host     string
	Port     string
	Dialect  string
	Schema   string
	Username string
	Password string
}

type AppConfig struct {
	Env     string
	Debug   bool
	Project string
	AppName string
}

type Config struct {
	Version    string
	CICD       bool
	App        *AppConfig
	ServerHTTP *ServerHTTPConfig
	ServiceDB  *ServiceDBConfig
}
