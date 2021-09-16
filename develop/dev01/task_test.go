package main

import (
	"testing"
	"time"

	"github.com/beevik/ntp"
)

func TestTimeNow(t *testing.T) {

	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		t.Errorf("Query ntp error: %v\n", err)
	}
	want := time.Now().Add(response.ClockOffset)
	got := timeNow()
	if want.Equal(got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
