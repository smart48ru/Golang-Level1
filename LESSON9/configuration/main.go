package configuration

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/url"
)

var Cfg = Config{}

type Config struct {
	Port        string `yaml:"port" json:"port"`
	DbUrl       string `yaml:"db_url" json:"dburl"`
	JaegerUrl   string `yaml:"jaeger_url" json:"jaegerurl"`
	SentryUrl   string `yaml:"sentry_url" json:"sentryurl"`
	KafkaBroker string `yaml:"kafka_broker" json:"kafkabroker"`
	SomeAppId   string `yaml:"some_app_id" json:"someappid"`
	SomeAppKey  string `yaml:"some_app_key" json:"someappkey"`
}

const (
	defaultPort     = "8080"
	defaultJaeger   = "http://jaeger:16686"
	defaultSentry   = "http://sentry:9000"
	defaultDBurl    = "sql://db-user:db-password@sql-db:3306"
	defaultKafka    = "kafka:9092"
	defaultConfFile = "config.yaml"
	defaultAppId    = "tesapptid"
	defaultAppKey   = "testkey"
)

//String - метод вывода конфига (незнаю долден ли он возвращать ошибку)!!
func (с *Config) String() (string, error) {
	d, err := yaml.Marshal(с)
	if err != nil {
		return " ", err
	}
	return fmt.Sprintf("%s", string(d)), nil
}

func GetConfiguration() (*Config, error) {
	conf := flag.Bool("c", false, "configuration file mode")
	confFile := flag.String("cfg", defaultConfFile, "config file name")
	confFormat := flag.String("format", defaultConfFile, "config file format. Supported json, yaml ")
	flag.StringVar(&Cfg.Port, "port", defaultPort, "network port")
	flag.StringVar(&Cfg.DbUrl, "dburl", defaultDBurl, "database url")
	flag.StringVar(&Cfg.JaegerUrl, "jaeger", defaultJaeger, "jaeger url")
	flag.StringVar(&Cfg.SentryUrl, "sentry", defaultSentry, "sentry url")
	flag.StringVar(&Cfg.KafkaBroker, "kafka", defaultKafka, "kafka broker")
	flag.StringVar(&Cfg.SomeAppId, "appid", defaultAppId, "app id")
	flag.StringVar(&Cfg.SomeAppKey, "appkey", defaultAppKey, "app key")
	flag.Parse()

	if *conf {
		content, err := checkCfgFromFile(*confFile, *confFormat)
		if err != nil {
			return nil, err
		}
		return content, nil
	} else {
		err := paramValidation()
		if err != nil {
			return nil, err
		}
	}
	return &Cfg, nil
}

//checkUrl - Проверяет валидность url на Scheme и Host, возвращает false если схема не верный формат схемы или хостнаме не указан
func checkUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

//checkCfgFromFile - Функция проверки формата и загрузки файла
func checkCfgFromFile(f, format string) (*Config, error) {
	switch format {
	case "yaml":
		err := laodYamlConfig(f)
		if err != nil {
			return nil, err
		}
	case "json":
		err := laodJsonConfig(f)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf(`Поддерживаемый формат конфигурационных файлов yaml и json. exemple --format="json"`)
	}
	return &Cfg, nil
}

//laodYamlConfig функуия загрузки конфигурации из вайла yaml принимает имя файла, возвращает  структуру Config и ошибку
func laodYamlConfig(f string) error {
	fmt.Println("Loading yaml config file", f)
	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &Cfg)
	if err != nil {
		return fmt.Errorf("Неверный формат конфигурационного файла: %v", err)
	}
	return nil
}

//laodJsonConfig функуия загрузки конфигурации из вайла json принимает имя файла, возвращает  структуру Config и ошибку
func laodJsonConfig(f string) error {
	fmt.Println("Loading json config file", f)
	jsonFile, err := configFile(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonFile, &Cfg)
	if err != nil {
		return fmt.Errorf("Неверный формат конфигурационного файла: %v", err)
	}
	return nil
}

//configFile функция чтения и валидации параметров файла пнинимает имя файла возвращает []byte
func configFile(f string) ([]byte, error) {
	file, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения конфигурационного файла: %v ", err)
	}
	err = paramValidation()
	if err != nil {
		return nil, err
	}
	return file, nil
}

//paramValidation - Функция валидации параметров jaeger и sentry
func paramValidation() error {
	churl := checkUrl(Cfg.JaegerUrl)
	if !churl {
		return fmt.Errorf("невреное значение параметра --jaeger = %v", Cfg.JaegerUrl)
	}
	churl = checkUrl(Cfg.SentryUrl)
	if !churl {
		return fmt.Errorf("невреное значение параметра --sentry = %v", Cfg.SentryUrl)
	}
	return nil
}
