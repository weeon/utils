package utils

import "testing"

func TestNewUUID(t *testing.T) {
	t.Log(NewUUID())
	t.Log(NewUUID())
}
