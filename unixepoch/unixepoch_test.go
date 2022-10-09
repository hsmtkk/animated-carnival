package unixepoch_test

import (
	"testing"
	"time"

	"github.com/hsmtkk/animated-carnival/unixepoch"
	"github.com/stretchr/testify/assert"
)

func TestStringToTime(t *testing.T) {
	want, err := time.Parse(time.RFC3339, "2019-06-08T22:20:00+09:00")
	assert.Nil(t, err)
	got, err := unixepoch.StringToTime("1560000000000000")
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
