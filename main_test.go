package main_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	hello "github.com/slonegd/otus/01_hello_now"
)

func TestPrintNow(t *testing.T) {
	date := time.Date(2019, time.May, 30, 12, 11, 10, 0, time.Local)
	hello.SetTime(date)

	got := bytes.Buffer{}
	hello.PrintNow(&got)

	assert.Equal(t, "2019-05-30 12:11:10\n", got.String())
}

func TestPrintNowError(t *testing.T) {
	hello.SetError()

	assert.Error(t, hello.PrintNow(&bytes.Buffer{}))
}
