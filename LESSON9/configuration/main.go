package configuration

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/url"
	"strings"
)

var Cfg = Config{}
var ErrCfg = Config{}
var FileCfg = Config{}

type Config struct {
	Port        string `yaml:"port",json:"port"`
	DbUrl       string `yaml:"db_url",json:"dburl"`
	JaegerUrl   string `yaml:"jaeger_url",json:"jaegerurl"`
	SentryUrl   string `yaml:"sentry_url",json:"sentryurl"`
	KafkaBroker string `yaml:"kafka_broker",json:"kafkabroker"`
	SomeAppId   string `yaml:"some_app_id",json:"someappid"`
	SomeAppKey  string `yaml:"some_app_key",json:"someappkey"`
}

const (
	defPort     = "8080"
	defJaeger   = "http://jaeger:16686"
	defSentry   = "http://sentry:9000"
	defDburl    = "sql://db-user:db-password@sql-db:3306"
	defKafka    = "kafka:9092"
	defConfFile = "config.yaml"
	defAppId    = "tesapptid"
	defAppKey   = "testkey"
)

//String - метод вывода конфига (незнаю долден ли он возвращать ошибку)!!
func (с Config) String() (string, error) {
	d, err := yaml.Marshal(с)
	if err != nil {
		return "", fmt.Errorf("%e", err)
	}
	return fmt.Sprintf("%s", string(d)), nil
}

func GetConfiguration() (Config, error) {
	conf := flag.Bool("c", false, "configuration file mode")
	confFile := flag.String("cfg", defConfFile, "config file name")
	flag.StringVar(&Cfg.Port, "port", defPort, "network port")
	flag.StringVar(&Cfg.DbUrl, "dburl", defDburl, "database url")
	flag.StringVar(&Cfg.JaegerUrl, "jaeger", defJaeger, "jaeger url")
	flag.StringVar(&Cfg.SentryUrl, "sentry", defSentry, "sentry url")
	flag.StringVar(&Cfg.KafkaBroker, "kafka", defKafka, "kafka broker")
	flag.StringVar(&Cfg.SomeAppId, "appid", defAppId, "app id")
	flag.StringVar(&Cfg.SomeAppKey, "appkey", defAppKey, "app key")
	flag.Parse()

	if *conf {
		content, err := checkExeption(*confFile)
		if err != nil {
			return content, err
		}
		return content, nil
	} else {
		churl := checkUrl(Cfg.JaegerUrl)
		if !churl {
			return ErrCfg, fmt.Errorf("невреное значение параметра --jaeger", Cfg.JaegerUrl)
		}
		churl = checkUrl(Cfg.SentryUrl)
		if !churl {
			return ErrCfg, fmt.Errorf("невреное значение параметра --sentry", Cfg.SentryUrl)
		}
	}
	return Cfg, nil
}

//checkUrl - Проверяет валидность url на Scheme и Host, возвращает false если схема не верный формат схемы или хостнаме не указан
func checkUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

//checkExeption функция проверяет расширение переданного конфигурационного файла
func checkExeption(f string) (Config, error) {
	extension := strings.Split(f, ".")
	lenSplit := len(extension)
	if lenSplit > 1 && extension[lenSplit-1] == "yaml" {
		_, err := laodYamlConfig(f)
		if err != nil {
			return ErrCfg, err
		}
	} else if lenSplit > 1 && extension[lenSplit-1] == "json" {
		_, err := laodJsonConfig(f)
		if err != nil {
			return ErrCfg, err
		}
	} else {
		return ErrCfg, fmt.Errorf("Поддерживаемое расширение и формат конфигурационных файлов .yaml и .json")
	}
	return FileCfg, nil
}

//laodYamlConfig функуия загрузки конфигурации из вайла yaml принимает имя файла, возвращает  структуру Config и ошибку
func laodYamlConfig(f string) (Config, error) {
	fmt.Println("Loading config file", f)
	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		return ErrCfg, err
	}
	err = yaml.Unmarshal(yamlFile, &FileCfg)
	if err != nil {
		return ErrCfg, err
	}
	churl := checkUrl(FileCfg.JaegerUrl)
	if !churl {
		return ErrCfg, fmt.Errorf("невреное значение параметра --jaeger ", FileCfg.JaegerUrl)
	}
	churl = checkUrl(FileCfg.SentryUrl)
	if !churl {
		return ErrCfg, fmt.Errorf("невреное значение параметра --sentry ", FileCfg.SentryUrl)
	}
	return FileCfg, nil
}

//laodJsonConfig функуия загрузки конфигурации из вайла json принимает имя файла, возвращает  структуру Config и ошибку
func laodJsonConfig(f string) (Config, error) {
	fmt.Println("Loading config file", f)
	jsonFile, err := configFile(f)
	if err != nil {
		return ErrCfg, err
	}
	err = json.Unmarshal(jsonFile, &FileCfg)
	if err != nil {
		return ErrCfg, fmt.Errorf("Неверный формат конфигурационного файла : ", err)
	}
	churl := checkUrl(FileCfg.JaegerUrl)
	if !churl {
		return ErrCfg, fmt.Errorf("невреное значение параметра --jaeger ", FileCfg.JaegerUrl)
	}
	churl = checkUrl(FileCfg.SentryUrl)
	if !churl {
		return ErrCfg, fmt.Errorf("невреное значение параметра --sentry ", FileCfg.SentryUrl)
	}
	return FileCfg, nil
}

//configFile функция чтения файла пнинимает имя файла возвращает []byte
func configFile(f string) ([]byte, error) {
	jsonFile, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения конфигурационного файла : ", err)
	}
	return jsonFile, nil
}
