// Package configuration implements methods for get config variables.
//
// env file.
package configuration

import (
	"errors"
	"strconv"
	"sync"
)

const (
	envLocal      = "local"
	envStaging    = "staging"
	envProduction = "production"
	envFile       = ".env"
	envRegexp     = `(?:APP_ENV=(?P<app_env>local|staging|production))|
(?:DB_HOST=(?P<db_host>[\S]+))|
(?:DB_DATABASE=(?P<db_database>[\S]+))|
(?:DB_USERNAME=(?P<db_username>[\S]+))|
(?:DB_PASSWORD=(?P<db_password>[\S]+))|
(?:MAIL_USERNAME=(?P<mail_username>[\S]+))|
(?:MAIL_PASSWORD=(?P<mail_password>[\S]+))|
(?:MAIL_ENABLED=(?P<mail_enabled>true|false))|
(?:RABBITMQ.HOST=(?P<rabbitmq_host>[\S]+))|
(?:RABBITMQ.PORT=(?P<rabbitmq_port>[\d]+))|
(?:RABBITMQ.USERNAME=(?P<rabbitmq_username>[\S]+))|
(?:RABBITMQ.PASSWORD=(?P<rabbitmq_password>[\S]+))`
)

var (
	structEnv   env
	errEmptyEnv = errors.New("empty env configuration")
	onceEnv     sync.Once
)

// Env is an alias to export env to outer space.
type Env = env

type env struct {
	env  string
	Db   db
	Rmq  rmq
	Mail mail
}

func (e *env) setEnv(env string) {
	e.env = env
}

func (e env) IsLocal() bool {
	return e.env == envLocal
}

func (e env) IsStaging() bool {
	return e.env == envStaging
}

func (e env) IsProduction() bool {
	return e.env == envProduction
}

type credential struct {
	username string
	password string
}

func (c credential) Username() string {
	return c.username
}

func (c credential) Password() string {
	return c.password
}

type db struct {
	Host        string
	Database    string
	Credentials credential
}

type rmq struct {
	Host        string
	Port        uint64
	Credentials credential
}

type mail struct {
	enabled     bool
	Credentials credential
}

func (m mail) IsEnabled() bool {
	return m.enabled
}

func envLoad() {
	buffer, err := cfg(envFile, envRegexp)

	if err != nil {
		panic(errEmptyEnv)
	}

	structEnv = env{
		env: string(buffer["app_env"]),
		Db: db{
			Host:     string(buffer["db_host"]),
			Database: string(buffer["db_database"]),
			Credentials: credential{
				username: string(buffer["db_username"]),
				password: string(buffer["db_password"]),
			},
		},
		Mail: mail{
			Credentials: credential{
				username: string(buffer["mail_username"]),
				password: string(buffer["mail_password"]),
			},
		},
		Rmq: rmq{
			Host: string(buffer["rabbitmq_host"]),
			Credentials: credential{
				username: string(buffer["rabbitmq_username"]),
				password: string(buffer["rabbitmq_password"]),
			},
		},
	}

	b, _ := strconv.ParseBool(string(buffer["mail_enabled"]))
	structEnv.Mail.enabled = b

	u, _ := strconv.ParseUint(string(buffer["rabbitmq_port"]), 10, 16)
	structEnv.Rmq.Port = u

	if (env{}) == structEnv {
		panic(errEmptyEnv)
	}
}

// EnvConfiguration expose ENV configuration data.
func EnvConfiguration() *env {
	onceEnv.Do(func() {
		envLoad()
	})

	return &structEnv
}
