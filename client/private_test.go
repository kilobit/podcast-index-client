/* Copyright 2021 Kilobit Labs Inc. */

package client

import _ "fmt"
import _ "errors"

import "kilobit.ca/go/tested/assert"
import "testing"

func TestPICClientPrivateTest(t *testing.T) {
	assert.Expect(t, true, true, "Failed sanity check.")
}

func TestCalculateAuth(t *testing.T) {

	tests := []struct {
		key     string
		secret  string
		datestr string
		exp     string
	}{
		{
			"THIRTYTHREETHIRTYTHREETHIRTYTHREE",
			"27c5f70d895655f3ed34d1b7380f2365dda26572",
			"1612224299",
			"f075b168719227b3e8880b5f99ab7449721db322",
		},
	}

	for _, test := range tests {

		act := calculateAuth(test.key, test.secret, test.datestr)
		assert.Expect(t, test.exp, act, test)
	}
}
