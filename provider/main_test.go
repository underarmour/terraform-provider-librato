package provider

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestNewRequest(t *testing.T) {
	p := NewProvider("user", "token", "baseurl")
	if _, err := p.NewRequest("", "", nil); err != nil {
		t.Fatal("Failed to get expected error")
	}
	data := map[string]string{
		"foo": "bar",
	}
	bs, err := json.Marshal(data)
	if err != nil {
		t.Fatal("Unable to marshal data: ", err)
	}
	req, err := p.NewRequest("method", "url", bytes.NewReader(bs))
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}
	user, token, _ := req.BasicAuth()
	if user != p.user {
		t.Fatal("Incorrect user in basic auth")
	}
	if token != p.token {
		t.Fatal("Incorrect token in basic auth")
	}
}

func TestNewProvider(t *testing.T) {
	NewProvider("foo", "bar", "baz")
}
