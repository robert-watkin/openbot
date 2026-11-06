package config

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	DefaultBaseURL   = "http://127.0.0.1:11434/v1"
	DefaultAPIKeyEnv = "OPENBOT_API_KEY"
)

// Bot is a YAML definition of a single assistant.
type Bot struct {
	Name         string `yaml:"name"`
	SystemPrompt string `yaml:"system_prompt"`
	Model        string `yaml:"model"`
	BaseURL      string `yaml:"base_url"`
	APIKeyEnv    string `yaml:"api_key_env"`
}

// Load reads a bot YAML file and applies defaults.
func Load(path string) (Bot, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Bot{}, fmt.Errorf("read bot file %q: %w", path, err)
	}

	var bot Bot
	if err := yaml.Unmarshal(data, &bot); err != nil {
		return Bot{}, fmt.Errorf("parse bot file %q: %w", path, err)
	}

	bot.Name = strings.TrimSpace(bot.Name)
	bot.SystemPrompt = strings.TrimSpace(bot.SystemPrompt)
	bot.Model = strings.TrimSpace(bot.Model)
	bot.BaseURL = strings.TrimSpace(bot.BaseURL)
	bot.APIKeyEnv = strings.TrimSpace(bot.APIKeyEnv)

	if bot.Name == "" {
		return Bot{}, fmt.Errorf("bot file %q: name is required", path)
	}
	if bot.SystemPrompt == "" {
		return Bot{}, fmt.Errorf("bot file %q: system_prompt is required", path)
	}
	if bot.Model == "" {
		return Bot{}, fmt.Errorf("bot file %q: model is required", path)
	}
	if bot.BaseURL == "" {
		bot.BaseURL = DefaultBaseURL
	}
	if bot.APIKeyEnv == "" {
		bot.APIKeyEnv = DefaultAPIKeyEnv
	}

	return bot, nil
}

// APIKey returns the key from the environment variable named by the bot, which may be empty.
func (b Bot) APIKey() string {
	return os.Getenv(b.APIKeyEnv)
}
