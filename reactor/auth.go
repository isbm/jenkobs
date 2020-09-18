package jenkobs_reactor

import (
	"github.com/isbm/go-nanoconf"
)

// AMQPAuth object
type AMQPAuth struct {
	User         string
	Password     string
	Fqdn         string
	Port         int
	ExchangeName string
	Vhost        string
}

// NewAMQPAuth constructor to AMQPAuth
func NewAMQPAuth(conf *nanoconf.Inspector) *AMQPAuth {
	mqa := new(AMQPAuth)
	mqa.User = conf.String("username", "")
	mqa.Password = conf.String("password", "")
	mqa.Fqdn = conf.String("fqdn", "")
	mqa.Port = conf.DefaultInt("port", "", 5672)
	mqa.ExchangeName = conf.String("exchange", "")
	mqa.Vhost = conf.String("vhost", "")

	return mqa
}

// JenkinsAuth object
type JenkinsAuth struct {
	User  string
	Token string
	Fqdn  string
	Port  int
}

// NewJenkinsAuth constructor to JenkinsAuth
func NewJenkinsAuth(conf *nanoconf.Inspector) *JenkinsAuth {
	jks := new(JenkinsAuth)
	jks.User = conf.String("username", "")
	jks.Token = conf.String("token", "")
	jks.Fqdn = conf.String("fqdn", "")
	jks.Port = conf.DefaultInt("port", "", 443)

	return jks
}
