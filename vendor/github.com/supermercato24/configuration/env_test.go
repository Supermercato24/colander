package configuration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	envMockup = []byte(`APP_ENV=local
DB_HOST=127.0.0.1
DB_DATABASE=foobar
DB_USERNAME=root
DB_PASSWORD=toor
MAIL_USERNAME=Foo
MAIL_PASSWORD=Bar
MAIL_ENABLED=true
RABBITMQ.HOST=127.0.0.1
RABBITMQ.PORT=5672
RABBITMQ.USERNAME=guest
RABBITMQ.PASSWORD=guest`)
)

func EnvModel(t *testing.T, env *env) {
	// env
	if env.IsProduction() {
		t.Error("env must be staging or local")
	}

	// db
	assert.NotEmpty(t, env.Db.Host)
	assert.NotEmpty(t, env.Db.Database)
	assert.NotEmpty(t, env.Db.Credentials.Username())
	assert.Exactly(t, "root", env.Db.Credentials.Username())
	assert.NotEmpty(t, env.Db.Credentials.Password())
	assert.Exactly(t, "toor", env.Db.Credentials.Password())

	// mail
	assert.NotEmpty(t, env.Mail.Credentials.Username())
	assert.Exactly(t, "Foo", env.Mail.Credentials.Username())
	assert.NotEmpty(t, env.Mail.Credentials.Password())
	assert.Exactly(t, "Bar", env.Mail.Credentials.Password())

	// Username
	assert.NotEqual(t, env.Db.Credentials.Username(), env.Mail.Credentials.Username())

	// Password
	assert.NotEqual(t, env.Db.Credentials.Password(), env.Mail.Credentials.Password())
}

func EnvMockup(t *testing.T) {
	fd, err := os.Open(envFile)
	if err != nil {
		fd, err := os.Create(envFile)
		assert.NoError(t, err)
		fd.Write(envMockup)
		assert.NoError(t, fd.Close())
	} else {
		assert.NoError(t, fd.Close())
	}
}

func TestEnv(t *testing.T) {
	assert.True(t, t.Run("mockup", EnvMockup))
	assert.True(t, assert.NotPanics(t, func() { EnvConfiguration() }))
	assert.NotPanics(t, func() { EnvModel(t, EnvConfiguration()) })
}

func BenchmarkEnv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EnvConfiguration()
	}
}
