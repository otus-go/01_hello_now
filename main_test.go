package main_test

import (
	"testing"
	"time"
	"bytes"
	"fmt"

	"github.com/stretchr/testify/assert"

	hello "github.com/slonegd/otus/01_hello_now"
)

func TestPrintNow(t *testing.T) {
	want := bytes.Buffer{}
	fmt.Fprintln(&want, time.Now().Format("2006-01-02 15:04:05"))

	got := bytes.Buffer{}
	hello.PrintNow(&got)

	assert.Equal(t, want, got)
}