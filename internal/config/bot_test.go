package config

import (
	"os"
	"path/filepath"
	"testing"
)

func writeBot(t *testing.T, contents string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "bot.yaml")
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestLoadAppliesDefaults(t *testing.T) {
	path := writeBot(t, `
name: helper
system_prompt: Be helpful.
model: llama3.2
`)
	bot, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if bot.BaseURL != DefaultBaseURL {
		t.Fatalf("BaseURL = %q, want %q", bot.BaseURL, DefaultBaseURL)
	}
	if bot.APIKeyEnv != DefaultAPIKeyEnv {
		t.Fatalf("APIKeyEnv = %q, want %q", bot.APIKeyEnv, DefaultAPIKeyEnv)
	}
}

func TestLoadMissingName(t *testing.T) {
	path := writeBot(t, `
system_prompt: Be helpful.
model: llama3.2
`)
	if _, err := Load(path); err == nil {
		t.Fatal("expected error for missing name")
	}
}

func TestLoadMissingFile(t *testing.T) {
	if _, err := Load(filepath.Join(t.TempDir(), "missing.yaml")); err == nil {
		t.Fatal("expected error for missing file")
	}
}
