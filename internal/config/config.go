package config

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/jonathanblade/yagpt-shell/internal/style"
	"github.com/spf13/viper"
)

const ConfigName = ".yagpt-shell"

type Config struct {
	ApiKey      string  `mapstructure:"API_KEY"`
	FolderID    string  `mapstructure:"FOLDER_ID"`
	Temperature float64 `mapstructure:"TEMPERATURE"`
}

func Read() *Config {
	viper.SetConfigName(ConfigName)
	viper.SetConfigType("env")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("%s not found. Please create it.", ConfigName)
		} else {
			log.Fatalf("Failed to read %s: %v", ConfigName, err)
		}
	}
	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalf("Failed to read %s. Make sure you set correct API_KEY, FOLDER_ID and TEMPERATURE.", ConfigName)
	}
	if conf.Temperature < 0 || conf.Temperature > 1 {
		log.Fatalf("Failed to read %s. Make sure you set correct TEMPERATURE (min = 0, max = 1).", ConfigName)
	}
	return &conf
}

func (c *Config) Show() {
	text := strings.Builder{}

	api_key := style.AccentTextStyle.Render("API key: ")
	api_key += lipgloss.NewStyle().PaddingLeft(4).Render(c.ApiKey)
	text.WriteString(api_key + "\n\n")

	folder_id := style.AccentTextStyle.Render("Folder ID: ")
	folder_id += lipgloss.NewStyle().PaddingLeft(2).Render(c.FolderID)
	text.WriteString(folder_id + "\n\n")

	temperature := style.AccentTextStyle.Render("Temperature: ")
	temperature += strconv.FormatFloat(c.Temperature, 'f', -1, 64)
	text.WriteString(temperature + "\n")

	fmt.Println(style.BorderStyle.Render(text.String()))
}
