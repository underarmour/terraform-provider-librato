package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

func TestProvider(t *testing.T) {
	provider := newProvider().(*schema.Provider)
	if err := provider.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
