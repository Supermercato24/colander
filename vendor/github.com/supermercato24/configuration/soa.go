// Package configuration implements methods for get config variables.
//
// config/soa file.
package configuration

import (
	"errors"
	"path/filepath"
	"sync"
)

const (
	soaFile   = "soa.php"
	soaRegexp = `(?:[\s]{2}'UA'[\s]=>[\s]'(?P<ua>[\S]+)')|
(?:[\s]{2,}'cp24'[\s]=>[\s]'(?P<cp24>http[\S]+)')|
(?:[\s]{2,}'cp24_token'[\s]=>[\s]'(?P<cp24_token>Basic[\s][\S]+)')|
(?:[\s]{2,}'s24'[\s]=>[\s]'(?P<s24>http[\S]+)')|
(?:[\s]{2,}'s24_token'[\s]=>[\s]'(?P<s24_token>Basic[\s][\S]+)')|
(?:[\s]{2,}'sm'[\s]=>[\s]'(?P<sm>http[\S]+)')|
(?:[\s]{2,}'sm_token'[\s]=>[\s]'(?P<sm_token>Basic[\s][\S]+)')`
)

var (
	structSoa   soa
	errEmptySoa = errors.New("empty Soa configuration")
	onceSoa     sync.Once
)

type soa struct {
	ua   string
	Cp24 service
	S24  service
	Sm   service
}

func (s soa) UA() string {
	return s.ua
}

type service struct {
	domain string
	token  string
}

func (s service) Domain() string {
	return s.domain
}

func (s service) Token() string {
	return s.token
}

func (s service) IsHttp() bool {
	return s.domain[0:5] == "http:"
}

func soaLoad() {
	buffer, err := cfg(filepath.Join(DirProjectConfig, soaFile), soaRegexp)

	if err != nil {
		panic(errEmptySoa)
	}

	structSoa = soa{
		ua: string(buffer["ua"]),
		Cp24: service{
			domain: string(buffer["cp24"]),
			token:  string(buffer["cp24_token"]),
		},
		S24: service{
			domain: string(buffer["s24"]),
			token:  string(buffer["s24_token"]),
		},
		Sm: service{
			domain: string(buffer["sm"]),
			token:  string(buffer["sm_token"]),
		},
	}

	if (soa{}) == structSoa {
		panic(errEmptySoa)
	}
}

// SoaConfiguration expose SOA configuration data.
func SoaConfiguration() *soa {
	onceSoa.Do(func() {
		soaLoad()
	})

	return &structSoa
}
