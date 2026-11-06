package llm

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChatSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/chat/completions" {
			t.Fatalf("path = %q", r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer test-key" {
			t.Fatalf("Authorization = %q", got)
		}
		raw, _ := io.ReadAll(r.Body)
		var payload map[string]any
		if err := json.Unmarshal(raw, &payload); err != nil {
			t.Fatal(err)
		}
		if payload["model"] != "llama3.2" {
			t.Fatalf("model = %v", payload["model"])
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"choices":[{"message":{"role":"assistant","content":"Hello from the mock."}}]}`)
	}))
	defer server.Close()

	client := Client{BaseURL: server.URL, APIKey: "test-key", HTTPClient: server.Client()}
	got, err := client.Chat(context.Background(), "llama3.2", "Be helpful.", "Hi")
	if err != nil {
		t.Fatal(err)
	}
	if got != "Hello from the mock." {
		t.Fatalf("got %q", got)
	}
}

func TestChatServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = io.WriteString(w, `{"error":{"message":"model unavailable"}}`)
	}))
	defer server.Close()

	client := Client{BaseURL: server.URL, HTTPClient: server.Client()}
	if _, err := client.Chat(context.Background(), "llama3.2", "Be helpful.", "Hi"); err == nil {
		t.Fatal("expected error")
	}
}

func TestChatMalformedJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "not-json")
	}))
	defer server.Close()

	client := Client{BaseURL: server.URL, HTTPClient: server.Client()}
	if _, err := client.Chat(context.Background(), "llama3.2", "Be helpful.", "Hi"); err == nil {
		t.Fatal("expected error")
	}
}
