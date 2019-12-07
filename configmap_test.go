package configmap

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestUnit_Set(t *testing.T) {
	str := "bar"
	cm := NewConfigMap()
	cm.Set("foo", str)
	act := cm.Get("foo")
	if act != str {
		t.Fatalf("act != str")
	}
}

func TestUnit_GetStringOrNil(t *testing.T) {
	cm := NewConfigMap()
	cm.Set("foo", "bar")
	baz := cm.GetStringOrNil("baz")
	bar := cm.GetString("foo")
	if baz != nil {
		t.Fatalf("baz != nil")
	}
	if bar != "bar" {
		t.Fatalf("bar != bar")
	}
}