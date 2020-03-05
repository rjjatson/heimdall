package hub

import (
	"heimdall/internal/client"
	"testing"
)

func TestClientAccess(t *testing.T) {
	h := &Hub{}
	ec := &client.Client{}

	h.AddClient("1", ec)
	ac, err := h.GetClient("1")
	if err != nil {
		t.Error("fail retrieving client")
	}
	if ac != ec {
		t.Error("incorrect retrieved client")
	}
}
