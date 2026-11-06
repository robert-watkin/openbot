# OpenBot

Self-host your own AI agents: local models or any API, with optional isolated runtimes.

OpenBot is a personal open-source project for running a team of AI bots on hardware you control. Today it can load one bot from a YAML file and complete a single chat turn against a local model server such as Ollama, or any OpenAI-compatible API.

The long-term vision is an open-source, self-hosted personal agent system that is easy enough to get running in a few minutes. There is a backend that manages bots, memory, connectors, and actually running them, and a web interface where you can see your bots in a sidebar, customize them, and open chats and group chats. The web UI comes first. Native wrappers for desktop or phone can come later if they are useful. If you run OpenBot on your own machine, the interface is just a local page in your browser, so it still works without Wi-Fi. If you host it on a server or VPS, you will need a network path to reach it.

You run multiple bots, each with their own name and character. Those bots can join group chats, message each other, keep their own tasks and routines, and connect to third-party apps. They can also help manage the system themselves: if you ask a bot to create another bot or set up its routines, it should be able to do that, with safety rules that do not depend on which model you are using. Anything like making a payment always needs your approval.

Setup should stay out of the way. After you install OpenBot, you connect it to Ollama (or another local server) and pick a model. OpenBot should detect available models where it can, and still let you choose from a dropdown or point it at an API instead. If you want a bot to have its own computer, you can turn that on and OpenBot will download the tooling needed to run a virtual machine on your device. That local option will have performance limits, so you can also connect a virtual machine hosted somewhere else.

The aim is that within about five minutes of install you have your first bot you can talk to, and it can have a computer.

## Quick start

You need Go 1.22 or newer. For a local model, run [Ollama](https://ollama.com/) (or another OpenAI-compatible server) separately.

```bash
git clone https://github.com/robert-watkin/openbot.git
cd openbot
go build -o openbot ./cmd/openbot

# Optional: many local servers do not need a key
# export OPENBOT_API_KEY=...

./openbot chat --bot examples/bots/assistant.yaml "What can you help me with?"
./openbot version
```

By default the example bot talks to `http://127.0.0.1:11434/v1` (Ollama's OpenAI-compatible endpoint) using the `llama3.2` model. Change `base_url` and `model` in the bot file to use a different server or model. If the server expects a key, set the environment variable named in `api_key_env` (default `OPENBOT_API_KEY`).

## Current status

- [x] MIT license and public source
- [x] YAML bot definition (`name`, `system_prompt`, `model`, `base_url`, `api_key_env`)
- [x] One-shot `openbot chat` against an OpenAI-compatible `/chat/completions` API
- [x] Unit tests against a mock server (no live model required)

## Roadmap

1. Easy first-run setup: install, connect Ollama or an API, detect local models, and pick one.
2. A web interface with your bots in a sidebar, room to customize them, and chats plus group chats. Native app wrappers can come later.
3. A backend that runs bots, stores memory, and manages connectors.
4. Multiple bots with their own names and characters, including choosing a bot by name.
5. Group chats, and bots that can message each other.
6. Tasks and routines for each bot, including letting a bot set those up when you ask.
7. Bot self-management (creating other bots, changing settings) with model-independent safety rules, including always asking you before a payment or similar action.
8. Streaming replies and lasting conversation memory.
9. Connectors to third-party applications, under your control.
10. An optional computer for each bot: a local toggle that installs the virtual machine tooling, plus the option to connect a machine hosted somewhere else.
11. Tools with permissions, so a bot can only read, write, or run what you have granted.
12. Automated tests, release builds, a short self-host guide, and a threat model covering secrets, filesystem access, and network access.

## License

MIT. See [LICENSE](LICENSE).
