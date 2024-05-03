package main

import (
	"testing"
	"time"
)

func TestGetDateOfBirth(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  time.Time
	}{
		{"29.01.1988 should be", "880129357787", time.Date(1988, time.Month(1), 29, 0, 0, 0, 0, time.UTC)},
		{"29.01.1993 should be", "930129357787", time.Date(1993, time.Month(1), 29, 0, 0, 0, 0, time.UTC)},
		{"29.01.2000 should be", "000129557787", time.Date(2000, time.Month(1), 29, 0, 0, 0, 0, time.UTC)},
		{"29.01.1888 should be", "880129157787", time.Date(1888, time.Month(1), 29, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := GetDateOfBirth(tt.input)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGender(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"male should be", "880129357787", "male"},
		{"female should be", "93012945678", "female"},
		{"digit 7 is not valid", "93012905678", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := GetGender(tt.input)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifyIIN(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  bool
	}{
		{"valid IIN should be", "991229351245", true},
		{"invalid IIN should be", "93012945678", false},
		{"invalid IIN should be", "1293", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VerifyIIN(tt.input)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
