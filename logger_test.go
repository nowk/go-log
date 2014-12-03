package log

import (
	"bytes"
	"github.com/nowk/assert"
	"log"
	"testing"
)

func TestMultipleWriters(t *testing.T) {
	var _a []byte
	var _b []byte
	a := bytes.NewBuffer(_a)
	b := bytes.NewBuffer(_b)

	l := New("PREFIX: ", 0, a, b)
	err := l.Output(2, "Hello World!")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "PREFIX: Hello World!\n", a.String())
	assert.Equal(t, "PREFIX: Hello World!\n", b.String())
}

func TestWithLogLevelValue(t *testing.T) {
	var b []byte
	w := bytes.NewBuffer(b)

	l := New("PREFIX: ", 0, w)
	err := l.Log(INFO, "%s!", "Hello World")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "PREFIX: [INFO] Hello World!\n", w.String())
}

func TestOutputAtOrAbove(t *testing.T) {
	SetLogLevel(WARN)
	defer SetLogLevel(INFO)

	var b []byte
	w := bytes.NewBuffer(b)

	l := New("", 0, w)
	err := l.Info("%s!", "Hello World")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "", w.String())

	err = l.Warn("All your bases!")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "[WARN] All your bases!\n", w.String())
}

func TestUsingLoggerInterface(t *testing.T) {
	var b []byte
	w := bytes.NewBuffer(b)

	l := log.New(w, "PREFIX: ", 0)
	err := Log(l, INFO, "%s!", "Hello World")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "PREFIX: [INFO] Hello World!\n", w.String())
}
