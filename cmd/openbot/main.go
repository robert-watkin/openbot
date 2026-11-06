package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/robert-watkin/openbot/internal/config"
	"github.com/robert-watkin/openbot/internal/llm"
)

const version = "0.1.0"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "openbot: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		printUsage(os.Stderr)
		return fmt.Errorf("expected a command")
	}

	switch args[0] {
	case "version", "--version", "-v":
		fmt.Println(version)
		return nil
	case "help", "--help", "-h":
		printUsage(os.Stdout)
		return nil
	case "chat":
		return runChat(args[1:])
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func runChat(args []string) error {
	fs := flag.NewFlagSet("chat", flag.ContinueOnError)
	botPath := fs.String("bot", "", "path to a bot YAML file")
	fs.SetOutput(os.Stderr)
	if err := fs.Parse(args); err != nil {
		return err
	}
	message := strings.TrimSpace(strings.Join(fs.Args(), " "))
	if *botPath == "" {
		return fmt.Errorf("chat: --bot is required")
	}
	if message == "" {
		return fmt.Errorf("chat: a message is required")
	}

	bot, err := config.Load(*botPath)
	if err != nil {
		return err
	}

	client := llm.Client{
		BaseURL: bot.BaseURL,
		APIKey:  bot.APIKey(),
	}
	reply, err := client.Chat(context.Background(), bot.Model, bot.SystemPrompt, message)
	if err != nil {
		return err
	}
	fmt.Println(reply)
	return nil
}

func printUsage(w *os.File) {
	fmt.Fprint(w, `OpenBot: self-host your own AI agents

Usage:
  openbot chat --bot <file> <message>
  openbot version
  openbot help
`)
}
