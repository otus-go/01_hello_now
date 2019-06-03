package main

import (
	"errors"
	"time"
)

func SetTime(t time.Time) {
	ntpTime = func(string) (time.Time, error) {
		return t, nil
	}
}

func SetError() {
	ntpTime = func(string) (time.Time, error) {
		return time.Time{}, errors.New("")
	}
}
