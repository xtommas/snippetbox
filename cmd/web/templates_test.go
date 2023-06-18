package main

import (
	"testing"
	"time"

	"github.com/xtommas/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	// Slice of anonymous structs that contains the test cases
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2023, 6, 18, 11, 55, 0, 0, time.UTC),
			want: "18 Jun 2023 at 11:55",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2022, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Mar 2022 at 09:15",
		},
	}

	// loop through the slice to test each case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}
