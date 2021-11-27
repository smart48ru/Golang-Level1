package configuration

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"net/url"
)

var Cfg = Config{}
var ErrCfg = Config{}

type Config struct {
	Port        string
	DbUrl       string
	JaegerUrl   string
	SentryUrl   string
	KafkaBroker string
	SomeAppId   string
	SomeAppKey  string
}
//String - метод вывода конфига (незнаю долден ли он возвращать ошибку)!!
func (с Config) String() (string, error) {
	d, err := yaml.Marshal(с)
	if err != nil {
		//log.Fatalf("error: %v", err)
		return "", fmt.Errorf("%e", err)
	}
	return fmt.Sprintf("%s", string(d)), nil
}

const (
	defPort   = "8080"
	defJaeger = "http://jaeger:16686"
	defSentry = "http://sentry:9000"
	defDburl  = "sql://db-user:db-password@sql-db:3306"
	defKafka  = "kafka:9092"
	defAppId  = "tesapptid"
	defAppKey = "testkey"
)

func ConfigurationFlag() (Config, error) {
	flag.StringVar(&Cfg.Port, "port", defPort, "network port")
	flag.StringVar(&Cfg.DbUrl, "dburl", defDburl, "database url")
	flag.StringVar(&Cfg.JaegerUrl, "jaeger", defJaeger, "jaeger url")
	flag.StringVar(&Cfg.SentryUrl, "sentry", defSentry, "sentry url")
	flag.StringVar(&Cfg.KafkaBroker, "kafka", defKafka, "kafka broker")
	flag.StringVar(&Cfg.SomeAppId, "appid", defAppId, "app id")
	flag.StringVar(&Cfg.SomeAppKey, "appkey", defAppKey, "app key")
	flag.Parse()

	if len(Cfg.JaegerUrl) != 0 {
		churl := checkUrl(Cfg.JaegerUrl)
		if !churl {
			err := fmt.Errorf("невреное значение параметра --jaeger %q", Cfg.JaegerUrl)
			return ErrCfg, err
		}
	}
	if len(Cfg.SentryUrl) != 0 {
		churl := checkUrl(Cfg.SentryUrl)
		if !churl {
			err := fmt.Errorf("невреное значение параметра  --sentry %q", Cfg.SentryUrl)
			return ErrCfg, err
		}
	}
	return Cfg, nil
}

//checkUrl - Проверяет валидность url на Scheme и Host, возвращает false если схема не верный формат схемы или хостнаме не указан
func checkUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
