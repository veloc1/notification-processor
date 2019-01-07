package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockedProcessor struct {
	Processor

	isProcessCalled bool
}

func (p *MockedProcessor) canHandle(script string) bool {
	return script == "bitbucket"
}

func (p *MockedProcessor) process(r *http.Request) (NotifyData, error) {
	p.isProcessCalled = true
	return NotifyData{}, nil
}

func TestNoService(t *testing.T) {
	handler := &WebhookHandler{}

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code incorrect, got: %d, want: %d.", resp.StatusCode, http.StatusBadRequest)
	}
}

func TestWrongService(t *testing.T) {
	handler := &WebhookHandler{}

	req := httptest.NewRequest("GET", "/?service=aaabbb", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code incorrect, got: %d, want: %d.", resp.StatusCode, http.StatusBadRequest)
	}
}

func TestBitbucketService(t *testing.T) {
	processor := &MockedProcessor{}
	handler := &WebhookHandler{
		processors: []Processor{
			processor,
		},
	}

	req := httptest.NewRequest("GET", "/?service=bitbucket", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code incorrect, got: %d, want: %d.", resp.StatusCode, http.StatusOK)
	}

	if !processor.isProcessCalled {
		t.Errorf("Method process was not called")
	}
}
