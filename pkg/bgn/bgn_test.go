package bgn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BGN_Parse(t *testing.T) {
	tests := []struct {
		name        string
		bgn         string
		shouldError bool
	}{
		{
			name:        "empty string should error",
			bgn:         "",
			shouldError: true,
		},
		{
			name:        "random string should error",
			bgn:         "12al12wt31nwauhwaj7haw",
			shouldError: true,
		},
		{
			name:        "missing Teams tag should error",
			bgn:         "[Seed \"123\"][Game \"Carcassonne\"]",
			shouldError: true,
		},
		{
			name:        "tags only should succeed",
			bgn:         "[Teams \"red, blue\"][Seed \"123\"][Game \"Carcassonne\"]",
			shouldError: false,
		},
		{
			name:        "nested right bracket should error",
			bgn:         "[[Teams] \"red, blue\"][Seed \"123\"][Game \"Carcassonne\"]",
			shouldError: true,
		},
		{
			name:        "nested left bracket should error",
			bgn:         "[Teams]] \"red, blue\"][Seed \"123\"][Game \"Carcassonne\"]",
			shouldError: true,
		},
		{
			name:        "missing action key should error",
			bgn:         "[Teams \"red, blue\"][Seed \"123\"][Game \"Carcassonne\"]0",
			shouldError: true,
		},
		{
			name:        "missing team index should error",
			bgn:         "[Teams \"red, blue\"][Seed \"123\"][Game \"Carcassonne\"]a",
			shouldError: true,
		},
		{
			name:        "large team index should succeed",
			bgn:         "[Teams \"red, blue\"][Seed \"123\"][Game \"Carcassonne\"]10a",
			shouldError: false,
		},
		{
			name:        "too many ampersands should error",
			bgn:         "[Teams \"red, blue\"][Seed \"123\"][Game \"Carcassonne\"]10a&1.2&1.1",
			shouldError: true,
		},
		{
			name:        "complex bgn should succeed",
			bgn:         "[Teams \"red, blue\"][Seed \"123\"][Game \"Carcassonne\"][Date \"11-01-2021\"]0a&1.2 {Comment} 0b&1.2.T.R 1c 1a&1.3 1b&1.3.K.B",
			shouldError: false,
		},
		{
			name: "multi line tags should succeed",
			bgn: `
			[Game "Tsuro"]
			[Teams "red, blue"]
			[Variant "Classic"]
			[Seed "1696036843787"]
	
			0p&4.5.CHDAEBFG`,
			shouldError: false,
		},
		{
			name: "multi line actions should succeed",
			bgn: `[Game "Tsuro"][Teams "red, blue"][Variant "Classic"][Seed "1696036843787"]

			0p&4.5.CHDAEBFG 1p&0.4.CDEAFGHB 0p&3.5.ACBGDHEF 1p&0.3.AEBDCGFH
			0p&5.5.EDFGHCAB
			1p&0.2.CHDFEAGB`,
			shouldError: false,
		},
		{
			name: "weird formatting should succeed",
			bgn: `
			
			[Game "Tsuro"]
			
		 [Teams "red, blue"]
			  
					[Variant "Classic"]
[Seed "1696036843787"]

			0p&4.5.CHDAEBFG   1p&0.4.CDEAFGHB 	
				0p&3.5.ACBGDHEF 1p&0.3.AEBDCGFH
			0p&5.5.EDFGHCAB

			 1p&0.2.CHDFEAGB
			
			`,
			shouldError: false,
		},
	}

	for _, test := range tests {
		_, err := Parse(test.bgn)
		assert.Equal(t, test.shouldError, err != nil, test.name)
	}
}
