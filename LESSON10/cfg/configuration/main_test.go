package configuration_test

import (
	. "github.com/smart48ru/Golang-Level1/LESSON10/cfg/configuration"
	"testing"
)

var cfg = Cfg
var sentryUrl = Cfg.SentryUrl

func TestCheckUrl(t *testing.T) {
	var (
		testTable = []struct {
			name string
			url  string
			want bool
		}{
			// Первый сценарий
			{
				name: "http://google.ru",
				url:  "http://google.ru",
				want: true,
			},
			// Второй сценарий
			{
				name: "http:///google.ru",
				url:  "http:///google.ru",
				want: false,
			},
			// Третий сценарий
			{
				name: "http://",
				url:  "http://",
				want: false,
			},
			// Четвертый сценарий
			{
				name: "",
				url:  "",
				want: false,
			},
		}
	)
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if CheckUrl(testCase.url) != testCase.want {
				t.Errorf("Division() = %v, want %v", CheckUrl(testCase.url), testCase.want)
			}
		})
	}
}

func TestСheckCfgFromFile(t *testing.T) {
	testTable := []struct {
		name      string
		format    string
		filename  string
		want      error
		jaegerUrl string
		sentryUrl string
		dbUrl string
		kafkaBroker string
		someAppId string
		someAppKey string
		port string
	}{
		//Первый сценарий
		{
			name:      "../config.json",
			format:    "json",
			filename:  "../config.json",
			want:      nil,
			jaegerUrl: "http://jaeger:26686",
			sentryUrl: "http://sentry:9002",
			dbUrl: "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable",
			kafkaBroker: "kafka:9092",
			someAppId: "testid",
			someAppKey: "testkey",
			port: "8084",
		},
		//Второй сценарий
		{
			name:      "../config.yaml",
			format:    "yaml",
			filename:  "../config.yaml",
			want:      nil,
			jaegerUrl: "http://jaeger:26686",
			sentryUrl: "http://sentry:9002",
			dbUrl: "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable",
			kafkaBroker: "kafka:9092",
			someAppId: "testid",
			someAppKey: "testkey",
			port: "8086",
		},
	}

	for _, testCase := range testTable {
			t.Run(testCase.name, func(t *testing.T) {
				file, err := СheckCfgFromFile(testCase.filename, testCase.format)
				if err != testCase.want {
					t.Errorf("ConfigFile() = %v, want %v", err, testCase.want)
				}
				if file.Port != testCase.port{
					t.Errorf("ConfigFile.Port = %v, want %v", file.Port, testCase.port)
				}
				if file.JaegerUrl != testCase.jaegerUrl{
					t.Errorf("ConfigFile.jaegerUrl = %v, want %v", file.JaegerUrl, testCase.jaegerUrl)
				}
				if file.SentryUrl != testCase.sentryUrl{
					t.Errorf("ConfigFile.SaegerUrl = %v, want %v", file.SentryUrl, testCase.sentryUrl)
				}
				if file.DbUrl != testCase.dbUrl{
					t.Errorf("ConfigFile.DbUrl = %v, want %v", file.DbUrl, testCase.dbUrl)
				}
				if file.KafkaBroker != testCase.kafkaBroker{
					t.Errorf("ConfigFile.KafkaBroker = %v, want %v", file.KafkaBroker, testCase.kafkaBroker)
				}
				if file.SomeAppId != testCase.someAppId{
					t.Errorf("ConfigFile.SomeAppId = %v, want %v", file.SomeAppId, testCase.someAppId)
				}
				if file.SomeAppKey != testCase.someAppKey{
					t.Errorf("ConfigFile.SomeAppKey = %v, want %v", file.SomeAppKey, testCase.someAppKey)
				}
				if file.SomeAppKey != testCase.someAppKey{
					t.Errorf("ConfigFile.SomeAppKey = %v, want %v", file.SomeAppKey, testCase.someAppKey)
				}
			})

	}
}
