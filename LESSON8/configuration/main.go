package configuration

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/url"
	"os"
)

type Config struct {
	Port        string `yaml:"port"`
	DbUrl       string `yaml:"db_url"`
	JaegerUrl   string `yaml:"jaeger_url"`
	SentryUrl   string `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker"`
	SomeAppId   string `yaml:"some_app_id"`
	SomeAppKey  string `yaml:"some_app_key"`
}

var (
	defPort = "8080"
	defJaeger = "http://jaeger:16686"
	defSentry = "http://sentry:9000"
	defDburl = "sql://db-user:db-password@sql-db:3306"
	defKafka = "kafka:9092"
	defConfFile = "config.yaml"
	defAppId = "tesapptid"
	defAppKey = "testkey"
)


func ConfigurationFlag() Config {
	Cfg := Config{}

	conf := flag.Bool("c", false, "configuration file mode" )
	confFile := flag.String("cfg",defConfFile, "config file name" )
	flag.StringVar(&Cfg.Port, "port", defPort, "network port")
	flag.StringVar(&Cfg.DbUrl, "dburl", defDburl, "database url")
	flag.StringVar(&Cfg.JaegerUrl, "jaeger", defJaeger, "jaeger url")
	flag.StringVar(&Cfg.SentryUrl, "sentry", defSentry, "sentry url")
	flag.StringVar(&Cfg.KafkaBroker, "kafka", defKafka, "kafka broker")
	flag.StringVar(&Cfg.SomeAppId, "appid", defAppId, "app id")
	flag.StringVar(&Cfg.SomeAppKey, "appkey", defAppKey, "app key")
	flag.Parse()

	if *conf{
		fmt.Println("Loading config file", *confFile)
		fileCfg := Config{}
		yamlFile, err := ioutil.ReadFile(*confFile)
		if err != nil {
			fmt.Println("Ошибка чтения конфигурационного файла : ",err)
		}
		err = yaml.Unmarshal(yamlFile, &fileCfg)
		if err != nil {
			fmt.Println("Неверный формат нфигурационного файла : ",err)
		}
		return fileCfg
	} else {
		fmt.Println("NO")
		if len(Cfg.JaegerUrl) != 0 {
			churl := checkUrl(Cfg.JaegerUrl)
			if !churl {
				fmt.Println("невреное значение параметра --jaeger", Cfg.JaegerUrl)
				os.Exit(0)
			}
		}

		if len(Cfg.SentryUrl) != 0 {
			churl := checkUrl(Cfg.SentryUrl)
			if !churl {
				fmt.Println("невреное значение параметра --sentry", Cfg.SentryUrl)
				os.Exit(0)
			}
		}
	}
	return Cfg
}

//checkUrl - Проверяет валидность url на Scheme и Host, возвращает false если схема не верный формат схемы или хостнаме не указан
func checkUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}