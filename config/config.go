package config

import (
	"fmt"
	"log"
	"time"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
)

var (
	k    = koanf.New(".")
	Conf Config
)

type Config struct {
	Port                int        `koanf:"port"`
	Storage             string     `koanf:"storage"`
	TaskFilePath        string     `koanf:"task_file_path"`
	SendBackgroundEmail bool       `koanf:"send_background_email"`
	EmailGroupBy        string     `koanf:"email_group_by"`
	EmailFreq           string     `koanf:"email_freq"`
	SMTP                SMTPConfig `koanf:"smtp"`
}

type SMTPConfig struct {
	Email    string `koanf:"email"`
	Password string `koanf:"password"`
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	To       string `koanf:"to"`
}

// LoadConfig loads the configuration from config.json
func LoadConfig() error {
	f := file.Provider("config/config.json")
	if err := k.Load(f, json.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	return k.Unmarshal("", &Conf)
}

func WatchConfig(onChange func()) {
	provider := file.Provider("config/config.json")

	go func() {
		err := provider.Watch(func(event interface{}, err error) {
			if err != nil {
				fmt.Printf("Watch error: %v\n", err)
				return
			}
			fmt.Println("config.json changed content")
			if err := k.Load(provider, json.Parser()); err != nil {
				fmt.Printf("Reload error: %v\n", err)
				return
			}
			if err := k.Unmarshal("", &Conf); err != nil {
				fmt.Printf("Unmarshal error: %v\n", err)
				return
			}
			onChange()
		})
		if err != nil {
			fmt.Printf("Failed to start watch: %v\n", err)
		}
	}()
}

// EmailFrequencyDuration returns EmailFreq as time.Duration
func EmailFrequencyDuration() (time.Duration, error) {
	return time.ParseDuration(Conf.EmailFreq)
}
