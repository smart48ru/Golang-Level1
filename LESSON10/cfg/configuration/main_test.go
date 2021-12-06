package configuration_test

import (
	"fmt"
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
	}{
		//Первый сценарий
		{
			name:      "../config.json",
			format:    "json",
			filename:  "../config.json",
			want:      nil,
			jaegerUrl: "http://jaeger:16686",
			sentryUrl: "http://sentry:9000",
		},
		//Второй сценарий
		{
			name:      "../config.yaml",
			format:    "yaml",
			filename:  "../config.yaml",
			want:      nil,
			jaegerUrl: "http://jaeger:16686",
			sentryUrl: "http://sentry:9000",
		},
	}

	for _, testCase := range testTable {
		Cfg.JaegerUrl = testCase.jaegerUrl
		Cfg.SentryUrl = testCase.sentryUrl


			_, err := СheckCfgFromFile(testCase.filename, testCase.format)
		fmt.Println(testCase.name)

			t.Run(testCase.name, func(t *testing.T) {
				if err != testCase.want {
					fmt.Printf("%T\n",err)
					fmt.Printf("%T\n",testCase.want)
					t.Errorf("ConfigFile() = %v, want %v", err, testCase.want)
				}
			})

	}
}
